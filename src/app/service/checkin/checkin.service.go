package checkin

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type IRepository interface {
	Checkin(*checkin.Checkin) error
	Checkout(*checkin.Checkin) error
	FindLastCheckin(string, int32, *checkin.Checkin) error
}

type ICacheRepository interface {
	SaveCache(string, interface{}, int) error
	GetCache(string, interface{}) error
	RemoveCache(key string) error
}

type Service struct {
	repo  IRepository
	cache ICacheRepository
	conf  config.App
}

func NewService(repo IRepository, cache ICacheRepository, conf config.App) *Service {
	return &Service{repo: repo, cache: cache, conf: conf}
}

func (s *Service) CheckinVerify(_ context.Context, req *proto.CheckinVerifyRequest) (*proto.CheckinVerifyResponse, error) {
	ciToken := &checkin.CheckinToken{}
	err := s.cache.GetCache(req.Id, ciToken)

	if err != redis.Nil {
		return &proto.CheckinVerifyResponse{
			CheckinToken: ciToken.Token,
			CheckinType:  ciToken.CheckinType,
		}, nil
	}

	ci := &checkin.Checkin{}
	err = s.repo.FindLastCheckin(req.Id, req.EventType, ci)

	var checkinType string

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().
			Err(err).
			Str("service", "checkin").
			Str("module", "checkinverify").
			Msg("Unknown db error")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	} else if errors.Is(err, gorm.ErrRecordNotFound) || ci.CheckoutAt != nil {
		checkinType = "check_in"
	} else {
		checkinType = "check_out"
	}

	_token, err := uuid.NewUUID()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "checkin").
			Str("module", "checkinverify").
			Msg("UUID broken")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	token := _token.String()

	s.cache.SaveCache(token, &checkin.TokenInfo{
		Id:          req.Id,
		CheckinType: checkinType,
		EventType:   req.EventType,
	}, 60)

	s.cache.SaveCache(req.Id, &checkin.CheckinToken{
		Token:       token,
		UserId:      req.Id,
		CheckinType: checkinType,
	}, 60)

	res := &proto.CheckinVerifyResponse{
		CheckinToken: token,
		CheckinType:  checkinType,
	}

	return res, nil
}

func (s *Service) CheckinConfirm(_ context.Context, req *proto.CheckinConfirmRequest) (*proto.CheckinConfirmResponse, error) {
	cached := &checkin.TokenInfo{}

	err := s.cache.GetCache(req.Token, cached)

	if err == redis.Nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	defer func() {
		if err := s.cache.RemoveCache(cached.Id); err != nil {
			log.Error().Err(err).
				Str("service", "Checkin").
				Str("module", "checkin confirm").
				Msg("Error while removing user cache")
		}
	}()

	defer func() {
		if err := s.cache.RemoveCache(req.Token); err != nil {
			log.Error().Err(err).
				Str("service", "Checkin").
				Str("module", "checkin confirm").
				Msg("Error while removing token cache")
		}
	}()

	switch cached.CheckinType {
	case "check_in":
		ci := newCheckin(cached.Id, cached.EventType)
		err = s.repo.Checkin(ci)
	case "check_out":
		ci := newCheckout(cached.Id, cached.EventType)
		err = s.repo.Checkout(ci)
	default:
		return nil, status.Error(codes.InvalidArgument, "Invalid checkin type")
	}

	if err != nil {
		log.Error().Err(err).
			Str("service", "Checkin").
			Str("module", "checkin confirm").
			Msg("Error while Checkin, Possibly due to invalid user-uuid")

		return nil, status.Error(codes.Internal, "Internal Error")
	}

	return &proto.CheckinConfirmResponse{
		Success: true,
	}, nil
}

func newCheckin(userid string, eventType int32) *checkin.Checkin {
	return &checkin.Checkin{
		UserId:    userid,
		EventType: eventType,
	}
}

func newCheckout(userid string, eventType int32) *checkin.Checkin {
	return &checkin.Checkin{
		UserId:     userid,
		CheckoutAt: utils.GetCurrentTimePtr(),
		EventType:  eventType,
	}
}

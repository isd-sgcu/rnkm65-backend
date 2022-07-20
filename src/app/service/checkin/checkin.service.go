package checkin

import (
	"context"

	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IRepository interface {
	Checkin(*checkin.Checkin) error
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Checkin(_ context.Context, req *proto.CheckInRequest) (*proto.CheckinResponse, error) {
	ci := newCheckin(req.Id)

	err := s.repo.Checkin(ci)

	if err == nil {
		return &proto.CheckinResponse{
			Success: true,
		}, nil
	}

	log.Error().
		Err(err).
		Str("service", "checkin").
		Str("module", "checkin").
		Msg("malformed data (invalid format / invalid userid)")

	return nil, status.Error(codes.InvalidArgument, "malformed data (invalid format / invalid userid)")
}

func newCheckin(userid string) *checkin.Checkin {
	return &checkin.Checkin{
		UserId: userid,
	}
}

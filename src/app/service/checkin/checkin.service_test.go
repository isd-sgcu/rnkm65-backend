package checkin

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	cmock "github.com/isd-sgcu/rnkm65-backend/src/mocks/cache"
	rmock "github.com/isd-sgcu/rnkm65-backend/src/mocks/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CheckinServiceTest struct {
	suite.Suite
	Ci          *checkin.Checkin
	Token       string
	CiToken     *checkin.CheckinToken
	CiTokenInfo *checkin.TokenInfo
	Conf        config.App
}

func TestCheckinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	t.Ci = &checkin.Checkin{
		UserId:     uuid.New().String(),
		EventType:  1,
		CheckinAt:  utils.GetCurrentTimePtr(),
		CheckoutAt: utils.GetCurrentTimePtr(),
	}

	t.Token = uuid.New().String()

	t.CiToken = &checkin.CheckinToken{
		Token:       t.Token,
		UserId:      t.Ci.UserId,
		CheckinType: "check_in",
	}

	t.CiTokenInfo = &checkin.TokenInfo{
		Id:          t.Ci.UserId,
		CheckinType: "check_in",
		EventType:   1,
	}

	t.Conf = config.App{
		CICacheTTL: 180,
	}
}

func (t *CheckinServiceTest) TestCheckinVerifyCached() {
	want := &proto.CheckinVerifyResponse{
		CheckinToken: t.Token,
		CheckinType:  "check_in",
	}

	req := &proto.CheckinVerifyRequest{
		Id:        t.Ci.UserId,
		EventType: t.Ci.EventType,
	}

	ciToken := &checkin.CheckinToken{}

	repo := &rmock.RepositoryMock{}

	cr := newCacheMock()
	cr.On("GetCache", req.Id, ciToken).Return(t.CiToken, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinVerify(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), actual, want)
}

func (t *CheckinServiceTest) TestCheckinVerifyNewToken() {
	req := &proto.CheckinVerifyRequest{
		Id:        t.Ci.UserId,
		EventType: t.Ci.EventType,
	}

	repo := &rmock.RepositoryMock{}
	ciToken := &checkin.CheckinToken{}
	cr := newCacheMock()

	repo.On("FindLastCheckin", t.Ci.UserId, t.Ci.EventType, &checkin.Checkin{}).Return(nil, gorm.ErrRecordNotFound)

	// New Token is generate in service
	cr.On("GetCache", req.Id, ciToken).Return(nil, redis.Nil)
	cr.On("SaveCache", mock.Anything, t.CiTokenInfo, 60).Return(nil)
	cr.On("SaveCache", t.Ci.UserId, mock.Anything, 60).Return(nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinVerify(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), len(actual.CheckinToken), 36)
}

func (t *CheckinServiceTest) TestCheckinVerifyUnknown() {
	req := &proto.CheckinVerifyRequest{
		Id:        t.Ci.UserId,
		EventType: t.Ci.EventType,
	}

	repo := &rmock.RepositoryMock{}

	ci := &checkin.Checkin{}
	repo.On("FindLastCheckin", t.Ci.UserId, t.Ci.EventType, ci).Return(nil, gorm.ErrInvalidData)

	cr := newCacheMock()
	cr.On("GetCache", req.Id, &checkin.CheckinToken{}).Return(nil, redis.Nil)

	service := NewService(repo, cr, t.Conf)

	_, err := service.CheckinVerify(context.Background(), req)

	assert.NotNil(t.T(), err)
}

func (t *CheckinServiceTest) TestCheckinVerifyInvalidToken() {
	repo := &rmock.RepositoryMock{}

	randomUUID, _ := uuid.NewUUID()

	req := &proto.CheckinConfirmRequest{
		Token: randomUUID.String(),
	}

	cr := newCacheMock()
	cr.On("GetCache", mock.AnythingOfType("string"), &checkin.TokenInfo{}).Return(nil, redis.Nil)

	service := NewService(repo, cr, t.Conf)

	_, err := service.CheckinConfirm(context.Background(), req)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Equal(t.T(), codes.Unauthenticated, st.Code())
}

func (t *CheckinServiceTest) TestCheckinConfirmUnknown() {
	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	repo := &rmock.RepositoryMock{}
	repo.On("Checkin", newCheckin(t.Ci.UserId, t.Ci.EventType)).Return(status.Error(codes.Internal, "Internal Error"))

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(&checkin.TokenInfo{
		Id:          t.Ci.UserId,
		CheckinType: "check_in",
		EventType:   t.Ci.EventType,
	}, nil)

	service := NewService(repo, cr, t.Conf)

	_, err := service.CheckinConfirm(context.Background(), req)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *CheckinServiceTest) TestCheckinConfirmSuccess() {
	want := &proto.CheckinConfirmResponse{
		Success: true,
	}

	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	repo := &rmock.RepositoryMock{}
	repo.On("Checkin", newCheckin(t.Ci.UserId, t.Ci.EventType)).Return(nil)

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(&checkin.TokenInfo{
		Id:          t.Ci.UserId,
		CheckinType: "check_in",
		EventType:   t.Ci.EventType,
	}, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinConfirm(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CheckinServiceTest) TestCheckoutConfirmSuccess() {
	want := &proto.CheckinConfirmResponse{
		Success: true,
	}

	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	repo := &rmock.RepositoryMock{}

	// repo.On("Checkout", co) cannot be used here
	// since There's little delay between
	// generating from t.ci and in service.Checkconfirm
	// args.Called cannot be used in the situation

	repo.On("Checkout", mock.Anything).Return(nil)

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(&checkin.TokenInfo{
		Id:          t.Ci.UserId,
		CheckinType: "check_out",
		EventType:   t.Ci.EventType,
	}, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinConfirm(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CheckinServiceTest) TestCheckinConfirmInvalidUserId() {
	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	randomUUID := uuid.New().String()

	repo := &rmock.RepositoryMock{}

	ci := newCheckin(randomUUID, t.Ci.EventType)

	repo.On("Checkin", ci).Return(errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails"))

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(&checkin.TokenInfo{
		Id:          randomUUID,
		CheckinType: "check_in",
		EventType:   t.Ci.EventType,
	}, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinConfirm(context.Background(), req)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func newCacheMock() *cmock.RepositoryMock {
	return &cmock.RepositoryMock{
		V: make(map[string]interface{}),
	}
}

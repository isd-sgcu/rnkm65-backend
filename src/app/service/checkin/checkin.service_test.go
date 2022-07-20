package checkin

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckinServiceTest struct {
	suite.Suite
	Ci *checkin.Checkin
}

func TestCheckinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	t.Ci = &checkin.Checkin{
		UserId: uuid.New().String(),
	}
}

func (t *CheckinServiceTest) TestCheckinSuccess() {
	want := &proto.CheckinResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Checkin", t.Ci).Return(nil)

	service := NewService(repo)

	actual, err := service.Checkin(context.Background(), &proto.CheckInRequest{Id: t.Ci.UserId})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CheckinServiceTest) TestCheckinFailed() {
	repo := &mock.RepositoryMock{}

	repo.On("Checkin", t.Ci).Return(status.Error(codes.Unknown, "any error"))

	service := NewService(repo)

	actual, err := service.Checkin(context.Background(), &proto.CheckInRequest{Id: t.Ci.UserId})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
	assert.Nil(t.T(), actual)
}

func (t *CheckinServiceTest) TestCheckinNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Checkin", t.Ci).Return(status.Error(codes.Unknown, "any error"))

	service := NewService(repo)

	actual, err := service.Checkin(context.Background(), &proto.CheckInRequest{Id: t.Ci.UserId})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
	assert.Nil(t.T(), actual)
}

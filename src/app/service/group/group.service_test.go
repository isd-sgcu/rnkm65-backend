package group

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/group"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"testing"
	"time"
)

type GroupServiceTest struct {
	suite.Suite
	Group              *group.Group
	GroupDto           *proto.Group
	CreateGroupReqMock *proto.CreateGroupRequest
	UpdateGroupReqMock *proto.UpdateGroupRequest
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(GroupServiceTest))
}

func (t *GroupServiceTest) SetupTest() {
	t.Group = &group.Group{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: faker.Word(),
		Token:    faker.Word(),
	}

	t.GroupDto = &proto.Group{
		Id:       t.Group.ID.String(),
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
	}

	t.CreateGroupReqMock = &proto.CreateGroupRequest{
		Group: &proto.Group{
			LeaderID: t.Group.LeaderID,
			Token:    t.Group.Token,
		},
	}

	t.UpdateGroupReqMock = &proto.UpdateGroupRequest{
		Group: &proto.Group{
			Id:       t.Group.ID.String(),
			LeaderID: t.Group.LeaderID,
			Token:    t.Group.Token,
		},
	}
}

func (t *GroupServiceTest) TestFindOneSuccess() {
	want := &proto.FindOneGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	repo.On("FindOne", t.Group.ID.String(), &group.Group{}).Return(t.Group, nil)

	srv := NewService(repo)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneGroupRequest{Id: t.Group.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindOneNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindOne", t.Group.ID.String(), &group.Group{}).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneGroupRequest{Id: t.Group.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestCreateSuccess() {
	want := &proto.CreateGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	in := &group.Group{
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
	}

	repo.On("Create", in).Return(t.Group, nil)

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateGroupReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestCreateInternalErr() {
	repo := &mock.RepositoryMock{}

	in := &group.Group{
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
	}

	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *GroupServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	repo.On("Update", t.Group.ID.String(), t.Group).Return(t.Group, nil)

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestUpdateNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.Group.ID.String(), t.Group).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestUpdateMalformed() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.Group.ID.String(), t.Group).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)

	t.UpdateGroupReqMock.Group.Id = "abc"

	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestDeleteSuccess() {
	want := &proto.DeleteGroupResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.Group.ID.String()).Return(nil)

	srv := NewService(repo)
	actual, err := srv.Delete(context.Background(), &proto.DeleteGroupRequest{Id: t.GroupDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestDeleteNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.Group.ID.String()).Return(errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.Delete(context.Background(), &proto.DeleteGroupRequest{Id: t.GroupDto.Id})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestFindByTokenSuccess() {
	want := &proto.FindByTokenGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	repo.On("FindByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)

	srv := NewService(repo)
	actual, err := srv.FindByToken(context.Background(), &proto.FindByTokenGroupRequest{Token: t.GroupDto.Token})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindByTokenNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindByToken", t.Group.Token, &group.Group{}).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.FindByToken(context.Background(), &proto.FindByTokenGroupRequest{Token: t.GroupDto.Token})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

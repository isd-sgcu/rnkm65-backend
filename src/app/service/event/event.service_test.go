package event

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/event"
	fMock "github.com/isd-sgcu/rnkm65-backend/src/mocks/file"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type EventServiceTest struct {
	suite.Suite
	Event              *event.Event
	EventDto           *proto.Event
	CreateEventReqMock *proto.CreateEventRequest
	UpdateEventReqMock *proto.UpdateEventRequest
}

func TestEventService(t *testing.T) {
	suite.Run(t, new(EventServiceTest))
}

func (t *EventServiceTest) SetupTest() {
	t.Event = &event.Event{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		NameTH:        faker.NameTH(),
		DescriptionTH: faker.DescriptionTH(),
		NameEN:        faker.NameEN,
		DescriptionEN: faker.DescriptionEN,
		Code:          faker.Code,
	}

	t.EventDto = &proto.Event{
		Id:            t.Event.ID.String(),
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
	}

	t.CreateEventReqMock = &proto.CreateEventRequest{
		Event: &proto.Event{
			Title:           t.Event.Title,
			Firstname:       t.Event.Firstname,
			Lastname:        t.Event.Lastname,
			Nickname:        t.Event.Nickname,
			StudentID:       t.Event.StudentID,
			Faculty:         t.Event.Faculty,
			Year:            t.Event.Year,
			Phone:           t.Event.Phone,
			LineID:          t.Event.LineID,
			Email:           t.Event.Email,
			AllergyFood:     t.Event.AllergyFood,
			FoodRestriction: t.Event.FoodRestriction,
			AllergyMedicine: t.Event.AllergyMedicine,
			Disease:         t.Event.Disease,
			CanSelectBaan:   *t.Event.CanSelectBaan,
			IsVerify:        *t.Event.IsVerify,
			GroupId:         t.Event.GroupID.String(),
		},
	}

	t.UpdateEventReqMock = &proto.UpdateEventRequest{
		Event: &proto.Event{
			Id:              t.Event.ID.String(),
			Title:           t.Event.Title,
			Firstname:       t.Event.Firstname,
			Lastname:        t.Event.Lastname,
			Nickname:        t.Event.Nickname,
			StudentID:       t.Event.StudentID,
			Faculty:         t.Event.Faculty,
			Year:            t.Event.Year,
			Phone:           t.Event.Phone,
			LineID:          t.Event.LineID,
			Email:           t.Event.Email,
			AllergyFood:     t.Event.AllergyFood,
			FoodRestriction: t.Event.FoodRestriction,
			AllergyMedicine: t.Event.AllergyMedicine,
			Disease:         t.Event.Disease,
			CanSelectBaan:   *t.Event.CanSelectBaan,
			IsVerify:        *t.Event.IsVerify,
			GroupId:         t.Event.GroupID.String(),
		},
	}
}

func (t *EventServiceTest) TestFindOneSuccess() {
	url := faker.URL()

	t.EventDto.ImageUrl = url

	want := &proto.FindOneEventResponse{Event: t.EventDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.Event.ID.String(), &event.Event{}).Return(t.Event, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.Event.ID.String()).Return(url, nil)

	srv := NewService(repo, fileSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneEventRequest{Id: t.Event.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestFindOneNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.Event.ID.String(), &event.Event{}).Return(nil, errors.New("Not found event"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneEventRequest{Id: t.Event.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *EventServiceTest) TestCreateSuccess() {
	want := &proto.CreateEventResponse{Event: t.EventDto}

	repo := &mock.RepositoryMock{}

	in := &event.Event{
		Title:           t.Event.Title,
		Firstname:       t.Event.Firstname,
		Lastname:        t.Event.Lastname,
		Nickname:        t.Event.Nickname,
		StudentID:       t.Event.StudentID,
		Faculty:         t.Event.Faculty,
		Year:            t.Event.Year,
		Phone:           t.Event.Phone,
		LineID:          t.Event.LineID,
		Email:           t.Event.Email,
		AllergyFood:     t.Event.AllergyFood,
		FoodRestriction: t.Event.FoodRestriction,
		AllergyMedicine: t.Event.AllergyMedicine,
		Disease:         t.Event.Disease,
		CanSelectBaan:   t.Event.CanSelectBaan,
		GroupID:         t.Event.GroupID,
	}

	repo.On("Create", in).Return(t.Event, nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	actual, err := srv.Create(context.Background(), t.CreateEventReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestCreateInternalErr() {
	repo := &mock.RepositoryMock{}

	in := &event.Event{
		Title:           t.Event.Title,
		Firstname:       t.Event.Firstname,
		Lastname:        t.Event.Lastname,
		Nickname:        t.Event.Nickname,
		StudentID:       t.Event.StudentID,
		Faculty:         t.Event.Faculty,
		Year:            t.Event.Year,
		Phone:           t.Event.Phone,
		LineID:          t.Event.LineID,
		Email:           t.Event.Email,
		AllergyFood:     t.Event.AllergyFood,
		FoodRestriction: t.Event.FoodRestriction,
		AllergyMedicine: t.Event.AllergyMedicine,
		Disease:         t.Event.Disease,
		CanSelectBaan:   t.Event.CanSelectBaan,
		GroupID:         t.Event.GroupID,
	}

	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	actual, err := srv.Create(context.Background(), t.CreateEventReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *EventServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateEventResponse{Event: t.EventDto}

	eventIn := *t.Event
	eventIn.IsVerify = nil

	repo := &mock.RepositoryMock{}

	repo.On("Update", t.Event.ID.String(), &eventIn).Return(t.Event, nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Update(context.Background(), t.UpdateEventReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestUpdateNotFound() {
	eventIn := *t.Event
	eventIn.IsVerify = nil

	repo := &mock.RepositoryMock{}
	repo.On("Update", t.Event.ID.String(), &eventIn).Return(nil, errors.New("Not found event"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Update(context.Background(), t.UpdateEventReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *EventServiceTest) TestUpdateMalformed() {
	eventIn := *t.Event
	eventIn.IsVerify = nil

	repo := &mock.RepositoryMock{}
	repo.On("Update", t.Event.ID.String(), eventIn).Return(nil, errors.New("Not found event"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	t.UpdateEventReqMock.Event.Id = "abc"

	actual, err := srv.Update(context.Background(), t.UpdateEventReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *EventServiceTest) TestDeleteSuccess() {
	want := &proto.DeleteEventResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.Event.ID.String()).Return(nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Delete(context.Background(), &proto.DeleteEventRequest{Id: t.EventDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestDeleteNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.Event.ID.String()).Return(errors.New("Not found event"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Delete(context.Background(), &proto.DeleteEventRequest{Id: t.EventDto.Id})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

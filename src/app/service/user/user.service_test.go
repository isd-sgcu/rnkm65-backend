package user

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	fMock "github.com/isd-sgcu/rnkm65-backend/src/mocks/file"
	mockFile "github.com/isd-sgcu/rnkm65-backend/src/mocks/file"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/user"
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

type UserServiceTest struct {
	suite.Suite
	User              *user.User
	UpdateUser        *user.User
	UserDto           *proto.User
	Url               string
	CreateUserReqMock *proto.CreateUserRequest
	UpdateUserReqMock *proto.UpdateUserRequest
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.User = &user.User{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Title:           faker.Word(),
		Firstname:       faker.FirstName(),
		Lastname:        faker.LastName(),
		Nickname:        faker.Name(),
		StudentID:       faker.Word(),
		Faculty:         faker.Word(),
		Year:            faker.Word(),
		Phone:           faker.Phonenumber(),
		LineID:          faker.Word(),
		Email:           faker.Email(),
		AllergyFood:     faker.Word(),
		FoodRestriction: faker.Word(),
		AllergyMedicine: faker.Word(),
		Disease:         faker.Word(),
		GroupID:         utils.UUIDAdr(uuid.New()),
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
	}

	t.UserDto = &proto.User{
		Id:              t.User.ID.String(),
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		CanSelectBaan:   *t.User.CanSelectBaan,
		IsVerify:        *t.User.IsVerify,
	}

	t.CreateUserReqMock = &proto.CreateUserRequest{
		User: &proto.User{
			Title:           t.User.Title,
			Firstname:       t.User.Firstname,
			Lastname:        t.User.Lastname,
			Nickname:        t.User.Nickname,
			StudentID:       t.User.StudentID,
			Faculty:         t.User.Faculty,
			Year:            t.User.Year,
			Phone:           t.User.Phone,
			LineID:          t.User.LineID,
			Email:           t.User.Email,
			AllergyFood:     t.User.AllergyFood,
			FoodRestriction: t.User.FoodRestriction,
			AllergyMedicine: t.User.AllergyMedicine,
			Disease:         t.User.Disease,
			CanSelectBaan:   *t.User.CanSelectBaan,
			IsVerify:        *t.User.IsVerify,
			GroupId:         t.User.GroupID.String(),
		},
	}

	t.UpdateUserReqMock = &proto.UpdateUserRequest{
		Id:              t.User.ID.String(),
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
	}

	t.UpdateUser = &user.User{
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
	}

	t.Url = faker.URL()
}

func createUrlList() []string {
	var result []string
	for i := 0; i < 3; i++ {
		result = append(result, faker.URL())
	}

	return result
}

func createUserList() []*user.User {
	var result []*user.User
	for i := 0; i < 3; i++ {
		usr := user.User{
			Base: model.Base{
				ID: uuid.New(),
			},
			Title:           faker.Word(),
			Firstname:       faker.FirstName(),
			Lastname:        faker.LastName(),
			Nickname:        faker.Name(),
			StudentID:       faker.Word(),
			Faculty:         faker.Word(),
			Year:            faker.Word(),
			Phone:           faker.Phonenumber(),
			LineID:          faker.Word(),
			Email:           faker.Email(),
			AllergyFood:     faker.Word(),
			FoodRestriction: faker.Word(),
			AllergyMedicine: faker.Word(),
			Disease:         faker.Word(),
			CanSelectBaan:   utils.BoolAdr(true),
			IsVerify:        utils.BoolAdr(true),
		}

		result = append(result, &usr)
	}
	return result
}

func createUserDto(userList []*user.User, urlList []string) []*proto.User {
	var result []*proto.User

	for i, u := range userList {
		usr := proto.User{
			Id:              u.ID.String(),
			Title:           u.Title,
			Firstname:       u.Firstname,
			Lastname:        u.Lastname,
			Nickname:        u.Nickname,
			StudentID:       u.StudentID,
			Faculty:         u.Faculty,
			Year:            u.Year,
			Phone:           u.Phone,
			LineID:          u.LineID,
			Email:           u.Email,
			AllergyFood:     u.AllergyFood,
			FoodRestriction: u.FoodRestriction,
			AllergyMedicine: u.AllergyMedicine,
			Disease:         u.Disease,
			ImageUrl:        urlList[i],
			CanSelectBaan:   *u.CanSelectBaan,
			IsVerify:        *u.IsVerify,
		}

		result = append(result, &usr)
	}

	return result
}

func (t *UserServiceTest) TestFindOneSuccess() {
	url := faker.URL()

	t.UserDto.ImageUrl = url

	want := &proto.FindOneUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return(url, nil)

	srv := NewService(repo, fileSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneSignUrlErr() {
	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return("", status.Error(codes.Unavailable, "Cannot get signed url"))

	srv := NewService(repo, fileSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Unavailable, st.Code())
}

func (t *UserServiceTest) TestFindOneSignUrlNotFound() {
	want := &proto.FindOneUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return("", status.Error(codes.NotFound, "Not found file"))

	srv := NewService(repo, fileSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(nil, gorm.ErrRecordNotFound)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestCreateSuccess() {
	want := &proto.CreateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	in := &user.User{
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		CanSelectBaan:   t.User.CanSelectBaan,
		GroupID:         t.User.GroupID,
	}

	repo.On("Create", in).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	actual, err := srv.Create(context.Background(), t.CreateUserReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestCreateInternalErr() {
	repo := &mock.RepositoryMock{}

	in := &user.User{
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		CanSelectBaan:   t.User.CanSelectBaan,
		GroupID:         t.User.GroupID,
	}

	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)

	actual, err := srv.Create(context.Background(), t.CreateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *UserServiceTest) TestVerifySuccess() {
	want := &proto.VerifyUserResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Verify", t.User.ID.String()).Return(nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Verify(context.Background(), &proto.VerifyUserRequest{StudentId: t.UserDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestVerifyNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Verify", t.User.ID.String()).Return(gorm.ErrRecordNotFound)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Verify(context.Background(), &proto.VerifyUserRequest{StudentId: t.UserDto.Id})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("Update", t.User.ID.String(), t.UpdateUser).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestUpdateNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.User.ID.String(), t.UpdateUser).Return(nil, errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestDeleteSuccess() {
	want := &proto.DeleteUserResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.User.ID.String()).Return(nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Delete(context.Background(), &proto.DeleteUserRequest{Id: t.UserDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestDeleteNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.User.ID.String()).Return(errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.Delete(context.Background(), &proto.DeleteUserRequest{Id: t.UserDto.Id})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestCreateOrUpdateSuccess() {
	userIn := *t.User
	userIn.IsVerify = nil
	userIn.GroupID = nil
	want := &proto.CreateOrUpdateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", &userIn).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestCreateOrUpdateMalformedID() {
	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", t.User).Return(t.User, nil)

	t.UserDto.Id = "abc"

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *UserServiceTest) TestCreateOrUpdateInternalErr() {
	userIn := *t.User
	userIn.IsVerify = nil
	userIn.GroupID = nil
	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", &userIn).Return(nil, errors.New("Something wrong"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *UserServiceTest) TestFindByStudentIDSuccess() {
	want := &proto.FindByStudentIDUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("FindByStudentID", t.User.StudentID, &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.FindByStudentID(context.Background(), &proto.FindByStudentIDUserRequest{StudentId: t.UserDto.StudentID})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindByStudentIDNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindByStudentID", t.User.StudentID, &user.User{}).Return(nil, errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.FindByStudentID(context.Background(), &proto.FindByStudentIDUserRequest{StudentId: t.UserDto.StudentID})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestGetImageSuccess() {
	want := t.Url

	repo := &mock.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return(t.Url, nil)

	srv := NewService(repo, fileSrv)
	actual, err := srv.GetUserImage(t.User.ID.String())

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestGetImageNotFound() {
	repo := &mock.RepositoryMock{}

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return(nil, status.Error(codes.NotFound, "Not found file"))

	srv := NewService(repo, fileSrv)
	actual, err := srv.GetUserImage(t.User.ID.String())

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Equal(t.T(), "", actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestGetImagesSuccess() {
	want := createUrlList()
	userList := createUserList()

	repo := &mock.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}
	for i, u := range userList {
		fileSrv.On("GetSignedUrl", u.ID.String()).Return(want[i], nil)
	}

	srv := NewService(repo, fileSrv)
	actual, err := srv.GetUserImages(userList)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestRawToDto() {
	want := &proto.User{
		Id:              t.User.ID.String(),
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		CanSelectBaan:   *t.User.CanSelectBaan,
		IsVerify:        *t.User.IsVerify,
		ImageUrl:        t.Url,
	}

	repo := &mock.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual := srv.RawToDto(t.User, t.Url)

	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestDtoToRaw() {
	want := &user.User{
		Base: model.Base{
			ID: t.User.ID,
		},
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		CanSelectBaan:   t.User.CanSelectBaan,
	}

	repo := &mock.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	srv := NewService(repo, fileSrv)
	actual, err := srv.DtoToRaw(t.UserDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

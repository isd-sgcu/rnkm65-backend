package baan

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	constant "github.com/isd-sgcu/rnkm65-backend/src/constant/baan"
	size "github.com/isd-sgcu/rnkm65-backend/src/constant/baan"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/baan"
	cacheMock "github.com/isd-sgcu/rnkm65-backend/src/mocks/cache"
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

type BaanServiceTest struct {
	suite.Suite
	baans   []*baan.Baan
	baan    *baan.Baan
	baanDto *proto.Baan
	conf    config.App
}

func TestBaanService(t *testing.T) {
	suite.Run(t, new(BaanServiceTest))
}

func (t *BaanServiceTest) SetupTest() {
	t.baans = make([]*baan.Baan, 0)

	t.baan = &baan.Baan{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		NameTH:        faker.Word(),
		DescriptionTH: faker.Paragraph(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Paragraph(),
		Size:          size.M,
		Facebook:      faker.URL(),
		Instagram:     faker.URL(),
		Line:          faker.URL(),
		ImageUrl:      faker.URL(),
	}

	baan2 := &baan.Baan{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		NameTH:        faker.Word(),
		DescriptionTH: faker.Paragraph(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Paragraph(),
		Size:          size.M,
		Facebook:      faker.URL(),
		Instagram:     faker.URL(),
		Line:          faker.URL(),
		ImageUrl:      faker.URL(),
	}

	baan3 := &baan.Baan{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		NameTH:        faker.Word(),
		DescriptionTH: faker.Paragraph(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Paragraph(),
		Size:          size.M,
		Facebook:      faker.URL(),
		Instagram:     faker.URL(),
		Line:          faker.URL(),
		ImageUrl:      faker.URL(),
	}

	t.baanDto = &proto.Baan{
		Id:            t.baan.ID.String(),
		NameTH:        t.baan.NameTH,
		DescriptionTH: t.baan.DescriptionTH,
		NameEN:        t.baan.NameEN,
		DescriptionEN: t.baan.DescriptionEN,
		Size:          proto.BaanSize(t.baan.Size),
		Facebook:      t.baan.Facebook,
		Instagram:     t.baan.Instagram,
		Line:          t.baan.Line,
		ImageUrl:      t.baan.ImageUrl,
	}

	t.baans = append(t.baans, t.baan, baan2, baan3)
	t.conf = config.App{
		Port:         3001,
		Debug:        false,
		BaanCacheTTL: 15,
	}
}

func (t *BaanServiceTest) TestGetAllBaanSuccess() {
	want := &proto.GetAllBaanResponse{Baans: createBaanDto(t.baans)}

	var baansIn []*baan.Baan

	r := mock.RepositoryMock{}
	r.On("FindAll", baansIn).Return(&t.baans, nil)

	c := cacheMock.RepositoryMock{
		V: map[string]interface{}{},
	}
	c.On("SaveCache", constant.BaanKey, &t.baans, t.conf.BaanCacheTTL).Return(nil)
	c.On("GetCache", constant.BaanKey, &baansIn).Return(&t.baans, nil)

	srv := NewService(&r, &c, t.conf)
	actual, err := srv.GetAllBaan(context.Background(), &proto.GetAllBaanRequest{})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func createBaanDto(in []*baan.Baan) []*proto.Baan {
	var result []*proto.Baan

	for _, b := range in {
		r := &proto.Baan{
			Id:            b.ID.String(),
			NameTH:        b.NameTH,
			DescriptionTH: b.DescriptionTH,
			NameEN:        b.NameEN,
			DescriptionEN: b.DescriptionEN,
			Size:          proto.BaanSize(b.Size),
			Facebook:      b.Facebook,
			Instagram:     b.Instagram,
			Line:          b.Line,
			ImageUrl:      b.ImageUrl,
		}

		result = append(result, r)
	}

	return result
}

func (t *BaanServiceTest) TestGetBaanSuccess() {
	want := &proto.GetBaanResponse{Baan: t.baanDto}

	r := mock.RepositoryMock{}
	r.On("FindOne", t.baan.ID.String(), &baan.Baan{}).Return(t.baan, nil)

	c := cacheMock.RepositoryMock{
		V: map[string]interface{}{},
	}
	c.On("GetCache", t.baan.ID.String(), &baan.Baan{}).Return(t.baan, nil)

	srv := NewService(&r, &c, t.conf)
	actual, err := srv.GetBaan(context.Background(), &proto.GetBaanRequest{Id: t.baan.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *BaanServiceTest) TestGetBaanSuccessNotCache() {
	want := &proto.GetBaanResponse{Baan: t.baanDto}

	r := mock.RepositoryMock{}
	r.On("FindOne", t.baan.ID.String(), &baan.Baan{}).Return(t.baan, nil)

	c := cacheMock.RepositoryMock{
		V: map[string]interface{}{},
	}
	c.On("GetCache", t.baan.ID.String(), &baan.Baan{}).Return(nil, redis.Nil)
	c.On("SaveCache", t.baan.ID.String(), t.baan, t.conf.BaanCacheTTL).Return(nil)

	srv := NewService(&r, &c, t.conf)
	actual, err := srv.GetBaan(context.Background(), &proto.GetBaanRequest{Id: t.baan.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *BaanServiceTest) TestGetBaanNotFound() {
	r := mock.RepositoryMock{}
	r.On("FindOne", t.baan.ID.String(), &baan.Baan{}).Return(nil, errors.New("Not found baan"))

	c := cacheMock.RepositoryMock{
		V: map[string]interface{}{},
	}
	c.On("GetCache", t.baan.ID.String(), &baan.Baan{}).Return(nil, redis.Nil)

	srv := NewService(&r, &c, t.conf)
	actual, err := srv.GetBaan(context.Background(), &proto.GetBaanRequest{Id: t.baan.ID.String()})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

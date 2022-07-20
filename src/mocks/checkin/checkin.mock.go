package checkin

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Checkin(ci *checkin.Checkin) error {
	args := r.Called(ci)

	return args.Error(0)
}

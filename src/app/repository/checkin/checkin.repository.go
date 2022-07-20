package checkin

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Checkin(ci *checkin.Checkin) error {
	return r.db.Create(ci).Error
}

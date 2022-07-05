package group

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindOne(id string, result *group.Group) error {
	return r.db.First(&result, "id = ?", id).Error
}

func (r *Repository) FindByToken(token string, result *group.Group) error {
	return r.db.First(&result, "token = ?", token).Error
}

func (r *Repository) Create(in *group.Group) error {
	return r.db.Create(&in).Error
}

func (r *Repository) Update(id string, result group.Group) error {
	return r.db.Where("id = ?", id).Updates(&result).First(&result).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.First(id).Delete(&group.Group{}).Error
}

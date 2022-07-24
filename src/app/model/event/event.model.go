package event

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"gorm.io/gorm"
)

type Event struct {
	model.Base
	NameTH        string `json:"name_th"`
	DescriptionTH string `json:"description_th"`
	NameEN        string `json:"name_en"`
	DescriptionEN string `json:"description_en"`
	Code          string `json:"code"`
	// Users         []model.Base.user.User `json:"users" gorm:"many2many:event_user"`
}

func (e *Event) BeforeCreate(_ *gorm.DB) error {
	e.ID = uuid.New()

	return nil
}

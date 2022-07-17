package baan

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/constant/baan"
)

type Baan struct {
	model.Base
	NameTH        string        `json:"name_th" gorm:"type:tinytext"`
	DescriptionTH string        `json:"description_th" gorm:"type:mediumtext"`
	NameEN        string        `json:"name_en" gorm:"type:tinytext"`
	DescriptionEN string        `json:"description_en" gorm:"type:mediumtext"`
	Size          baan.BaanSize `json:"size" gorm:"type:tinyint"`
	Facebook      string        `json:"facebook" gorm:"type:tinytext"`
	Instagram     string        `json:"instagram" gorm:"type:tinytext"`
	Line          string        `json:"line" gorm:"type:tinytext"`
	Members       []*user.User  `json:"members"`
}

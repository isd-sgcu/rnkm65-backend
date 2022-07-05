package group

import "github.com/isd-sgcu/rnkm65-backend/src/app/model"

type Group struct {
	model.Base
	LeaderID string `json:"leader_id"`
	Token    string `json:"token" gorm:"type:tinyText"`
}

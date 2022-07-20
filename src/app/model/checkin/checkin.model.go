package checkin

import (
	"time"

	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
)

type Checkin struct {
	model.Base
	UserId     string     `json:"user_id" gorm:"index"`
	CheckinAt  *time.Time `json:"check_in_at" gorm:"type:datetime;autoCreateTime:nano"`
	CheckoutAt *time.Time `json:"check_out_at" gorm:"type:datetime"`
	EventType  int32      `json:"event_type"`
}

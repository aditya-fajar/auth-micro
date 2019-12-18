package model

import "time"

type AccessToken struct {
	ID      uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Token   string    `json:"token" gorm:"column:token"`
	UserID  uint64    `json:"user_id" gorm:"column:user_id"`
	Expired time.Time `json:"expired"   gorm:"column:expired_at"`
}

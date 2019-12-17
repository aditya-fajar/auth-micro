package model

type Token struct {
	ID      uint64 `json:"id" gorm:"primary_key"`
	Token   string `json:"token" gorm:"column:token"`
	Expired string `json:"expired"   gorm:"column:expired_at"`
}

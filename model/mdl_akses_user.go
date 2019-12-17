package model

type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password"   gorm:"column:password"`
}

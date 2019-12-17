package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pfpmaniac123/auth-micro/model"
)

type User struct {
	DB *gorm.DB
}

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func New(db *gorm.DB) User {
// 	return User{DB: db}
// }

func (u *User) Register(c *gin.Context) {
	var request model.User

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	u.DB.Create(&request)
	// fmt.Println(request)
	c.JSON(200, gin.H{
		"message": request,
	})
}

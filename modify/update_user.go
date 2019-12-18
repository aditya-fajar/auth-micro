package modify

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pfpmaniac123/auth-micro/model"
)

type User struct {
	DB *gorm.DB
}

//Test

type user struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetUsers(c *gin.Context) {
	db := u.DB
	var user []model.User

	db.Find(&user)

	c.JSON(200, gin.H{
		"data": user,
	})
}

func (u *User) UpdateUser(c *gin.Context) {
	var request user

	id := c.Param("id")

	var user model.User

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	u.DB.Where("id = ?", id).Find(&user)

	user.Name = request.Name
	user.Email = request.Email

	u.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
		"user":    user,
	})
}

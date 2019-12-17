package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Login struct {
	DB *gorm.DB
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(DB *gorm.DB) Login {
	return Login{DB: DB}
}

func (l Login) LoginUser(c *gin.Context) {
	var request LoginBody

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	l.DB.Where("name = ?", "jinzhu").First(&DB)
	//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	//p.DB.Create(&request)
	c.JSON(200, gin.H{
		"message": request,
	})

}

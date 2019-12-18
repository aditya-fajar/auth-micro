package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pfpmaniac123/auth-micro/model"

	jwt "github.com/dgrijalva/jwt-go"
)

type Login struct {
	DB *gorm.DB
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyClaims struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}

var APPLICATION_NAME = "Auth-Micro"

// var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Day
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("auth-micro123")

func New(DB *gorm.DB) Login {
	return Login{DB: DB}
}

func (l Login) LoginUser(c *gin.Context) {
	var request LoginBody
	var user model.User
	var accessToken model.AccessToken

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	l.DB.Where("email = ? AND password = ?", request.Email, request.Password).Find(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	expired := time.Now().Add(time.Hour * 48)

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expired.Unix(),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken.Token = signedToken
	accessToken.UserID = user.ID
	accessToken.Expired = expired

	l.DB.Create(&accessToken)
	c.JSON(200, gin.H{
		"data":   user,
		"token":  signedToken,
		"access": accessToken,
	})

}

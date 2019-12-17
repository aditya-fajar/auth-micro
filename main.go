package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pfpmaniac123/auth-micro/auth"
	"github.com/pfpmaniac123/auth-micro/config"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	login := auth.New(db)
	register := auth.User{DB: db}
	//product := product.Product{DB: db}

	//r.GET("/products", product.GetProducts)
	//r.POST("/products", product.CreateProduct)

	//r.GET("/test", func)
	//Menu Login (AgungW)
	r.POST("/login", login.LoginUser)

	r.POST("/register", register.Register)

	r.PATCH("/update", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update",
		})
	})

	r.Run()
}

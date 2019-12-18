package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pfpmaniac123/auth-micro/config"
	"github.com/pfpmaniac123/auth-micro/modify"
	"github.com/pfpmaniac123/auth-micro/auth"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	login := auth.New(db)
	register := auth.User{DB: db}
	//login := auth.New(db)
	update := modify.User{DB: db}
	//product := product.Product{DB: db}

	//r.GET("/products", product.GetProducts)
	//r.POST("/products", product.CreateProduct)

	//r.GET("/test", func)
	//Menu Login (AgungW)
	r.POST("/login", login.LoginUser)
	//r.POST("/login", login.LoginUser)

	r.POST("/register", register.Register)

	r.PATCH("/update", update.GetUsers)

	r.Run()
}

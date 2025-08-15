package main

import (
	"go-micro/controller"
	db "go-micro/db/connection"
	"go-micro/model"
	"go-micro/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatalln("cannot connet DB: ", err)
		return
	}

	userModel := model.NewAuthRepository(db)
	userService := service.NewAuthService(userModel)
	userController := controller.NewAuthController(userService)

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/Register", userController.Register)
		auth.POST("/Login", userController.Login)
	}
	r.Run(":8001")
}

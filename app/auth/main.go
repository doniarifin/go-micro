package main

import (
	rabbitmq "go-micro/config"
	"go-micro/controller"
	db "go-micro/db/connection"
	"go-micro/model"
	"go-micro/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbconn, err := db.ConnectDB()
	if err != nil {
		log.Fatalln("cannot connet DB: ", err)
		return
	}
	defer db.CloseConnection()

	if err := rabbitmq.Connect(); err != nil {
		return
	}
	defer rabbitmq.CloseConnection()

	userModel := model.NewAuthRepository(dbconn)
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

package main

import (
	rabbitmq "go-micro/config"
	"go-micro/controller"
	db "go-micro/db/connection"
	"go-micro/middleware"
	"go-micro/model"
	"go-micro/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := rabbitmq.Connect(); err != nil {
		return
	}

	defer rabbitmq.CloseConnection()

	go func() {
		if err := rabbitmq.Consume("email", db); err != nil {
			return
		}
	}()

	msg := model.NewMsgRepository(db)
	msgService := service.NewMsgService(msg)
	msgController := controller.NewMsgController(msgService)

	r := gin.Default()
	app := r.Group("/msg")

	app.Use(middleware.JWTMiddleware("admin"))

	{
		app.POST("/save", msgController.Save)
		app.POST("/gets/:id", msgController.Get)
		app.POST("/delete", msgController.Delete)
	}

	r.Run(":8002")
}

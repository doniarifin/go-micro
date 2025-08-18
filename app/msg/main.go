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
	dbConn, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.CloseConnection()

	if err := rabbitmq.Connect(); err != nil {
		return
	}

	defer rabbitmq.CloseConnection()

	go func() {
		if err := rabbitmq.Consume("email", dbConn); err != nil {
			log.Println("Error consuming RabbitMQ:", err)
			return
		}
	}()

	msg := model.NewMsgRepository(dbConn)
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

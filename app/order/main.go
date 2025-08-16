package main

import (
	"go-micro/controller"
	db "go-micro/db/connection"
	"go-micro/model"
	"go-micro/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := db.ConnectDB()

	if err != nil {
		return
	}

	r := gin.Default()
	app := r.Group("/order")

	order := model.NewOrderRepository(db)
	orderService := service.NewOrderService(order)
	orderController := controller.NewOrderController(orderService)

	{
		app.POST("/Create", orderController.Create)
	}
	r.Run(":8003")
}

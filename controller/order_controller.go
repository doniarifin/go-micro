package controller

import (
	"go-micro/model"
	"go-micro/service"
	"time"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) *orderController {
	return &orderController{service}
}

type OrderRequest struct {
	Model *model.Order
}

func (controller *orderController) Create(c *gin.Context) {

	payload := OrderRequest{}

	if err := c.ShouldBindJSON(&payload.Model); err != nil {
		return
	}

	order := &model.Order{}
	order = payload.Model
	order.Date = time.Now()
	order.CreatedAt = time.Now()

	if err := controller.service.Create(order); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    order,
	})
}
func (controller *orderController) Get(id string, c *gin.Context) {
	Id := c.Param(id)
	order, err := controller.service.Get(&model.Order{ID: Id})

	if err != nil {
		return
	}

	c.JSON(200, &order)
}
func (controller *orderController) Update(c *gin.Context) {
	payload := OrderRequest{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		return
	}

	order := &model.Order{}
	order = payload.Model

	if err := controller.service.Create(order); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    order,
	})
}
func (controller *orderController) Delete(id string, c *gin.Context) {
	Id := c.Param(id)

	order := &model.Order{}
	order.ID = Id

	err := controller.service.Delete(order)

	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "success order delete",
	})
}

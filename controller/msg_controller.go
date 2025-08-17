package controller

import (
	"go-micro/model"
	"go-micro/service"

	"github.com/gin-gonic/gin"
)

type msgController struct {
	service service.MsgService
}

func NewMsgController(srv service.MsgService) *msgController {
	return &msgController{srv}
}

type MsgRequest struct {
	model *model.Message
}

func (ctrl *msgController) Save(c *gin.Context) {
	payload := &MsgRequest{}

	if err := c.ShouldBindJSON(payload); err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	msg := &model.Message{}
	msg = payload.model

	if err := ctrl.service.Insert(msg); err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    msg,
	})
}

func (ctrl *msgController) Gets(c *gin.Context) {
	ID := c.Param("id")

	data, err := ctrl.service.Gets(&model.Message{ID: ID})

	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (ctrl *msgController) Delete(c *gin.Context) {
	ID := c.Param("id")
	if err := ctrl.service.Delete(&model.Message{ID: ID}); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "success delete msg",
	})
}

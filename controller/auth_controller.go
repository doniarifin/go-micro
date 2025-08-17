package controller

import (
	"encoding/json"
	rabbitmq "go-micro/config"
	"go-micro/model"
	"go-micro/service"
	helper "go-micro/utils"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type authController struct {
	auth service.AuthService
}

func NewAuthController(auth service.AuthService) *authController {
	return &authController{auth}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	// Token   string `json:"token"`
}

func (ac *authController) Register(c *gin.Context) {
	payload := &RegisterRequest{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		return
	}

	//hash password
	pass, err := helper.HashPassword(payload.Password)
	if err != nil {
		log.Println("Hash password error:", err)
		c.JSON(500, gin.H{"error": "failed register user"})
		return
	}

	user := &model.AuthUser{}
	user.ID = helper.NewUUID()
	user.Email = payload.Email
	user.Password = pass

	if err := ac.auth.Register(user); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Println("register error:", err)
			c.JSON(500, gin.H{"error": "email already exist"})
			return
		}
		log.Println("DB error:", err)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	msg := &model.Message{}
	msg.Sender = "system"
	msg.Receiver = user.Email
	msg.MsgType = "email"
	msg.MsgBody = "created account success"
	msg.CreatedBy = user.Email
	msg.CreatedAt = time.Now()
	msg.UpdateAt = time.Now()

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return
	}

	if err := rabbitmq.Publish("email", string(msgJSON)); err != nil {
		return
	}

	c.JSON(200, &RegisterResponse{
		Message: "success",
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (ac *authController) Login(c *gin.Context) {
	payload := &LoginRequest{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		return
	}

	user := &model.AuthUser{}
	user.Email = payload.Email
	user.Password = payload.Password

	userLogin, err := ac.auth.Login(user)

	if err != nil {
		log.Println("DB error:", err)
		c.JSON(401, gin.H{"message": "email not found"})
		return
	}

	if !helper.CheckHashPassword(userLogin.Password, payload.Password) {
		c.JSON(401, gin.H{"message": "wrong password"})
		return
	}

	regReq := &helper.AuthUser{
		ID:    user.ID,
		Email: user.Email,
	}

	token, err := helper.GenerateToken(regReq)

	if err != nil {
		log.Println("error:", err)
		c.JSON(500, gin.H{"error": "failed generate token"})
		return
	}

	c.JSON(200, &LoginResponse{
		Message: "success",
		Token:   token,
	})
}

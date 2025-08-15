package controller

import (
	"go-micro/model"
	"go-micro/service"
	helper "go-micro/utils"
	"log"

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

	user := &model.AuthUser{}
	user.ID = helper.NewUUID()
	user.Email = payload.Email
	user.Password = payload.Password

	err := ac.auth.Register(user)

	if err != nil {
		log.Println("DB error:", err)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	// regReq := &helper.AuthUser{
	// 	ID:    user.ID,
	// 	Email: user.Email,
	// }

	// token, err := helper.GenerateToken(regReq)

	// if err != nil {
	// 	log.Println("error:", err)
	// 	c.JSON(500, gin.H{"error": "failed generate token"})
	// 	return
	// }

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

	_, err := ac.auth.Login(user)

	if err != nil {
		log.Println("DB error:", err)
		c.JSON(500, gin.H{"error": "internal server error"})
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

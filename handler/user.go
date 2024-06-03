package handler

import (
	"kilobite/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//Tangkap input dari user
	//Map input dari user ke struct RegisterUserInput
	//struct di atas kita parsing sebagai parameter service
}

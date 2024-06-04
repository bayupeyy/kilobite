package handler

import (
	"kilobite/helper"
	"kilobite/user"
	"net/http"

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

	var input user.RegisterUserInput //Membuat variabel untuk menangkap input dari user

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	//Dibawah ini untuk mengubah format response
	//Jika terdapat Token maka membutuhkan perinta token, err := h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "TOKENTOKENTOKENTOKEN")

	response := helper.APIResponse("Akun berhasil registrasi", http.StatusOK, "success", formatter)

	//Untuk kembalikan Status
	c.JSON(http.StatusOK, response)
}

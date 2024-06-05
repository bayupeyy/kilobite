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

		errors := helper.FormatValidationError(err) //Fungsi untuk mengatur format error dari helper.go
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage) //kode ini berguna untuk menampilkan format error
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//Dibawah ini untuk mengubah format response
	//Jika terdapat Token maka membutuhkan perinta token, err := h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "TOKENTOKENTOKENTOKEN")

	response := helper.APIResponse("Akun berhasil registrasi", http.StatusOK, "success", formatter)

	//Untuk kembalikan Status
	c.JSON(http.StatusOK, response)
}

// Func untuk login user
func (h *userHandler) Login(c *gin.Context) {
	//User memasukkan input email & password
	//Input ditangkap handler
	//Mapping dari input user ke input struct
	//input struct passing service
	//di service mencari dg bantuan repository user dengan email x
	//Mencocokan password
}

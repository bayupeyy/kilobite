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

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage) //kode ini berguna untuk menampilkan format error
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	loggedinUser, err := h.userService.Login(input)

	// Pengecekan error
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage) //kode ini berguna untuk menampilkan format error
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokenentetoken")

	response := helper.APIResponse("Berhasil Masuk", http.StatusOK, "success", formatter)

	//Untuk kembalikan Status
	c.JSON(http.StatusOK, response)

}
func (h *userHandler) CheckEmailAvaibility(c *gin.Context) {
	//Ada input email dari user
	//Input email di-mapping ke struct input
	//Struct input di-passing ke service
	//Service akan memanggil reposiory - email apakah sudah ada atau belum
	//Repository - db

	// Berfungsi untuk menangkap input
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	//Melakukan pengecekan apakah ada error atau tidak
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)

	//Melakukan pengecekan apakah ada error
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Berfungsi untuk menampung hasil nilai dari pengecekan
	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	//Contoh tes
	var metaMessage string

	if IsEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// Fungsi untuk Upload avatar (gambar)
func (h *userHandler) UploadAvatar(c *gin.Context) {
	//Input dari user
	//simpan gambar di folder "images/"
	//di service kita panggil repo
	//JWT ( Sementara menggunakan hardcore, seakan akan user yg login ID = 1)
	//repo ambil data user yg ID = 1
	//repo update data user simpan lokasi file
}

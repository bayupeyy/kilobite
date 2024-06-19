package main

import (
	"fmt"
	"kilobite/auth"
	"kilobite/handler"
	"kilobite/helper"
	"kilobite/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/kilobite?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	fmt.Println(authService.GenerateToken(1001))

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	//Router
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	// Membuat fungsi untuk Middleware
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer tokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

//Ambil nilai header Authorization : Bearer token token
// Dari header Authorization, kita ambil nilai tokennya saja
// Kita validasi token
// Jika token valid, kita lanjutkan ke handler
// kita ambil user_id
// Ambil user dari db berdasarkan user_id lewat service
// Kita set context isinya user
//====================================================

//Berguna untuk mencari email by user
// userByEmail, err := userRepository.FindByEmail("bayuajike@gmail.com")
// if err != nil {
// 	fmt.Println(err.Error())
// }

// if userByEmail.ID == 0 {
// 	fmt.Println("User not found")
// } else {
// 	fmt.Println(userByEmail.Name)
// }

//Sementara di disable dulu karena perintah ini hanya untuk  uji coba
// userInput := user.RegisterUserInput{}
// userInput.Name = "John Doe"
// userInput.Email = "john@example.com"
// userInput.Occupation = "Programming"
// userInput.Password = "Karman"

//userService.RegisterUser(userInput)

// 	fmt.Println("Koneksi Berhasil")

// 	var users []user.User
// 	db.Find(&users)

// 	for _, users := range users {
// 		fmt.Println(users.Name)
// 		fmt.Println(users.Email)
// 		fmt.Println("===========")
// 	}

// router := gin.Default()
// router.GET("/handler", handler)
// router.Run()

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/kilobite?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }

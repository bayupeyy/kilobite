package main

import (
	"fmt"
	"kilobite/auth"
	"kilobite/campaign"
	"kilobite/handler"
	"kilobite/user"
	"log"

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
	campaignRepository := campaign.NewRepository(db)

	campaign, err := campaignRepository.FindByUserID(1)

	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println(len(campaign))
	for _, campaign := range campaign {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println("Jumlah Gambar")
			fmt.Println(len(campaign.CampaignImages))
			fmt.Println(campaign.CampaignImages[0].FileName)
		}

	}

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
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	//Ambil nilai header Authorization : Bearer token token
	// Dari header Authorization, kita ambil nilai tokennya saja
	// Kita validasi token
	// Jika token valid, kita lanjutkan ke handler
	// kita ambil user_id
	// !Ambil user dari db berdasarkan user_id lewat service
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
}

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

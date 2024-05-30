package main

import (
	"kilobite/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// 	dsn := "root:@tcp(127.0.0.1:3306)/kilobite?charset=utf8mb4&parseTime=True&loc=Local"
	// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}

	// 	fmt.Println("Koneksi Berhasil")

	// 	var users []user.User
	// 	db.Find(&users)

	// 	for _, users := range users {
	// 		fmt.Println(users.Name)
	// 		fmt.Println(users.Email)
	// 		fmt.Println("===========")
	// 	}

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/kilobite?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}

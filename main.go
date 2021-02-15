package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// setting database
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	// connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// when error
	if err != nil {
		log.Fatal(err.Error())
	}

	// print if connection success
	fmt.Println("Connection to database is good")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	// api versioning with group router
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

}
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yaswanthsaivendra/prod_mang/database"
	"github.com/yaswanthsaivendra/prod_mang/handlers"

	"github.com/yaswanthsaivendra/prod_mang/middleware"
	"github.com/yaswanthsaivendra/prod_mang/model"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()

}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Product{})
	database.Database.AutoMigrate(&model.Image{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", handlers.Register)
	publicRoutes.POST("/login", handlers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	router.MaxMultipartMemory = 8 << 20
	protectedRoutes.POST("/product", handlers.HandleProductUpload)
	protectedRoutes.GET("/product", handlers.GetAllProductsOfUser)

	router.Run(":8000")
	fmt.Println("server running on port 8000")
}

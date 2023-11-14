package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yaswanthsaivendra/prod_mang/database"
	"github.com/yaswanthsaivendra/prod_mang/handlers"

	"github.com/yaswanthsaivendra/prod_mang/middleware"
	"github.com/yaswanthsaivendra/prod_mang/model"
)

func main() {
	loadDB()
	serveApplication()

}

func Migrate(db *gorm.DB) {
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Product{})
	database.Database.AutoMigrate(&model.Image{})
}

func loadDB() {
	db := database.InitDB()
	Migrate(db)
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

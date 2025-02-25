package main

import (
	_ "e-commerce/cmd/docs"
	"e-commerce/internal/domain/services"
	"e-commerce/internal/infrastructure/config"
	"e-commerce/internal/infrastructure/http"
	"e-commerce/internal/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Commerce API
// @version 1.0
// @description This is a sample API for an e-commerce application
// @host localhost:8080
// @BasePath /
func main() {
	db := config.DbConnection()

	// Dependency Injection
	userRepo := persistence.NewUserRepositoryImp(db)
	service := services.NewUserService(userRepo)
	handler := http.NewUserHandler(service)

	httpServer := gin.Default()

	// Swagger route
	httpServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer.GET("/users", handler.GetAllUsers)
	httpServer.POST("/register", handler.RegisterUser)
	httpServer.PUT("/update", handler.Update)
	httpServer.GET("/get-by-id", handler.FindByID)
	httpServer.DELETE("/delete", handler.Delete)
	httpServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to E-Commerce API",
			"APIs":    "GET /users, POST /register, PUT /update, GET /get-by-id, DELETE /delete",
		})
	})

	httpServer.Run(": 8080")
}

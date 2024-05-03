package main

import (
	"gin/controllers"
	"gin/infra"

	"gin/repositories"
	"gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemContoroller(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	authRouter := r.Group("/auth")

	itemRouter.GET("/items", itemController.FindAll)
	itemRouter.GET("/items/:id", itemController.FindById)
	itemRouter.POST("/items", itemController.Create)
	itemRouter.PUT("/items/:id", itemController.Update)
	itemRouter.DELETE("/items/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

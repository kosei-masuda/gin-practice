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

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouter.GET("/items", itemController.FindAll)
	itemRouter.GET("/items/:id", itemController.FindById)
	itemRouter.POST("/items", itemController.Create)
	itemRouter.PUT("/items/:id", itemController.Update)
	itemRouter.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

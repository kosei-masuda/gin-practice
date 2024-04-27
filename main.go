package main

import (
	"gin/controllers"
	"gin/models"
	"gin/repositories"
	"gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{
			ID:          1,
			Name:        "商品1",
			Price:       1000,
			Description: "説明1",
			Soldout:     false,
		},
		{
			ID:          2,
			Name:        "商品2",
			Price:       1000,
			Description: "説明2",
			Soldout:     true,
		},
		{
			ID:          3,
			Name:        "商品3",
			Price:       1000,
			Description: "説明3",
			Soldout:     false,
		},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemContoroller(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}

package main

import (
	"gin/infra"
	"gin/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Faild to migrate database")
	}

}

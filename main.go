package main

import (
	"github.com/Coolenov/Fusion-api/api/controllers"
	"github.com/Coolenov/Fusion-library/gormDb"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	gormDb.DbConnect("root:firstpass@tcp(database:3306)/Fusion_db?utf8mb4&loc=Local")
}

func main() {
	r := gin.Default()

	r.GET("/contents/all/all", controllers.GetAllContents)
	r.GET("/tags/all/all", controllers.GetAllTags)
	r.GET("/sources", controllers.GetAllSources)

	r.POST("/t", controllers.GetContentsByTag)
	r.POST("/content/id", controllers.GetContentById)
	r.POST("/content/source", controllers.GetLastContentBySource)
	r.POST("/next", controllers.GetNextContent)
	r.POST("/previous", controllers.GetPreviousContent)

	err := r.Run(":10000")
	if err != nil {
		log.Println(err)
	}
}

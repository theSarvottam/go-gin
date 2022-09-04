package main

import (
	gin "github.com/gin-gonic/gin"
	api "github.com/theSarvottam/go-gin/api"
)

func main() {
	router := gin.Default()

	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.CreateAlbums)

	router.Run("localhost:8000")

}

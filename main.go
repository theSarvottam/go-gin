package main

import (
	gin "github.com/gin-gonic/gin"
	api "github.com/theSarvottam/go-gin/api"
)

func main() {
	router := gin.Default()

	router.GET("/album", api.GetAlbums)

	router.Run("localhost:8000")

}

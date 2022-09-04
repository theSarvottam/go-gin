package api

import (
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
	model "github.com/theSarvottam/go-gin/model"
	service "github.com/theSarvottam/go-gin/service"
)

func GetAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, service.GetAlbums())

}

func CreateAlbums(c *gin.Context) {
	var albumRequest model.CreateAlbumRequest

	if err := c.BindJSON(&albumRequest); err != nil {
		log.Fatal(err)
		return
	}

	albumResponse := service.CreateAlbum(albumRequest)

	c.IndentedJSON(http.StatusCreated, albumResponse)

}

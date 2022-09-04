package api

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
	service "github.com/theSarvottam/go-gin/service"
)

func GetAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, service.GetAlbums())

}

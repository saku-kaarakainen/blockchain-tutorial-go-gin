package main

import (
	"blockchain-tutorial-go-gin/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", endpoints.GetAlbums)

	router.Run("localhost:8080")
}

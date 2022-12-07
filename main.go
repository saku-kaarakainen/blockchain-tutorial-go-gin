package main

import (
	"blockchain-tutorial-go-gin/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	endpoints.NewTransaction(router)

	router.Run("localhost:8080")
}

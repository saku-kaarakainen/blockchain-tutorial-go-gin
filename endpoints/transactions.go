package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blockchain-tutorial-go-gin/blockchain"
)

// POST transactions/new
// func NewTransaction(context *gin.Context) {
func NewTransaction(router *gin.Engine) {
	// TODO: GET is only for testing. Make it POST
	router.GET("/transactions/new", func(context *gin.Context) {
		sender := context.Query("sender")
		recipient := context.Query("recipient")
		amount := context.Query("amount")

		index := blockchain.NewTransaction(sender, recipient, amount)
		context.String(http.StatusOK, "Transaction will be added to Block %d", index)
	})
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/receipts/process", processReceipt)

	router.Run(":8080")
}

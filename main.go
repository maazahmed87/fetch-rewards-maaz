package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)

	router.Run(":8080")
}

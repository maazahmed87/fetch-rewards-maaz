package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = make(map[string]Receipt)

func processReceipt(c *gin.Context) {
	var receipt Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getPoints(c *gin.Context) {
	id := c.Param("id")
	receipt, exists := receipts[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	points := calculatePoints(receipt)
	c.JSON(http.StatusOK, gin.H{"Points": points})
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Storing receipts in-memory
var receipts = make(map[string]Receipt)

// processReceipt handles the creation of a new receipt. It binds the incoming JSON data to the Receipt struct, validates it
// and then stores it in memory in the receipts in-memory map with a unique ID
func processReceipt(c *gin.Context) {
	var receipt Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// getPoints handles the retrieval of points for a given receipt by its ID
// It validates the ID then gets the receipt by the ID and calculates the points
func getPoints(c *gin.Context) {
	id := c.Param("id")

	if err := ValidateID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receipt, exists := receipts[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	points := calculatePoints(receipt)
	c.JSON(http.StatusOK, gin.H{"Points": points})
}

// getHealth provides a simple health check endpoint to see if the API is running
func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

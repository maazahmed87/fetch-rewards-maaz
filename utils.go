package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Regular expressions for fiels validation
var (
	idPattern               = regexp.MustCompile(`^\S+$`)        // Checks for no whitespace in ID
	retailerPattern         = regexp.MustCompile(`^[\w\s\-&]+$`) // Allows for alphanumericm, whitespace, hyphens and ampersand
	totalPattern            = regexp.MustCompile(`^\d+\.\d{2}$`) // Checks total for two decimal places
	shortDescriptionPattern = regexp.MustCompile(`^[\w\s\-]+$`)  // Allows alphanumeric, whitespace and hyphens
)

// ValidateReceipt checks each field of a receipt based on regex patterns and minimum item count
func ValidateReceipt(receipt Receipt) error {
	if !retailerPattern.MatchString(receipt.Retailer) {
		return errors.New("invalid format for retailer")
	}
	if !totalPattern.MatchString(receipt.Total) {
		return errors.New("invalid format for total")
	}
	if len(receipt.Items) == 0 {
		return errors.New("items array must contain at least one item")
	}
	for _, item := range receipt.Items {
		if !shortDescriptionPattern.MatchString(item.ShortDescription) {
			return errors.New("invalid format for item shortDescription")
		}
		if !totalPattern.MatchString(item.Price) {
			return errors.New("invalid format for item price")
		}
	}
	return nil
}

// ValidateID ensures that the receipt ID follows the specified pattern
func ValidateID(id string) error {
	if !idPattern.MatchString(id) {
		return errors.New("invalid format for ID")
	}
	return nil
}

// calculatePoints calculates total points for the a receipt by following the given rules
func calculatePoints(receipt Receipt) int {
	points := calculateRetailerPoints(receipt.Retailer)
	points += calculateTotalPoints(receipt.Total)
	points += calculateItemsPoints(receipt.Items)
	points += calculateDatePoints(receipt.PurchaseDate)
	points += calculateTimePoints(receipt.PurchaseTime)
	return points
}

// calculateRetailerPoints adds points based on alphanumeric characters in the retailer's name
func calculateRetailerPoints(retailer string) int {
	points := 0
	for _, char := range retailer {
		if isAlphanumeric(char) {
			points++
		}
	}
	return points
}

// calculateTotalPoints adds points based on total amount
func calculateTotalPoints(totalStr string) int {
	points := 0
	total, err := strconv.ParseFloat(totalStr, 64)
	if err != nil {
		log.Printf("Error parsing total: %v", err)
		return points
	}

	// check if total is negative
	if total < 0 {
		log.Printf("Negative total amount: %v", total)
		return points
	}

	// Check if the total is an integer value and not decimal
	if total == float64(int(total)) {
		points += 50
	}

	// Check if the total is divisible by 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	return points
}

// calculateItemsPoints adds points based on each item's short description length and price
func calculateItemsPoints(items []Item) int {
	points := (len(items) / 2) * 5
	for _, item := range items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			} else {
				log.Printf("Error parsing item price: %v", err)
			}
		}
	}
	return points
}

// calculateDatePoints adds points if the purchase date is an odd day
func calculateDatePoints(dateStr string) int {
	points := 0
	purchaseDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("Error parsing purchase date: %v", err)
		return points
	}
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}
	return points
}

// calculateTimePoints adds points if the purchase time is between 2:00 pm and 3:59 pm
func calculateTimePoints(timeStr string) int {
	points := 0
	purchaseTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		log.Printf("Error parsing purchase time: %v", err)
		return points
	}

	// Check if the purchase time is between 2:00 pm and 3:59 pm
	if purchaseTime.Hour() == 14 || purchaseTime.Hour() == 15 {
		points += 10
	}
	return points
}

// isAlphanumeric checks if a character is alphanumeric.
func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

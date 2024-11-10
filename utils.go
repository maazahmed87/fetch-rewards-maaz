package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// calculatePoints calculates total points for the given receipt by applying multiple rules.
func calculatePoints(receipt Receipt) int {
	points := calculateRetailerPoints(receipt.Retailer)
	points += calculateTotalPoints(receipt.Total)
	points += calculateItemsPoints(receipt.Items)
	points += calculateDatePoints(receipt.PurchaseDate)
	points += calculateTimePoints(receipt.PurchaseTime)
	return points
}

// calculateRetailerPoints adds points based on alphanumeric characters in the retailers name
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

	if total < 0 {
		log.Printf("Negative total amount: %v", total)
		return points
	}

	if err != nil {
		log.Printf("Error parsing total: %v", err)
		return points
	}

	if math.Mod(total, 1.0) == 0 {
		points += 50
	}
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

// calculateDatePoints adds points if purchase date is an odd day.
func calculateDatePoints(dateStr string) int {
	points := 0
	purchaseDate, err := time.Parse("2006-01-02", dateStr)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}
	return points
}

// calculateTimePoints adds points if purchase time is between 2:00pm and 3:59pm.
func calculateTimePoints(timeStr string) int {
	points := 0
	purchaseTime, err := time.Parse("15:04", timeStr)
	if err == nil && (purchaseTime.Hour() == 14 || purchaseTime.Hour() == 15) {
		points += 10
	}
	return points
}

// isAlphanumeric checks if a character is alphanumeric.
func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

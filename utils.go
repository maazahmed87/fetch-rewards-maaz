package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 10) == 0 {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

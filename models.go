package main

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"Price"`
}

type Receipt struct{
	Retailer string `json: "retailer"`
	PurchaseDate string `json:"purchaseDaete"`
	PurchaseTime string `json:"purchaseTime"`
	Items []Item `json:"items"`
	Total string `json:"total`
}
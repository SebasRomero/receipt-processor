package models

import "time"

type Receipt struct {
	Id           string
	Retailer     string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Total        float64
	Items        []Item
	Points       int
}

type SaveReceipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        string
	Items        []Item
}
type Item struct {
	ShortDescription string
	Price            string
}

type ProcessResponse struct {
	Id string
}

type PointsResponse struct {
	Points int
}

package models

import "time"

type Receipt struct {
	retailer     string
	purchaseDate time.Time `json:"purchase_date"`
	purchaseTime time.Time `json:"purchase_time"`
}

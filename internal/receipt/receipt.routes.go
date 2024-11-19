package receipt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	custom_errors "github.com/sebasromero/receipt-processor/internal/custom-errors"
	"github.com/sebasromero/receipt-processor/internal/db"
	"github.com/sebasromero/receipt-processor/internal/models"
)

func Process(w http.ResponseWriter, r *http.Request) {
	receipt := &models.SaveReceipt{}
	err := json.NewDecoder(r.Body).Decode(receipt)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(custom_errors.ErrorDecoding))
		return
	}

	newId, err := GenerateId()

	if err != nil {
		w.Write([]byte(custom_errors.ErrorGeneratingId))
		return
	}

	purchaseDate, err := ParseDate(receipt.PurchaseDate)

	if err != nil {
		w.Write([]byte(custom_errors.ErrorParsingDate))
		return
	}

	purchaseTime, err := ParseTime(receipt.PurchaseTime)

	if err != nil {
		w.Write([]byte(custom_errors.ErrorParsingTime))
		return
	}

	parsedPrice, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	newReceipt := &models.Receipt{
		Id:           newId,
		Retailer:     receipt.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Items:        receipt.Items,
		Total:        parsedPrice,
	}

	db.Receipts = append(db.Receipts, *newReceipt)

	fmt.Println(db.Receipts)

	json.NewEncoder(w).Encode(&models.ProcessResponse{
		Id: newId,
	})
}

func Points(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Points"))
}

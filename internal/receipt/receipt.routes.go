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
		w.WriteHeader(http.StatusBadGateway)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorGeneratingId,
		})
		return
	}

	purchaseDate, err := ParseDate(receipt.PurchaseDate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorParsingDate,
		})
		return
	}

	purchaseTime, err := ParseTime(receipt.PurchaseTime)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorParsingTime,
		})
		return
	}

	parsedPrice, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorParsingPrice,
		})
		return
	}

	if !ValidateAllItemsAreCorrect(receipt.Items) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorParsingPrice,
		})
		return
	}

	if !ValidatePriceArePositive(receipt.Items) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorNegativeNumber,
		})
		return
	}

	if ValidatePriceIsPositive(receipt.Total) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorNegativeNumber,
		})
		return
	}

	if ValidateYear(purchaseDate) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": custom_errors.ErrorBadYearInput,
		})
		return
	}
	newReceipt := &models.Receipt{
		Id:           newId,
		Retailer:     receipt.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Items:        receipt.Items,
		Total:        parsedPrice,
		Points:       0,
	}
	newReceipt.Points = calculatePoints(*newReceipt)

	db.Receipts = append(db.Receipts, *newReceipt)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&models.ProcessResponse{
		Id: newId,
	})

}

func Points(w http.ResponseWriter, r *http.Request) {
	for _, item := range db.Receipts {
		if r.PathValue("id") == item.Id {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&models.PointsResponse{
				Points: item.Points,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Receipt not found",
	})
}

func Receipts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&db.Receipts)
}

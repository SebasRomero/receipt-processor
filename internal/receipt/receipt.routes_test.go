package receipt

import (
	"strconv"
	"testing"

	custom_errors "github.com/sebasromero/receipt-processor/internal/custom-errors"
	"github.com/sebasromero/receipt-processor/internal/models"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		name     string
		receipt  models.SaveReceipt
		expected models.PointsResponse
	}{
		{
			name: "Target",
			receipt: models.SaveReceipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Total:        "35.35",
				Items: []models.Item{
					models.Item{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
					models.Item{
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					},
					models.Item{
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					},
					models.Item{
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					},
					models.Item{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
			},
			expected: models.PointsResponse{
				Points: 28,
			},
		},
		{
			name: "M&M Corner Market",
			receipt: models.SaveReceipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Total:        "9.00",
				Items: []models.Item{
					models.Item{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					models.Item{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					models.Item{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					models.Item{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					models.Item{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
				},
			},
			expected: models.PointsResponse{
				Points: 109,
			},
		},
	}
	for _, tc := range testCases {
		newId, err := GenerateId()

		if err != nil {
			t.Errorf(custom_errors.ErrorGeneratingId)
		}

		purchaseDate, err := ParseDate(tc.receipt.PurchaseDate)

		if err != nil {
			t.Errorf(custom_errors.ErrorParsingDate)
		}

		purchaseTime, err := ParseTime(tc.receipt.PurchaseTime)

		if err != nil {
			t.Errorf(custom_errors.ErrorParsingTime)
		}

		parsedPrice, err := strconv.ParseFloat(tc.receipt.Total, 64)

		if err != nil {
			t.Errorf(err.Error())
		}
		newReceipt := models.Receipt{
			Id:           newId,
			Retailer:     tc.receipt.Retailer,
			PurchaseDate: purchaseDate,
			PurchaseTime: purchaseTime,
			Total:        parsedPrice,
			Items:        tc.receipt.Items,
			Points:       0,
		}
		res := calculatePoints(newReceipt)
		if res != tc.expected.Points {
			t.Error("points are different than expected", tc.expected, res)
		}
	}
}

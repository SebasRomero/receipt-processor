package receipt

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sebasromero/receipt-processor/internal/models"
)

func GenerateId() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}

func ParseTime(unParsedTime string) (time.Time, error) {
	layout := "15:04"
	t, err := time.Parse(layout, unParsedTime)
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}

func calculatePoints(receipt models.Receipt) int {
	totalPoints := 0
	totalPoints += checkAlphaNumeric(receipt.Retailer)

	if checkTotalIsRound(receipt.Total) {
		totalPoints += 50
	}

	if checkTotalIsMultipleOfPoint25(receipt.Total) {
		totalPoints += 25
	}

	totalPoints += checkEveryTwoItems(receipt.Items)

	totalPoints += checkItemDescription(receipt.Items)

	if checkDayOdd(receipt.PurchaseDate) {
		totalPoints += 6
	}
	if checkTimeOfPurchase(receipt.PurchaseTime) {
		totalPoints += 10
	}

	return totalPoints
}

func checkAlphaNumeric(retailerName string) int {
	count := 0
	for _, v := range retailerName {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			count++
		}
	}

	return count
}

func checkTotalIsRound(total float64) bool {
	return total == float64(int64(total))
}

func checkTotalIsMultipleOfPoint25(total float64) bool {

	return math.Mod(total, 0.25) == 0
}

func checkEveryTwoItems(items []models.Item) int {
	return (len(items) / 2) * 5
}

func checkItemDescription(items []models.Item) int {
	points := 0
	for _, item := range items {
		itemDescriptionTrimmed := strings.Trim(item.ShortDescription, " ")
		if len(itemDescriptionTrimmed)%3 == 0 {
			res, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(res * 0.2))
		}
	}

	return points
}

func checkDayOdd(date time.Time) bool {
	return date.Day()%2 != 0
}

func checkTimeOfPurchase(date time.Time) bool {
	return date.Hour() >= 14 && date.Hour() < 16 && date.Minute() >= 1 && date.Minute() <= 59
}

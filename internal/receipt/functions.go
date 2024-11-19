package receipt

import (
	"time"

	"github.com/google/uuid"
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

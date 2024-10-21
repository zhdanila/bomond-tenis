package utils

import (
	"bomond-tenis/pkg/db/query"
	"fmt"
	"time"
)

func CheckTimeBookingBusy(existing, newBooking query.BookedCourt) (bool, error) {
	// Парсимо дату і час початку бронювання
	existingStartTime, err := time.Parse(time.RFC3339, existing.Date.String())
	if err != nil {
		return true, err
	}
	newStartTime, err := time.Parse(time.RFC3339, newBooking.Date.String())
	if err != nil {
		return true, err
	}

	existingEndTime := existingStartTime.Add(time.Duration(existing.Duration) * time.Minute)
	newEndTime := newStartTime.Add(time.Duration(newBooking.Duration) * time.Minute)

	fmt.Println(existingStartTime, existingEndTime)
	fmt.Println(newStartTime, newEndTime)
	fmt.Println()
	fmt.Println()

	if (newStartTime.Before(existingEndTime) && newEndTime.After(existingStartTime)) ||
		(existingStartTime.Before(newEndTime) && existingEndTime.After(newStartTime)) {

		return true, nil
	}

	return false, nil
}

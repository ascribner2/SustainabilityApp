package repos

import (
	"fmt"
	"time"
)

func CalcTimes(timespan string) (string, string) {
	now := time.Now()

	if timespan == "Annual" {
		year, _, _ := now.Date()

		return fmt.Sprintf("%d-01-01 00:00:00", year), now.Format("2006-01-02 15:04:05")
	}

	if timespan == "Month" {
		year, month, _ := now.Date()

		return fmt.Sprintf("%d-%02d-01 00:00:00", year, month), now.Format("2006-01-02 15:04:05")
	}

	if timespan == "Week" {
		weekday := now.Weekday()
		weekdayOffset := int(weekday) * -1
		weekStart := now.AddDate(0, 0, weekdayOffset).Format("2006-01-02")

		return fmt.Sprintf("%s 00:00:00", weekStart), now.Format("2006-01-02 15:04:05")
	}

	if timespan == "Day" {
		year, month, day := now.Date()

		return fmt.Sprintf("%d-%02d-%02d 00:00:00", year, month, day), now.Format("2006-01-02 15:04:05")
	}

	return "", ""
}

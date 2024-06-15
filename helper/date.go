package helper

import (
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02" // yyyy-mm-dd
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	// Set time to 23:59:59
	t = t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	return t, nil
}

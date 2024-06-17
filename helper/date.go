package helper

import (
	"strconv"
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

func ParseSeederTime(dateStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func Atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

package util

import (
	"os"
	"time"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func NowTime() string {
	return time.Now().Format("2006-01-02T15:04:05.999")
}

func ElapsedDay(startDate string, endDate string) int {
	st, _ := time.Parse("2006-01-02", startDate)
	et, _ := time.Parse("2006-01-02", endDate)
	return int(et.Sub(st).Hours() / 24)
}

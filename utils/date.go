package utils

import "time"

func GetCurrentDate() int64 {
	now := time.Now()
	return now.UnixMilli()
}

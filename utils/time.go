package utils

import (
	"fmt"
	"time"
)

func FormatDate(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func FormatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

package utils

import (
	"fmt"
	"time"
)

func FormatTime(dt time.Time) string {
	toFormat, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%v", dt))
	formatted := fmt.Sprintf("%v %v %v", toFormat.Day(), toFormat.Month(), toFormat.Year())

	return formatted
}

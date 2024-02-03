package helper

import (
	"fmt"
	"time"
)

func addOrdinal(n int) string {
	switch n {
	case 1, 21, 31:
		return fmt.Sprintf("%dst", n)
	case 2, 22:
		return fmt.Sprintf("%dnd", n)
	case 3, 23:
		return fmt.Sprintf("%drd", n)
	default:
		return fmt.Sprintf("%dth", n)
	}
}

func ToOrdinalDate(date time.Time) string {
	newOrdinalDate := fmt.Sprintf("%s %s, %s", date.Format("Jan"), addOrdinal(date.Day()), date.Format("15:04 AM"))
	return newOrdinalDate
}

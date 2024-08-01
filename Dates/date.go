package Dates

import (
	"time"
)

type AppDate struct {
	Date time.Time
}

func (date AppDate) GetDate() string {
	return date.Date.String()
}

func (date AppDate) FormatDate(Format string) string {
	return date.Date.Format(Format)
}

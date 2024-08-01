package Dates

import (
	"fmt"
	"time"
)

func GetDate(Format string) string {
	var currentDate string

	switch Format {
	case "mmddyyyy":
		currentDate = time.Now().Format("01022006")
	case "walletName":
		currentDate = time.Now().Format("01022006-1504")
	case "timestamp":
		currentDate = fmt.Sprintf("%d", time.Now().Unix())
	case "yyyy-mm-dd":
		currentDate = time.Now().Format("2006-01-02")
	case "dd-mm-yyyy":
		currentDate = time.Now().Format("02-01-2006")
	case "dd-mm-yyyy:HH:MM:SS":
		currentDate = time.Now().Format("02-01-2006:15:04:05")
	case "DT":
		currentDate = time.Now().Format("02-01-2006-15-04")
	case "yyyyMMdd":
		currentDate = time.Now().Format("20060102")
	case "logfile_yyyyMMdd":
		currentDate = time.Now().Format("logfile_20060102")
	default:
		currentDate = time.Now().Format("01022006:15:04:05")
	}

	return currentDate
}

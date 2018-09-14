package main

import (
	"strconv"
	"strings"
	"time"
)

func parseDateTime(year int, dateStr, timeStr string) time.Time {
	yearStr := strconv.Itoa(year)
	dateStr = strings.Replace(dateStr, "/", "-", 1)
	yearDateTimeStr := yearStr + "-" + dateStr + "T" + timeStr + ":00+08:00"

	t, err := time.Parse(time.RFC3339, yearDateTimeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func parseIPDateTime(year int, str string) (ip string, time time.Time) {
	strSlc := strings.Split(str, " ")
	if len(strSlc) == 2 { // IP address doesn't exist
		dateStr, timeStr := strSlc[0], strSlc[1]
		return "", parseDateTime(year, dateStr, timeStr)
	}
	// IP address exists
	ipStr, dateStr, timeStr := strSlc[0], strSlc[1], strSlc[2]
	return ipStr, parseDateTime(year, dateStr, timeStr)
}

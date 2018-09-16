package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

func sanitize(s, sep string) string {
	slc := strings.Split(s, sep)
	if len(slc) > 2 {
		return ""
	}
	var ans string
	for _, str := range slc {
		padded := padZero(str)
		if padded == "" {
			return ""
		}
		ans = ans + padded + sep
	}
	return strings.TrimRight(ans, sep)
}

func padZero(s string) string {
	if len(s) > 2 {
		return ""
	}
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Printf("error parse string [%s] to number\n", s)
		return ""
	}
	return fmt.Sprintf("%02d", u)
}

func parseDateTime(articleTime time.Time, dateStr, timeStr string) time.Time {
	if dateStr == "" || timeStr == "" {
		return articleTime
	}
	yearStr := strconv.Itoa(articleTime.Year())
	dateStr = strings.Replace(dateStr, "/", "-", 1)
	yearDateTimeStr := yearStr + "-" + dateStr + "T" + timeStr + ":00+08:00"

	t, err := time.Parse(time.RFC3339, yearDateTimeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func parseCommentIPDateTime(articleTime time.Time, str string) (ip string, time time.Time) {
	var ipStr, dateStr, timeStr string
	strSlc := strings.Split(str, " ")
	for _, str := range strSlc {
		if netIP := net.ParseIP(str); netIP != nil {
			ipStr = netIP.String()
		} else if strings.ContainsRune(str, '/') {
			dateStr = sanitize(str, "/")
		} else if strings.ContainsRune(str, ':') {
			timeStr = sanitize(str, ":")
		}
	}
	return ipStr, parseDateTime(articleTime, dateStr, timeStr)
}

func parseANSICTime(timeStr string) time.Time {
	t, err := time.Parse(time.ANSIC, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func loadModel(hit *elastic.SearchHit, v interface{}) {
	sourceBytes := []byte(*hit.Source)
	err := json.Unmarshal(sourceBytes, v)
	if err != nil {
		panic(err)
	}
}

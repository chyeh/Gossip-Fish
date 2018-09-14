package main

import (
	"testing"
	"time"
)

func TestParseIPDateTime(t *testing.T) {
	tests := []struct {
		input        string
		expectedIP   string
		expectedTime time.Time
	}{
		{"09/14 10:20", "", time.Date(2018, time.September, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"10/14 10:20", "", time.Date(2018, time.October, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"11/14 10:20", "", time.Date(2018, time.November, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"09/04 10:20", "", time.Date(2018, time.September, 4, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"10/04 10:20", "", time.Date(2018, time.October, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"11/04 10:20", "", time.Date(2018, time.November, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"10.20.30.40 09/14 10:20", "10.20.30.40", time.Date(2018, time.September, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"10.20.30.40 10/14 10:20", "10.20.30.40", time.Date(2018, time.October, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"127.0.0.1 11/14 10:20", "127.0.0.1", time.Date(2018, time.November, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"127.0.0.1 09/04 10:20", "127.0.0.1", time.Date(2018, time.September, 4, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"8.8.8.8 10/04 10:20", "8.8.8.8", time.Date(2018, time.October, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"8.8.8.8 11/04 10:20", "8.8.8.8", time.Date(2018, time.November, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
	}
	for i, v := range tests {
		actualIP, actualTime := parseIPDateTime(2018, v.input)
		if actualIP != v.expectedIP {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actualIP, v.expectedIP)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actualIP, v.expectedIP)

		if !actualTime.Equal(v.expectedTime) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actualTime, v.expectedTime)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actualTime, v.expectedTime)
	}
}

func TestParseDateTime(t *testing.T) {
	tests := []struct {
		inputDate string
		inputTime string
		expected  time.Time
	}{
		{"09/14", "10:20", time.Date(2018, time.September, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"10/14", "10:20", time.Date(2018, time.October, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"11/14", "10:20", time.Date(2018, time.November, 14, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"09/04", "10:20", time.Date(2018, time.September, 4, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"10/04", "10:20", time.Date(2018, time.October, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"11/04", "10:20", time.Date(2018, time.November, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},

		{"09/14", "00:20", time.Date(2018, time.September, 14, 0, 20, 0, 0, time.FixedZone("+0800", +8*60*60))},
		{"09/14", "10:00", time.Date(2018, time.September, 14, 10, 0, 0, 0, time.FixedZone("+0800", +8*60*60))},
	}
	for i, v := range tests {
		actual := parseDateTime(2018, v.inputDate, v.inputTime)
		if !actual.Equal(v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

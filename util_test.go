package main

import (
	"testing"
	"time"
)

func TestParseCommentIPDateTime(t *testing.T) {
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
		actualIP, actualTime := parseCommentIPDateTime(time.Date(2018, time.November, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60)), v.input)
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
	articleTime := time.Date(2018, time.November, 04, 10, 20, 0, 0, time.FixedZone("+0800", +8*60*60))
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
		{"09/14", "", articleTime},
		{"", "10:00", articleTime},
		{"", "", articleTime},
	}
	for i, v := range tests {
		actual := parseDateTime(articleTime, v.inputDate, v.inputTime)
		if !actual.Equal(v.expected) {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestPadZero(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-1", ""},
		{"-10", ""},
		{"a", ""},
		{"123", ""},
		{"1", "01"},
		{"9", "09"},
		{"10", "10"},
		{"11", "11"},
		{"19", "19"},
		{"20", "20"},
		{"23", "23"},
		{"24", "24"},
		{"24", "24"},
		{"59", "59"},
	}
	for i, v := range tests {
		actual := padZero(v.input)
		if actual != v.expected {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

func TestSanitize(t *testing.T) {
	tests := []struct {
		inputStr string
		inputSep string
		expected string
	}{
		{"9/1", "/", "09/01"},
		{"9/01", "/", "09/01"},
		{"09/1", "/", "09/01"},
		{"09/123", "/", ""},
		{"-9/1", "/", ""},
		{"a/1", "/", ""},

		{"9:1", ":", "09:01"},
		{"9:01", ":", "09:01"},
		{"09:1", ":", "09:01"},
		{"09:123", ":", ""},
		{"-9:1", ":", ""},
		{"a:1", ":", ""},

		{"22:11:34", ":", ""},
	}
	for i, v := range tests {
		actual := sanitize(v.inputStr, v.inputSep)
		if actual != v.expected {
			t.Errorf("Error on case %d: %v(actual) != %v(expected)", i, actual, v.expected)
		}
		t.Logf("Case %d: %v(actual) == %v(expected)", i, actual, v.expected)
	}
}

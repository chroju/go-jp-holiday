package holiday

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	cases := []struct {
		date   time.Time
		result bool
	}{
		{date: time.Date(2019, 4, 30, 0, 0, 0, 0, time.Local), result: true},
		{date: time.Date(2019, 7, 30, 0, 0, 0, 0, time.Local), result: false},
	}

	for _, c := range cases {
		if result, _ := IsHoliday(c.date); result != c.result {
			t.Error(c.result)
		}
	}
}

func TestGetHolidayName(t *testing.T) {
	cases := []struct {
		date   time.Time
		result string
	}{
		{date: time.Date(2019, 4, 30, 0, 0, 0, 0, time.Local), result: "休日"},
		{date: time.Date(2019, 7, 30, 0, 0, 0, 0, time.Local), result: ""},
	}

	for _, c := range cases {
		if result, _ := GetHolidayName(c.date); result != c.result {
			t.Error(c.result)
		}
	}
}

func TestGetHolidayNameToday(t *testing.T) {
	expected, err := GetHolidayName(time.Now())
	if err != nil {
		t.Error(err.Error())
	}

	if result, _ := GetHolidayNameToday(); result != expected {
		t.Error(expected)
	}
}

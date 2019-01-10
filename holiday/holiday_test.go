package holiday

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	cases := []struct {
		date   time.Time
		result string
	}{
		{date: time.Date(2019, 4, 30, 0, 0, 0, 0, time.Local), result: "休日"},
		{date: time.Date(2019, 7, 30, 0, 0, 0, 0, time.Local), result: ""},
	}

	for _, c := range cases {
		if IsHoliday(c.date) != c.result {
			t.Error(c.result)
		}
	}
}
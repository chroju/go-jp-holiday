package holiday

import (
	"encoding/csv"
	"io"
	"net/http"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	HolidayCsvURL = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
)

func IsHoliday(t time.Time) string {
	date := t.Format("2006/1/2")
	resp, err := http.Get(HolidayCsvURL)
	if err != nil {
		panic("get csv error")
	}
	defer resp.Body.Close()

	data := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))

	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if record[0] == date {
			return record[1]
		}
	}

	return ""
}

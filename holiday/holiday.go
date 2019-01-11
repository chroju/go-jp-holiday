package holiday

import (
	"encoding/csv"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	HolidayCsvURL = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
)

func IsHoliday(t time.Time) (bool, error) {
	result, err := checkHoliday(t.Format("2006/1/2"))
	if err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}
	return false, nil
}

func IsHolidayToday() (bool, error) {
	result, err := checkHoliday(time.Now().Format("2006/1/2"))
	if err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}
	return false, nil
}

func GetHolidayName(t time.Time) (string, error) {
	result, err := checkHoliday(t.Format("2006/1/2"))
	if err != nil {
		return "", err
	}

	return result, nil
}

func GetHolidayNameToday() (string, error) {
	result, err := checkHoliday(time.Now().Format("2006/1/2"))
	if err != nil {
		return "", err
	}

	return result, nil
}

func checkHoliday(date string) (string, error) {
	csvBody, err := getCsvBody()
	if err != nil {
		return "", err
	}

	for _, v := range csvBody {
		if v[0] == date {
			return v[1], nil
		}
	}

	return "", nil
}

func getCsvBody() ([][]string, error) {
	resp, err := http.Get(HolidayCsvURL)
	if err != nil {
		return nil, errors.Wrap(err, "Get CSV failed")
	}
	defer resp.Body.Close()

	data := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	csvData, err := data.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Read CSV failed")
	}

	return csvData, nil
}

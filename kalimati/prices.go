package kalimati

import (
	"encoding/json"
	"fmt"
)

const (
	API_URL = "https://kalimatimarket.gov.np/api/daily-prices"
)

type DailyPrice struct {
	Status int      `json:"status"`
	Date   string   `json:"date"`
	Prices []Prices `json:"prices"`
}

type Prices struct {
	Commodityname string `json:"commodityname"`
	Commodityunit string `json:"commodityunit"`
	Minprice      string `json:"minprice"`
	Maxprice      string `json:"maxprice"`
	Avgprice      string `json:"avgprice"`
}

func GetPrices(lang string) (DailyPrice, error) {
	var dailyPrice DailyPrice
	url := fmt.Sprintf("%s/%s", API_URL, lang)
	resp, err := NewHTTP("GET", url)
	if err != nil {
		return dailyPrice, err
	}

	if err = json.Unmarshal(resp, &dailyPrice); err != nil {
		return dailyPrice, err
	}

	return dailyPrice, nil
}

package helpers

import (
	"encoding/json"
	"io/ioutil"

	"kalimati/kalimati"
)

type JSONConverter struct{}

func (c *JSONConverter) Write(prices kalimati.DailyPrice) error {
	year, month, day := splitDate(prices.Date)
	file := fileName(year, month, day, "json")

	data, err := json.MarshalIndent(prices, "", " ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(file, data, 0o644); err != nil {
		return err
	}
	return nil
}

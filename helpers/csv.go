package helpers

import (
	"encoding/csv"
	"os"

	"kalimati/kalimati"
)

type CSVConverter struct{}

var headers = []string{"Date", "Product", "Unit", "Max Price", "Min Price", "Avg Price"}

func (c *CSVConverter) Write(data kalimati.DailyPrice) error {
	year, month, day := splitDate(data.Date)
	fileName := fileName(year, month, day, "csv")

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, item := range data.Prices {
		err = writer.Write([]string{data.Date, item.Commodityname, item.Commodityunit, item.Maxprice, item.Minprice, item.Avgprice})
		if err != nil {
			return err
		}
	}
	return nil
}

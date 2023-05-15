package helpers

import (
	"fmt"

	"kalimati/kalimati"

	"github.com/xuri/excelize/v2"
)

type XLSXConverter struct{}

func (c *XLSXConverter) Convert(data kalimati.DailyPrice) []map[string]string {
	var prices []map[string]string
	for index := range data.Prices {
		price := c.GetColumns(data.Prices[index], index, data.Date)
		prices = append(prices, price)
	}
	return prices
}

func (c *XLSXConverter) GetHeaders() map[string]string {
	return map[string]string{
		"A1": "Date", "B1": "Product", "C1": "Unit", "D1": "Max Price", "E1": "Min Price", "F1": "Avg Price",
	}
}

func (c *XLSXConverter) GetColumn(column string, num int) string {
	return fmt.Sprintf("%s%d", column, num+2)
}

func (c *XLSXConverter) GetColumns(price kalimati.Prices, num int, date string) map[string]string {
	return map[string]string{
		c.GetColumn("A", num): date,
		c.GetColumn("B", num): price.Commodityname,
		c.GetColumn("C", num): price.Commodityunit,
		c.GetColumn("D", num): price.Maxprice,
		c.GetColumn("E", num): price.Minprice,
		c.GetColumn("F", num): price.Avgprice,
	}
}

func (c *XLSXConverter) WriteFile(fileName string, headers map[string]string, data []map[string]string) {
	f := excelize.NewFile()
	for k, v := range headers {
		f.SetCellValue("Sheet1", k, v)
	}

	for _, vals := range data {
		for k, v := range vals {
			f.SetCellValue("Sheet1", k, v)
		}
	}

	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}

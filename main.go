package main

import (
	"kalimati/helpers"
	"kalimati/kalimati"
)

var langMap = map[string]string{
	"en": "en",
	"np": "np",
}

func main() {
	for lang := range langMap {
		prices, err := kalimati.GetPrices(lang)
		if err != nil {
			panic(err)
		}

		xlsx := helpers.XLSXConverter{}
		json := helpers.JSONConverter{}
		csv := helpers.CSVConverter{}

		if err := xlsx.Write(prices); err != nil {
			panic(err)
		}

		if err := json.Write(prices); err != nil {
			panic(err)
		}

		if err := csv.Write(prices); err != nil {
			panic(err)
		}
	}
}

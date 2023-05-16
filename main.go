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
	converters := []helpers.HelperInterface{
		&helpers.XLSXConverter{},
		&helpers.JSONConverter{},
		&helpers.CSVConverter{},
	}
	for lang := range langMap {
		prices, err := kalimati.GetPrices(lang)
		if err != nil {
			panic(err)
		}

		for _, converter := range converters {
			if err := converter.Write(prices); err != nil {
				panic(err)
			}
		}
	}
}

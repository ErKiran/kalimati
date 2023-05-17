package main

import (
	"github.com/ErKiran/kalimati/kalimati"

	"github.com/ErKiran/kalimati/helpers"
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

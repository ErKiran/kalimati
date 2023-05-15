package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

		date := prices.Date

		splittedDate := strings.Split(date, "-")
		year, month, day := splittedDate[0], splittedDate[1], splittedDate[2]

		file := fileName(year, month, day, "excel")

		xlsx.WriteFile(file, xlsx.GetHeaders(), xlsx.Convert(prices))

		// write data to json file

		data, _ := json.MarshalIndent(prices, "", " ")

		filej := fileName(year, month, day, "json")
		_ = ioutil.WriteFile(filej, data, 0o644)
	}
}

func pathChecker(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileName(year, month, day, output string) string {
	if output == "excel" {
		dirPath := fmt.Sprintf("data/excel/%s/%s", year, month)

		if err := pathChecker(dirPath); err != nil {
			panic(err)
		}

		fileName := fmt.Sprintf("%s/%s.xlsx", dirPath, day)
		return fileName
	}

	if output == "json" {
		dirPath := fmt.Sprintf("data/json/%s/%s", year, month)

		if err := pathChecker(dirPath); err != nil {
			panic(err)
		}

		fileName := fmt.Sprintf("%s/%s.json", dirPath, day)
		return fileName
	}
	return ""
}

package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/ErKiran/kalimati/kalimati"
)

type HelperInterface interface {
	Write(data kalimati.DailyPrice) error
}

var extMap = map[string]string{
	"excel": "xlsx",
	"json":  "json",
	"csv":   "csv",
}

func pathChecker(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func fileName(year, month, day, output string) string {
	dirPath := fmt.Sprintf("data/%s/%s/%s", output, year, month)

	if err := pathChecker(dirPath); err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s/%s.%s", dirPath, day, extMap[output])
	return fileName
}

func splitDate(date string) (string, string, string) {
	splittedDate := strings.Split(date, "-")
	return splittedDate[0], splittedDate[1], splittedDate[2]
}

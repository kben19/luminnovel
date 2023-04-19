package report

import (
	"github.com/tealeg/xlsx"
)

func readFile(path string) ([][][]string, error) {
	return xlsx.FileToSlice(path)
}

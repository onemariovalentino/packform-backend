package helper

import (
	"encoding/csv"
	"os"
)

func ReadCsvFile(csvFile string) (*csv.Reader, *os.File, error) {
	f, err := os.Open(csvFile)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}

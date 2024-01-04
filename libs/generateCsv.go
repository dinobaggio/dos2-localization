package libs

import (
	"encoding/csv"
	"log"
	"os"
)

func GenerateCSV(fileName string, data [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	log.Println("File CSV berhasil dibuat:", fileName)
	return nil
}

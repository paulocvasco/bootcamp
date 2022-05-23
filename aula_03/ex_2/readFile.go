package ex2

import (
	"bootcamp/aula_03/ex_2/model"
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile() error {

	// array to store results
	var data []model.CSVFormat

	file, err := os.Open("./saved_data.txt")
	if err != nil {
		return err
	}

	defer file.Close()

	csvData, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, csvLine := range csvData {
		dataLine := model.CSVFormat{
			ID:       csvLine[0],
			Price:    csvLine[1],
			Quantity: csvLine[2],
		}

		data = append(data, dataLine)
	}

	for i, d := range data {
		fmt.Printf("%d. ID: %s, Price: %s, Quantity: %s\n", i+1, d.ID, d.Price, d.Quantity)
	}

	return nil
}

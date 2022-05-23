package ex1

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func WriteData(id string, price string, quantity string) error {
	var file *os.File
	defer file.Close()

	// Validate input data
	if id == "" {
		return errors.New("empty ID")
	}

	if price == "" {
		return errors.New("invalid price")
	}

	if quantity == "" {
		return errors.New("invalid quantity")
	}

	// check if exists a file to save data
	_, err := os.Stat("./saved_data.txt")
	if err != nil {
		log.Print("File doesn't exist. Creating...")

		file, err = os.Create("./saved_data.txt")
		if err != nil {
			return err
		}

		log.Print("File saved_data.txt created on currently folder.")
	} else {
		file, err = os.Open("./saved_data.txt")
		if err != nil {
			return err
		}
	}

	// format data to csv format
	data := fmt.Sprintf("%s, %s, %s\n", id, price, quantity)

	// save data
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

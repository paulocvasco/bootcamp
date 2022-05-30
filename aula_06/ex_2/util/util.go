package util

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func GenerateID() string {
	id := uuid.New()
	return id.String()
}

func OpenFile() (*os.File, error) {
	var file *os.File
	// check if exists a file to save data
	_, err := os.Stat("./customers/customers.txt")
	if err != nil {
		_, err := os.Stat("./customers")
		if err != nil {
			log.Print("Folder doesn't exist. Creating...")

			err := os.MkdirAll("./customers", os.ModePerm)
			if err != nil {
				return nil, err
			}
			log.Print("Folder customers created on currently folder.")
		}
		file, err = os.Create("./customers/customers.txt")
		if err != nil {
			return nil, err
		}

		log.Print("File customers.txt created on ./customers folder.")
	} else {
		file, err = os.OpenFile("./customers/customers.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func SearchUser(file *os.File, id string) error {
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if scan.Text() == id {
			userFile := fmt.Sprintf("./customers/%s", id)
			file, err := os.OpenFile(userFile, os.O_RDONLY, os.ModeAppend)
			if err != nil {
				return err
			}
			defer file.Close()

			scanUser := bufio.NewScanner(file)
			for scanUser.Scan() {
				fmt.Println(scanUser.Text())
			}
			return nil
		}
	}
	return errors.New("user not found")
}

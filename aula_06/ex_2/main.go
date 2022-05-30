package main

import (
	"bootcamp/aula_06/ex_2/util"
	"log"
)

func main() {

	file, err := util.OpenFile()
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

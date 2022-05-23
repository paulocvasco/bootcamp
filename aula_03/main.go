package main

import (
	ex1 "bootcamp/aula_03/ex_1"
	ex2 "bootcamp/aula_03/ex_2"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	option := args[0]

	switch option {
	case "-w":
		err := ex1.WriteData(args[1], args[2], args[3])
		if err != nil {
			fmt.Println(err)
		}
	case "-r":
		err := ex2.ReadFile()
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Invalid option. Usage:")
		fmt.Println("Save data: aula_3 -w id price quantity")
		fmt.Println("Read data: aula_3 -r")
	}
}

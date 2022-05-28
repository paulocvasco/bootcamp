package main

import (
	"bootcamp/aula_04/ex_1/types"
	"fmt"
	"os"
)

func main() {
	var option string
	var args []string

	var users types.Users

	if len(os.Args) > 1 {
		option = os.Args[1]
		args = os.Args[2:]
	}

	switch option {
	case "newuser":
		if len(args) == 5 {
			users.AddUser(args[0], args[1], args[2], args[3], args[4])
		} else {
			fmt.Println("Missing argument. Usage:")
			fmt.Println("ex_1 newuser name lastname age email pass")
		}
	default:
		fmt.Println("Invalid option. Usage:")
		fmt.Println("Save data: aula_3 -w [id] [price] [quantity]")
		fmt.Println("Read data: aula_3 -r")
	}

	users.ShowAllUsers()
}

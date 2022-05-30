package main

import (
	"bootcamp/aula_06/ex_2/users"
	"bootcamp/aula_06/ex_2/util"
	"fmt"
	"log"
	"os"
)

func main() {

	var option string
	var args []string

	file, err := util.OpenFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if len(os.Args) > 1 {
		args = os.Args[1:]
		option = args[0]
	}

	switch option {
	case "-w":
		if len(args) == 5 {
			user := users.CreateUser(args[1], args[2], args[3], args[4])
			err = user.Save(file)
			if err != nil {
				log.Print(err)
			}
		} else {
			fmt.Println("Missing or invalid argument. Usage:")
			fmt.Println("Save user: ex_2 -w [name] [rg] [tel] [address]")
		}
	case "-s":
		if len(args) == 2 {
			err = util.SearchUser(file, args[1])
			if err != nil {
				log.Print(err)
			}
		} else {
			fmt.Println("Missing or invalid argument. Usage:")
			fmt.Println("Search user: ex_2 -s [id]")
		}
	default:
		fmt.Println("Invalid option. Usage:")
		fmt.Println("Save user: ex_2 -w [name] [rg] [tel] [address]")
		fmt.Println("Read data: ex_2 -s [id]")
	}

}

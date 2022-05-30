package main

import (
	"bootcamp/aula_04/ex_1/types"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var option string
	var args []string

	var users types.Users

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		args = strings.Split(input.Text(), " ")
		option = args[0]

		switch option {
		case "newuser":
			if len(args) == 6 {
				users.AddUser(args[1], args[2], args[3], args[4], args[5])
			} else {
				fmt.Println("Missing argument. Usage:")
				fmt.Println("newuser name lastname age email pass")
			}
		case "changename":
			if len(args) == 3 {
				users.ChangeName(args[1], args[2])
			} else {
				fmt.Println("Missing argument. Usage:")
				fmt.Println("changename username newname")
			}
		case "changeage":
			if len(args) == 3 {
				users.ChangeAge(args[1], args[2])
			} else {
				fmt.Println("Missing argument. Usage:")
				fmt.Println("changeage username newage")
			}
		case "changelastname":
			if len(args) == 3 {
				users.ChangeLastName(args[1], args[2])
			} else {
				fmt.Println("Missing argument. Usage:")
				fmt.Println("changelastname username newlastname")
			}
		case "changeemail":
			if len(args) == 3 {
				users.ChangeEmail(args[1], args[2])
			} else {
				fmt.Println("Missing argument. Usage:")
				fmt.Println("changeemail username newemail")
			}
		case "changepass":
			if len(args) == 3 {
				users.ChangePassword(args[1], args[2])
			} else {
				fmt.Println("Missing argument. Usage:")
				fmt.Println("changpass username newpass")
			}
		case "showusers":
			users.ShowAllUsers()
		default:
			fmt.Println("Invalid option. Usage:")
			fmt.Println("[newuser | changename | changeage | changelastname | changeemail | changepass] [params ...]")
		}
	}
}

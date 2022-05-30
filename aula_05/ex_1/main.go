package main

import (
	"bootcamp/aula_05/ex_1/errors"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var args []string
	var error errors.CustomError

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		args = strings.Split(input.Text(), " ")

		value, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		if value < 15000 {
			fmt.Println(error.Error())
		} else {
			fmt.Println("Pagar imposto")
		}
	}
}

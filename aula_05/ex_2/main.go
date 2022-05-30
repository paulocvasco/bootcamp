package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var args []string

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		args = strings.Split(input.Text(), " ")

		value, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		if value < 15000 {
			fmt.Println(errors.New("salario abaixo do minimo"))
		} else {
			fmt.Println("Pagar imposto")
		}
	}
}

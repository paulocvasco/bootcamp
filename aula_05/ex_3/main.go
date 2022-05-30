package main

import (
	"bufio"
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
			fmt.Println(fmt.Errorf("minimo tributavel: 15000, o salario informado eh: %d", value))
		} else {
			fmt.Println("Pagar imposto")
		}
	}
}

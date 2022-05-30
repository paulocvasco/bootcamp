package main

import (
	"bootcamp/aula_05/ex_4/util"
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
		if len(args) < 2 {
			fmt.Println(errors.New("missing argument. Usage: [worked hours] [salary per hour]"))
		} else {
			hours, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
			}
			value, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println(err)
			}
			salary, err := util.CalculateSalary(hours, value)
			if err != nil {
				fmt.Println(fmt.Errorf("failed to calculate salary: %s", err))
			} else {
				fmt.Printf("Salary %.2f\n", salary)
			}
		}
	}
}

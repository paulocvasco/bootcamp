package main

import "bootcamp/aula_02_tarde/ex_1/types"

func main() {

	students := []types.StudentData{
		{Name: "A", LastName: "B", RG: "111.111.111-1", StartDate: "01/01/2001"},
		{Name: "C", LastName: "D", RG: "222.222.222-2", StartDate: "02/02/2002"},
	}

	for _, student := range students {
		student.PrintInfo()
	}

}

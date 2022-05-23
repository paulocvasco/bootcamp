package types

import (
	"fmt"
)

type StudentData struct {
	Name      string
	LastName  string
	RG        string
	StartDate string
}

func (sd *StudentData) PrintInfo() {
	fmt.Println("============================")
	fmt.Printf("Student info:\n")
	fmt.Printf("Name: %s %s\n", sd.Name, sd.LastName)
	fmt.Printf("RG: %s\n", sd.RG)
	fmt.Printf("Start date: %s\n", sd.StartDate)
	fmt.Println("============================")
}

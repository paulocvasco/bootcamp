package types

import (
	"fmt"
	"time"
)

type StudentData struct {
	Name      string
	LastName  string
	RG        string
	StartDate time.Time
}

func (sd *StudentData) PrintInfo() {
	fmt.Printf("Student info:\n")
	fmt.Printf("Name: %s %s\n", sd.Name, sd.LastName)
	fmt.Printf("RG: %s\n", sd.RG)
	fmt.Printf("Start date: %s\n", sd.StartDate.Format("01/02/2006"))
}

package main

import "os"

func main() {
	file, err := os.OpenFile("customers.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic("file not found")
	}
	file.WriteString("foo")
}

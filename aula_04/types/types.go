package types

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/mail"
	"strconv"
)

type user struct {
	name     string
	lastName string
	age      int
	email    string
	password [32]byte
}

func (u *user) ChangeName(newName string) {
	u.name = newName
}

func (u *user) ChangeLastName(newLastName string) {
	u.lastName = newLastName
}

func (u *user) ChangeAge(newAge string) {
	value, err := strconv.Atoi(newAge)
	if err != nil {
		log.Printf("Invalid argument: %s\n", err)
		return
	}
	if value <= 0 {
		log.Print("Age must be positive\n")
		return
	}
	u.age = value
}

func (u *user) ChangeEmail(newEmail string) {
	// validate email
	_, err := mail.ParseAddress(newEmail)
	if err != nil {
		log.Printf("Invalid email: %s", err)
		return
	}
	u.email = newEmail
}

func (u *user) ChangePass(newPass string) {
	// save hash
	u.password = sha256.Sum256([]byte(newPass))
}

func (u *user) PrintContent() {
	fmt.Println("========================")
	fmt.Printf("Name: %s\n", u.name)
	fmt.Printf("Last name: %s\n", u.lastName)
	fmt.Printf("Age: %d\n", u.age)
	fmt.Printf("Email: %s\n", u.email)
	fmt.Printf("Hash password: %s\n", u.password)
	fmt.Printf("========================\n\n")
}

type Users []user

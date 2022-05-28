package types

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
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

func (u *user) changeName(newName string) {
	u.name = newName
}

func (u *user) changeLastName(newLastName string) {
	u.lastName = newLastName
}

func (u *user) changeAge(newAge string) {
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

func (u *user) changeEmail(newEmail string) {
	// validate email
	_, err := mail.ParseAddress(newEmail)
	if err != nil {
		log.Printf("Invalid email: %s", err)
		return
	}
	u.email = newEmail
}

func (u *user) changePass(newPass string) {
	// save hash
	u.password = sha256.Sum256([]byte(newPass))
}

func (u *user) printContent() {
	fmt.Println("========================")
	fmt.Printf("Name: %s\n", u.name)
	fmt.Printf("Last name: %s\n", u.lastName)
	fmt.Printf("Age: %d\n", u.age)
	fmt.Printf("Email: %s\n", u.email)
	fmt.Printf("Hash password: %s\n", base64.StdEncoding.EncodeToString(u.password[:]))
	fmt.Printf("========================\n\n")
}

type Users []user

func (u *Users) AddUser(name, lastname, age, email, pass string) {
	newUser := user{}

	newUser.changeName(name)
	newUser.changeLastName(lastname)
	newUser.changeAge(age)
	newUser.changeEmail(email)
	newUser.changePass(pass)

	*u = append(*u, newUser)
}

func (u Users) ShowUsers() {
	for _, user := range u {
		user.printContent()
	}
}

func (u *Users) ChangeName(username string, newname string) {
	user, err := u.searchUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	user.changeName(newname)
}

func (u *Users) ChangeLastName(username string, newLastName string) {
	user, err := u.searchUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	user.changeLastName(newLastName)
}

func (u *Users) ChangeEmail(username string, newEmail string) {
	user, err := u.searchUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	user.changeEmail(newEmail)
}

func (u *Users) ChangePassword(username string, newPass string) {
	user, err := u.searchUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	user.changePass(newPass)
}

func (u Users) ShowAllUsers() {
	for _, user := range u {
		user.printContent()
	}
}

func (u Users) searchUser(username string) (user, error) {
	for _, user := range u {
		if user.name == username {
			return user, nil
		}
	}
	return user{}, errors.New("user not found")
}

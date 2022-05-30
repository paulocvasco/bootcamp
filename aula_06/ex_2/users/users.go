package users

import (
	"bootcamp/aula_06/ex_2/util"
	"fmt"
	"os"
)

type User struct {
	Id      string
	Name    string
	Rg      string
	Tel     string
	Address string
}

func CreateUser(name, rg, tel, address string) User {
	newUser := User{
		Id:      util.GenerateID(),
		Name:    name,
		Rg:      rg,
		Tel:     tel,
		Address: address,
	}

	return newUser
}

func (u User) Save(allUsers *os.File) error {
	data := fmt.Sprintf("%s\n", u.Id)
	_, err := allUsers.WriteString(data)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("./customers/%s", u.Id)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	userData := fmt.Sprintf("ID: %s\nName: %s\nRG: %s\nTel: %s\nAddress: %s\n", u.Id, u.Name, u.Rg, u.Tel, u.Address)
	_, err = file.WriteString(userData)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

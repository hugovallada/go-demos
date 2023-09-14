package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type UserDomain struct {
	ID             string        `copier:"Id"`
	Name           string        `copier:"Name"`
	Age            int8          `copier:"Age"`
	LastName       string        `copier:"LastName"`
	Email          string        `copier:"Email"`
	DocumentNumber string        `copier:"DocumentNumber"`
	Address        AddressDomain `copier:"Address"`
}

type AddressDomain struct {
	Street     string `copier:"Street"`
	Number     int    `copier:"Number"`
	CEP        string `copier:"CEP"`
	State      string `copier:"State"`
	City       string `copier:"City"`
	Complement string `copier:"Complement"`
}

type UserRequest struct {
	ID             string         `copier:"Id"`
	Name           string         `copier:"Name"`
	Age            int8           `copier:"Age"`
	LastName       string         `copier:"LastName"`
	Email          string         `copier:"Email"`
	DocumentNumber string         `copier:"DocumentNumber"`
	Address        AddressRequest `copier:"Address"`
}

type AddressRequest struct {
	Street     string `copier:"Street"`
	Number     int    `copier:"Number"`
	CEP        string `copier:"CEP"`
	State      string `copier:"State"`
	City       string `copier:"City"`
	Complement string `copier:"Complement"`
}

func main() {
	userRequest := UserRequest{
		ID:             "123",
		Name:           "Code",
		Age:            27,
		LastName:       "Test",
		Email:          "test@gmail.com",
		DocumentNumber: "1234",
		Address: AddressRequest{
			Street:     "teste r",
			Number:     123,
			CEP:        "12344",
			State:      "SP",
			City:       "RP",
			Complement: "X",
		},
	}

	var userDomain UserDomain

	copier.Copy(&userDomain, &userRequest)

	fmt.Printf("Converção: %v\n", userDomain)
}

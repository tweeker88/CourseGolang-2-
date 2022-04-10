package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Age        int    `json:"age"`
	Social     Social `json:"social"`
}

type Social struct {
	Vk       string `json:"vk"`
	Facebook string `json:"facebook"`
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.FirstName)
	fmt.Printf("Type: %s\n", u.MiddleName)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social. VK: %s and FB: %s\n", u.Social.Vk, u.Social.Facebook)
}

func main() {
	handler, err := os.Open("users.json")
	defer handler.Close()

	if err != nil {
		log.Fatal("File is not open")
	}

	byteValue, err := ioutil.ReadAll(handler)

	if err != nil {
		log.Fatal("File is not readable")
	}

	var users Users
	json.Unmarshal(byteValue, &users)

	for _, user := range users.Users {
		PrintUser(&user)
	}
}

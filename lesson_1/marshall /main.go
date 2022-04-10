package main

import (
	"encoding/json"
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

func main() {
	users := Users{
		[]User{
			{"Maxim", "Teplov", "Olegovich", 25, Social{Vk: "vk.com/id123", Facebook: "facebook.com/id123213"}},
			{"Jenya", "Bandurina", "Nikolaevna", 22, Social{Vk: "vk.com/id123", Facebook: "facebook.com/id1766371"}},
		},
	}

	byteArray, err := json.MarshalIndent(users, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("users.json", byteArray, 0666)

	if err != nil {
		log.Fatal(err)
	}

}

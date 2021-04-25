package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// Id was of type string before

func LoadUser() map[string]User {
	u := make(map[string]User)
	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return u
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}
	return u
}

func StoreUser(u map[string]User) {
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(u)
	if err != nil {
		fmt.Println(err)
	}

}

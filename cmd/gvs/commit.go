package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//SSHkey defines ssh key and it's extra properties
type SSHkey struct {
	KeyName string `json:"keyname"`
	SSHkey  string `json:"ssh"`
	Public  bool   `json:"public"`
}

//Ref defines ssh key and it's extra properties
type Ref struct {
	ID string `json:"id"`
	Link  string `json:"link"`
	Other  bool   `json:"other"`
}

//User defined by name, email and set of ssh keys
type User struct {
	User    string `json:"user"`
	Email   string `json:"email"`
	SSHkeys []SSHkey
}

//Commit represents whole config file
type Commit struct {
	name string `json:"name"`
	description string `json:"description"`
	Date string `json:"date"`
	Author User
	Reference []Ref
}

//CreateCommitJSON dump values to config file
func CreateCommitJSON(conf GVSconfig) {
	b, err := json.Marshal(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

//CreateCommit dump values to config file
func CreateCommit() {
	// config := new(GVSconfig)
	// b, err := json.Marshal(config)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))
}

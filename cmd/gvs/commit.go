package gvs

import (
	"encoding/json"
	"fmt"
)

//Ref defines ssh key and it's extra properties
type Ref struct {
	ID    string `json:"id"`
	Link  string `json:"link"`
	Other bool   `json:"other"`
}

//Commit represents whole config file
type Commit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Author      User
	Reference   []Ref
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

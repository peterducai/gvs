package models

import (
	"encoding/json"
	"fmt"
)

//Change defines ssh key and it's extra properties
type Change struct {
	ID            string `json:"id"`
	FilePath      string `json:"filepath"`
	DiffPath      bool   `json:"diffpath"`
	PatchPath     bool   `json:"patchpath"`
	SelinuxChange string `json:"selinuxchange"` //unconfined_u:object_r:user_home_t:s0
}

//Commit represents whole config file
type Commit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Author      User
	Changes     []Change
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

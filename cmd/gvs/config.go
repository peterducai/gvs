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

//User defined by name, email and set of ssh keys
type User struct {
	User    string `json:"user"`
	Email   string `json:"email"`
	SSHkeys []SSHkey
}

//GVSconfig represents whole config file
type GVSconfig struct {
	Users []User
}

//ReadConfig reads gvs.config
func ReadConfig() {

}

//CreateConfig dump values to config file
func CreateConfig(conf GVSconfig) {
	b, err := json.Marshal(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

//CreateDefaultConfig dump values to config file
func CreateDefaultConfig() {
	config := new(GVSconfig)
	b, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
		return false
	}
	return !info.IsDir()
}

//ConfigHealthCheck checks for healthy config file
func ConfigHealthCheck() {
	fmt.Println("running CONFIG checks")
	fileExists("~/.gvsconfig.json")
}

//getFilePerm will get file permissions
func getFilePerm(fpath string) {

	fileInfo, err := os.Stat(fpath)

	// check if there is an error
	if err != nil {

		// check if error is file does not exist
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
		}

	}

	mode := fileInfo.Mode()

	fmt.Println(fileInfo, "mode is ", mode)
}

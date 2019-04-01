package main

import "os"

//Config
type SSHkey struct {
	SSHkey string `json:"ssh"`
}
type GVSconfig struct {
	User   string	`json:"user"`
	Email  string	`json:"email"`
	SSHkeys []SSHkey
}

//ReadConfig reads gvs.config
func ReadConfig(){

}

func CreateConfig(){

}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
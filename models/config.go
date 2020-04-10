package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//GVSconfig represents whole config file
type GVSconfig struct {
	Roles []Role
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
func CreateDefaultConfig(path string) {
	config := new(GVSconfig)
	b, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.Write(b)
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		//fmt.Println("File doesn't exist")
		return false
	}
	return !info.IsDir()
}

//ConfigHealthCheck checks for healthy config file
func ConfigHealthCheck() {
	usrdir, err := os.UserHomeDir()
	check(err)
	usrdir = usrdir + "/.gvsconfig.json"
	if !fileExists(usrdir) {
		CreateDefaultConfig(usrdir)
	}
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

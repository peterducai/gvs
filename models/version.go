package models

//Version of server
type Version struct {
	MAJOR    int    //version when you make incompatible API changes,
	MINOR    int    //version when you add functionality in a backwards compatible manner, and
	PATCH    int    //version when you make backwards compatible bug fixes.
	HASH     string //to verify integrity of software
	Startime string
}

var GvsVersion Version

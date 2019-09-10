package gvs

import (
	"encoding/json"
	"fmt"
)

type DocumentType int

const (
	GenericDoc DocumentType = iota
	Incident
	Change
	Problem
	Bug
	MeetingNote
	ProjectPlan
	Requirement
	HowTo
	Blog
)

func (dt DocumentType) String() string {
	return [...]string{"GenericDoc", "Incident", "Change", "Problem"}[dt]
}

//xSSHkey defines ssh key and it's extra properties
type Document struct {
	Id string `json:"id"`
	Path  string `json:"path"`
	Public  bool   `json:"public"`
}

//xCreateCommitJSON dump values to config file
func xCreateCommitJSON(conf GVSconfig) {
	b, err := json.Marshal(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

//xCreateCommit dump values to config file
func xCreateCommit() {
	// config := new(GVSconfig)
	// b, err := json.Marshal(config)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))
}

package gvs

import (
	"encoding/json"
	"fmt"
)

type Document int

const (
	GenericDoc Document = iota
	Tncident
	Change
	Problem
	Bug
	MeetingNote
	ProjectPlan
	Requirement
	HowTo
	Blog
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

//xSSHkey defines ssh key and it's extra properties
type xSSHkey struct {
	KeyName string `json:"keyname"`
	SSHkey  string `json:"ssh"`
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

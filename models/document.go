package models

import (
	"encoding/json"
	"fmt"
)

//DocumentType type of doc
type DocumentType int

//GenericDoc ia any doc  TODO: rewrite
const (
	GenericDoc DocumentType = iota
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

//Document defines ssh key and it's extra properties
type Document struct {
	ID           string       `json:"id"`
	Path         string       `json:"path"`
	Permissions  bool         `json:"permissions"`
	Encrypted    bool         `json:"encrypted"`
	DocumentType DocumentType `json:"documenttype"`
}

//CreateDocument dump values to config file
func CreateDocument(doc Document) {
	b, err := json.Marshal(doc)
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

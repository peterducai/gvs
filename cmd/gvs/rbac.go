package gvs

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

//Perm is permission for role/user
type Perm struct {
}

//Role to which user can be bind
type Role struct {
	Permissions []Perm
}

//GVSconfig represents whole config file
type GVSconfig struct {
	Roles []Role
}
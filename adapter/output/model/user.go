package model

type User struct {
	Email 	string
	Checked bool
}

type UserCode struct {
	Code  	string
	Email   string
}

var UserData = make(map[string]User)
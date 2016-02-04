package models

type Status struct {
	Status string `json:"status"`
}

type Json struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Status   Status `json:"status"`
}

type JsonSlice []Json

type User struct {
	Sys_Id    string `json:"sys_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Userslice []User

type Configuration struct {
	User     string
	Password string
	Instance string
}

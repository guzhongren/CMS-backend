package main

type Role struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	RoleId int    `json:"roleId"`
}

type Roleesponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

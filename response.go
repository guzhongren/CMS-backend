package main

type Response struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

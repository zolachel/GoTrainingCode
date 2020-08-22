package main

import (
	"encoding/json"
	"fmt"
)

//Todo ...
type Todo struct {
	ID     int    `json:id`
	Title  string `json:title`
	Status string `json:status`
}

func main() {
	data := []byte(`{
	"id": 1,
	"title": "pay credit card",
	"status": "active"
	}`)
	var t Todo
	err := json.Unmarshal(data, &t)
	fmt.Printf("% #v\n", t)
	fmt.Println(err)
}

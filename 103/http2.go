package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Todo ...
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		doPOST(w, req)
		return
	} else if req.Method == "GET" {
		doGET(w, req)
		return
	} else if req.Method == "PUT" {
		doPUT(w, req)
		return
	}
}
func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("starting...")

	log.Fatal(http.ListenAndServe(":1234", nil))

	fmt.Println("end")
}

func doPOST(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "error : %v", err)
		return
	}

	t := Todo{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Fprintf(w, "error: ", err)
	}

	todos = append(todos, t)
	fmt.Printf("%#v\n", todos)

	fmt.Printf("body : % #v\n", t)
	fmt.Fprintf(w, "hello %s created todos", "POST")
}

func doGET(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func doPUT(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "error : %v", err)
		return
	}

	t := Todo{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Fprintf(w, "error: ", err)
	}

	for i := 0; i < len(todos); i++ {
		fmt.Printf("i: %d, todos[i].ID: %v, t.ID: %v\n", i, todos[i].ID, t.ID)
		if todos[i].ID == t.ID {
			todos[i].ID = t.ID
			todos[i].Status = t.Status
		}
	}
}

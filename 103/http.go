package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Fprintf(w, "Hello: %s", req.Method)
	}

	resp := []byte(`{"name": "beer"}`)
	w.Header().Add("content-type", "application/json")
	w.Write(resp)
}
func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("starting...")
	//err := http.ListenAndServe(":1234567890", nil)
	//log.Fatal(err)
	log.Fatal(http.ListenAndServe(":1234", nil))

	fmt.Println("end")
}

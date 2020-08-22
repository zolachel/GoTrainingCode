package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./error.go")

	if err != nil {
		log.Fatal(err)

		return errors.Errorf("can't openfile: %s", err.Error())
	}

	f.Close()

	fmt.Println("Finished")
}

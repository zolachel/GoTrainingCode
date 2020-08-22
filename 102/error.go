package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
	Err      error
	Code     int
	Severity int
}

func (buss *BusinessError) Error() string {
	return fmt.Printf("err=%s, code=%d, serverity=%d", b.Err.Error(), b.Code, b.Severity)
}

func PrintErr(err error) {
	fmt.Println(err)
}

func main() {
	err := errors.New("my error msg")
	PrintErr(err)
	berr := &BusinessError{Err: err, Code: 555, serverity: 1}
	PrintErr(berr)

	if err != nil {

	}
}

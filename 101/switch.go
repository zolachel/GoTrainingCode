package main

import (
	"fmt"
	"runtime"
)

func main() {

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("It's linux")
		fallthrough
	case "darwin":
		fmt.Println("It's darwin")
	default:
		fmt.Printf("It's %s\n", os)
	}
}

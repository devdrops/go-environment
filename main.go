package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO VERSION IS", runtime.Version())
	fmt.Println("GO ROOT IS", runtime.GOROOT())

	info := `Instructions:
	1. Run docker-compose up
	2. Edit any Go file and see the changes being executed`

	fmt.Println(info)
}

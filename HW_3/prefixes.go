package main

import (
	"fmt"
	"os"
)

func main() {
	inputString := os.Args[1]
	for i := 0; i < len(inputString)+1; i++ {
		fmt.Println(inputString[:i])
	}
}

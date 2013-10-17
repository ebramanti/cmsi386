//	Problem #4
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename, error := os.Open(os.Args[1])
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(filename)
	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && line[0] != '#' {
			count += 1
		}
	}
	fmt.Println(count)
}

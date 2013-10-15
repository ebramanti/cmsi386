package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	filename, error := os.Open(os.Args[1]), ...
	scanner := bufio.NewScanner(filename)
	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if  line != "" && line[0] != '#' {
			count += 1
		}
	}
}
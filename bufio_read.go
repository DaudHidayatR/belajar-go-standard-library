package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := strings.NewReader("hello world")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
	}
}

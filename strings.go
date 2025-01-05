package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.Trim("...test...", "."))
	fmt.Println(strings.Replace("test", "e", "E", -1))

}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("e[a-z]+")
	fmt.Println(regex.MatchString("example"))
	fmt.Println(regex.MatchString("is"))
	fmt.Println(regex.MatchString("good"))
	fmt.Println(regex.MatchString("Example is good"))

}

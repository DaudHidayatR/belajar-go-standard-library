package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error parsing bool: ", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("error parsing int: ", err.Error())
	} else {
		fmt.Println(resultInt)
	}

}

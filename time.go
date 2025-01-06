package main

import (
	"fmt"
	"time"
)

func main() {
	new := time.Now()
	fmt.Println(new)

	utc := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := time.Now()
	valuetime, err := time.Parse(formatter, value.Format(formatter))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valuetime)
	}
}

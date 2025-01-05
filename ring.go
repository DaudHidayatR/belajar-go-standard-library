package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(3)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		fmt.Println(data.Value)
		data = data.Next()

	}
	//
	//data.Value = "daud"
	//data = data.Next()
	//
	//data.Value = "hidayat"
	//data = data.Next()
	//
	//data.Value = "ramadhan"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}

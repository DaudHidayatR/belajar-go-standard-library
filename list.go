package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(1)
	data.PushBack("daud")
	data.PushBack("hidayat")
	data.PushBack("ramadhan")

	head := data.Front()
	fmt.Println(head.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("time.go"))
	fmt.Println(path.Base("time.go"))
	fmt.Println(path.Ext("time.go"))
	fmt.Println(path.Join("a", "b", "c"))

}

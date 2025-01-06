package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	fmt.Println(filepath.Dir("time.go"))
	fmt.Println(filepath.Base("time.go"))
	fmt.Println(filepath.Ext("time.go"))
	fmt.Println(filepath.Join("a", "b", "c"))

}

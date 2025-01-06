package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"John", "Doe"}
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(name, "daud"))
	fmt.Println(slices.Contains(name, "Doe"))
	fmt.Println(slices.Contains(values, 3))

}

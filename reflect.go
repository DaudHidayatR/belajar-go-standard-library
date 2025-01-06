package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `require:"true" max:"10"`
}

type Person struct {
	Name string `require:"true" max:"10"`
	Age  int    `require:"true"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type:", valueType)
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Println(" Field:", field.Name)
		fmt.Println("  Type:", field.Type)
		fmt.Println("  Tag:", field.Tag)
		fmt.Println("  Require:", field.Tag.Get("require"))
		fmt.Println("  Max:", field.Tag.Get("max"))
	}
}

func Isvalid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("require") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if !result {
				return false
			}
		}
	}
	return result
}

func main() {
	sample := Sample{Name: "daud"}
	person := Person{Name: "hidayat", Age: 22}

	readField(sample)
	readField(person)
	fmt.Println(Isvalid(person))
}

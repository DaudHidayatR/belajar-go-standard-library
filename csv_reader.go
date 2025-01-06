package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	// Code for CSV Reader
	csvString := "daud,hidayat,ramadhan\n" +
		"22,20,21\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(record)
	}
}

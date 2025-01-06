package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "daud Hidayat Ramadhan"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded: ", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Decoded: ", string(decoded))
	}

}

package main

import (
	"bufio"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(message)
	return nil
}
func readFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil

}
func addFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(message)
	return nil
}

func main() {
	//createNewFile("test.txt", "hello world")
	//message, _ := readFile("test.txt")
	//fmt.Println(message)
	addFile("test.txt", "\nhello world")
}

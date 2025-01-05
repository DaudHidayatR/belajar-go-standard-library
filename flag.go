package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "a string")
	port := flag.Int("port", 80, "an int")
	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
}

package main

import (
	"fmt"
	"log"
)

func getPlainMessage() {
	var plainMessage string
	fmt.Println("Please input your message:")
	count, err := fmt.Scanln(&plainMessage)
	//plainMessage, err := fmt.Scanln(&plainMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(plainMessage)
}

func main() {
	getPlainMessage()
}

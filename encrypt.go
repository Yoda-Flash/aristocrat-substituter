package main

import (
	"fmt"
)

func getPlainMessage() {
	var plainMessage string
	fmt.Println("Please input your message:")
	fmt.Scanln(&plainMessage)
	//plainMessage, err := fmt.Scanln(&plainMessage)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(plainMessage)
}

func main() {
	getPlainMessage()
}

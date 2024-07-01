package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getPlainMessage(scanner *bufio.Scanner) string {
	var plainMessage string
	fmt.Println("Please input your message:")
	scanner.Scan()
	err := scanner.Err()
	plainMessage = scanner.Text()
	//count, err := fmt.Scanln(&plainMessage)
	////plainMessage, err := fmt.Scanln(&plainMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(plainMessage)
	//fmt.Println(plainMessage)
	//fmt.Println(count)
	//for i := 0; i < count; i++ {
	//	fmt.Println(plainMessage)
	//}
	return plainMessage
}

func getKey(scanner *bufio.Scanner) string {
	fmt.Println("Please input your key, without any spaces in between:")
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	var key string
	key = scanner.Text()
	fmt.Println(key)
	return key
}

func getCodedMessage(plainMessage string, key string) string {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	plainMessage := getPlainMessage(scanner)
	key := getKey(scanner)
	codedMessage := getCodedMessage(plainMessage, key)
	fmt.Println(codedMessage)
}

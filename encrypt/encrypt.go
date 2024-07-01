package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getPlainMessage(scanner *bufio.Scanner) string {
	var plainMessage string
	fmt.Println("Please input the message you wish to encrypt:")
	scanner.Scan()
	err := scanner.Err()
	plainMessage = scanner.Text()
	plainMessage = strings.TrimSpace(plainMessage)
	plainMessage = strings.ToLower(plainMessage)
	//count, err := fmt.Scanln(&plainMessage)
	////plainMessage, err := fmt.Scanln(&plainMessage)
	if err != nil {
		log.Fatal(err)
	}
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
	key = strings.TrimSpace(key)
	key = strings.ToLower(key)
	key += " "
	return key
}

func returnCodedMessage(plainMessage string, key string) string {
	var codedMessage string
	alphabet := "abcdefghijklmnopqrstuvwxyz "
	for i := 0; i < len(plainMessage); i++ {
		for j := 0; j < 27; j++ {
			if plainMessage[i] == alphabet[j] {
				codedMessage += string(key[j])
			}
		}
	}
	return codedMessage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	plainMessage := getPlainMessage(scanner)
	key := getKey(scanner)
	codedMessage := returnCodedMessage(plainMessage, key)
	fmt.Println(codedMessage)
}

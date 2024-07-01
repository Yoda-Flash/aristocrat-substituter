package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCodedMessage(scanner *bufio.Scanner) string {
	var codedMessage string
	fmt.Println("Please input the message you wish to decrypt:")
	scanner.Scan()
	err := scanner.Err()
	codedMessage = scanner.Text()
	codedMessage = strings.TrimSpace(codedMessage)
	codedMessage = strings.ToLower(codedMessage)
	if err != nil {
		log.Fatal(err)
	}
	return codedMessage
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

func returnPlainMessage(codedMessage string, key string) string {
	var plainMessage string
	alphabet := "abcdefghijklmnopqrstuvwxyz "
	for i := 0; i < len(codedMessage); i++ {
		for j := 0; j < 27; j++ {
			if codedMessage[i] == key[j] {
				plainMessage += string(alphabet[j])
			}
		}
	}
	return plainMessage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	codedMessage := getCodedMessage(scanner)
	key := getKey(scanner)
	plainMessage := returnPlainMessage(codedMessage, key)
	fmt.Println(plainMessage)
}

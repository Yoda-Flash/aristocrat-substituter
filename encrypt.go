package aristocrat_substituter

import (
	"fmt"
)

func getPlainMessage() {
	var plainMessage string
	fmt.Println("Please input your message:")
	fmt.Scanln(plainMessage)
	fmt.Println(plainMessage)
}

func main() {
	getPlainMessage()
}

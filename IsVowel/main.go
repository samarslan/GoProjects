package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter chracter:")
	chracter, _, _ := reader.ReadRune()

	if chracter == 'a' || chracter == 'e' || chracter == 'i' || chracter == 'o' || chracter == 'u' || chracter == 'A' || chracter == 'E' || chracter == 'I' || chracter == 'O' || chracter == 'U' {
		fmt.Println("You entered a vowel.\n")
	} else {
		fmt.Println("You entered a consonant.\n")
	}
}

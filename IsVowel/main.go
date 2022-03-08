package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter character:")
	character, _, _ := reader.ReadRune()

	if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' || character == 'A' || character == 'E' || character == 'I' || character == 'O' || character == 'U' {
		fmt.Println("You entered a vowel.\n")
	} else {
		fmt.Println("You entered a consonant.\n")
	}
}

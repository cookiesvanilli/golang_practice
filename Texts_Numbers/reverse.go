package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the text to be reversed: ")
	a := bufio.NewReader(os.Stdin)
	input, err := a.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	text := []rune(input)
	num := len(text)
	for i := 0; i < num/2; i++ {
		text[i], text[num-i-2] = text[num-i-2], text[i]
	}
	fmt.Println(string(text))
}
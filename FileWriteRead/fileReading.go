package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fmt.Println("--- Scanner ---")
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	file.Seek(0,0)//указатель сместили на начало

	fmt.Println("--- Reader---")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
	}
}

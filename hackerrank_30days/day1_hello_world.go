package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a line and press enter")

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Hello, World")
	fmt.Print(line)
}

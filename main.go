package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	tokens, err := tokenize(f)
	if err != nil {
		fmt.Println("tokenization error:", err)
		os.Exit(1)
	}

	err = parseObject(tokens)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("valid JSON")
}

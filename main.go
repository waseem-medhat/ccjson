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

	err = parse(tokenize(f))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("valid JSON")
}

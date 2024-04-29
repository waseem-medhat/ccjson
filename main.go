package main

import (
	"bufio"
	"fmt"
	"os"
)

type Token int

const (
	OpenBrace    Token = iota // Opening curly brace `{` token
	CloseBrace                // Closing curly brace `}` token
	OpenBracket               // Opening square bracket `[` token
	CloseBracket              // Closing square bracket `]` token
	True                      // Boolean `true` token
	False                     // Boolean `false` token
	Null                      // `null` token
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

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

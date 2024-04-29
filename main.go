package main

import (
	"bufio"
	"fmt"
	"os"
)

type TokenType int

const (
	BeginObject    TokenType = iota // Opening curly brace `{` token
	EndObject                       // Closing curly brace `}` token
	BeginArray                      // Opening square bracket `[` token
	EndArray                        // Closing square bracket `]` token
	NameSeparator                   // Colon `:` token
	ValueSeparator                  // Comma `,` token
	True                            // Boolean `true` token
	False                           // Boolean `false` token
	Null                            // `null` token
	String
	Number
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

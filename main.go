package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

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
	defer f.Close()

	err = isValid(tokenize(f))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
}

func tokenize(f *os.File) []Token {
	TokenMap := map[string]TokenType{
		"{":     BeginObject,
		"}":     EndObject,
		"[":     BeginArray,
		"]":     EndArray,
		":":     NameSeparator,
		",":     ValueSeparator,
		"true":  True,
		"false": False,
		"null":  Null,
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	tokens := []Token{}
	for s.Scan() {
		c := s.Text()
		if tokenType, ok := TokenMap[s.Text()]; ok {
			tokens = append(tokens, Token{Type: tokenType, Value: c})
		}
	}

	return tokens
}

func isValid(tokens []Token) error {
	nBeginObject := 0
	nEndObject := 0

	for _, t := range tokens {
		switch t.Type {
		case BeginObject:
			nBeginObject++
		case EndObject:
			if nBeginObject == 0 {
				return errors.New("unexpected closing brace")
			}
			nEndObject++
		}
	}

	return nil
}

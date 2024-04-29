package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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

	err = parse(tokenize(f))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("valid JSON")
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
	inString := false
	strBuilder := strings.Builder{}
	for s.Scan() {
		c := s.Text()
		if tokenType, ok := TokenMap[s.Text()]; ok {
			tokens = append(tokens, Token{Type: tokenType, Value: c})
			continue
		}

		if c == "\"" && inString {
			inString = false
			tokens = append(tokens, Token{Type: String, Value: strBuilder.String()})
			continue
		}

		if c == "\"" && !inString {
			inString = true
			strBuilder.Reset()
			continue
		}

		if inString {
			strBuilder.Write([]byte(c))
		}
	}

	return tokens
}

func parse(tokens []Token) error {
	if len(tokens) == 0 || tokens[0].Type != BeginObject {
		return errors.New("not an object")
	}

	nBeginObject := 1
	nEndObject := 0
	for i := 1; i < len(tokens); i++ {
		t := tokens[i]
		tPrev := tokens[i-1]
		fmt.Println(t) // TODO: delete

		if (tPrev.Type == BeginObject || tPrev.Type == ValueSeparator) && t.Type != String {
			return errors.New("non-string key")
		}

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

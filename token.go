package main

import (
	"bufio"
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

package main

import (
	"bufio"
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

func tokenize(f *os.File) ([]Token, error) {
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
	literalBuilder := strings.Builder{}
	for s.Scan() {
		c := s.Text()
		if tokenType, ok := TokenMap[c]; ok {
			tokens = append(tokens, Token{Type: tokenType, Value: c})
			continue
		}

		if c == "\"" && inString {
			inString = false
			tokens = append(tokens, Token{Type: String, Value: strBuilder.String()})
			strBuilder.Reset()
			continue
		}

		if c == "\"" && !inString {
			inString = true
			continue
		}

		if inString {
			strBuilder.Write([]byte(c))
			continue
		}

		if _, ok := TokenMap[c]; ok || strings.TrimSpace(c) == "" {
			newLiteral := literalBuilder.String()
			if newLiteral == "" {
				continue
			}

			if tokenType, ok := TokenMap[newLiteral]; ok {
				tokens = append(tokens, Token{Type: tokenType, Value: newLiteral})
				literalBuilder.Reset()
				continue
			}

			return tokens, fmt.Errorf("unexpected literal %v", newLiteral)
		}

		literalBuilder.Write([]byte(c))
	}

	for _, t := range tokens {
		fmt.Println(t.Value)
	}

	return tokens, nil
}

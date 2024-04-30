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
	tokenMap := validTokens()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	tokens := []Token{}
	inString := false
	strBuilder := strings.Builder{}
	literalBuilder := strings.Builder{}
	for s.Scan() {
		c := s.Text()
		if tokenType, ok := tokenMap[c]; ok {
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

		if _, ok := tokenMap[c]; ok || strings.TrimSpace(c) == "" {
			newLiteral := literalBuilder.String()
			if newLiteral == "" {
				continue
			}

			if tokenType, ok := tokenMap[newLiteral]; ok {
				tokens = append(tokens, Token{Type: tokenType, Value: newLiteral})
				literalBuilder.Reset()
				continue
			}

			return tokens, fmt.Errorf("unexpected literal %v", newLiteral)
		}

		literalBuilder.Write([]byte(c))
	}

	return tokens, nil
}

func validTokens() map[string]TokenType {
	return map[string]TokenType{
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
}

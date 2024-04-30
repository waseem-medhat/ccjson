package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		if tokenType, ok := tokenMap[c]; ok || strings.TrimSpace(c) == "" {
			literal := strings.TrimSpace(literalBuilder.String())
			newToken, err := tokenizeLiteral(literal)
			if err != nil && err.Error() == "invalid literal" {
				return tokens, fmt.Errorf("tokenization error: %v %v", err, literal)
			}

			if err == nil {
				tokens = append(tokens, newToken)
				literalBuilder.Reset()
			}

			if ok {
				tokens = append(tokens, Token{Type: tokenType, Value: c})
			}
			continue
		}

		literalBuilder.Write([]byte(c))
	}

	return tokens, nil
}

func tokenizeLiteral(literal string) (Token, error) {
	if literal == "" {
		return Token{}, fmt.Errorf("empty input")
	}

	tokenMap := validTokens()
	if tokenType, ok := tokenMap[literal]; ok {
		return Token{Type: tokenType, Value: literal}, nil
	}

	if _, err := strconv.ParseFloat(literal, 64); err == nil {
		return Token{Type: Number, Value: literal}, nil
	}

	return Token{}, fmt.Errorf("invalid literal")
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

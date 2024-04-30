package main

import (
	"errors"
	"fmt"
)

func parse(tokens []Token) error {
	if len(tokens) == 0 || tokens[0].Type != BeginObject {
		return errors.New("not an object")
	}

	nBeginObject := 1
	nEndObject := 0
	for i := 1; i < len(tokens); i++ {
		t := tokens[i]
		tPrev := tokens[i-1]
		// fmt.Println(t) // TODO: delete

		if tPrev.Type == ValueSeparator && t.Type != String {
			return fmt.Errorf("unexpected token: %v", t.Value)
		}

		if tPrev.Type == BeginObject && !(t.Type == String || t.Type == EndObject) {
			return fmt.Errorf("unexpected token: %v", t.Value)
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

package main

import (
	"os"
	"testing"
)

type subtest struct {
	fileName    string
	validTokens bool
	validJSON   bool
}

func run(subtests []subtest, t *testing.T) {
	for _, st := range subtests {
		f, err := os.Open(st.fileName)
		if err != nil {
			t.Fatalf("couldn't open file %v", st.fileName)
		}
		defer f.Close()

		tokens, err := tokenize(f)
		if err != nil && st.validTokens {
			t.Fatalf("%v with valid tokens got error: %v", st.fileName, err)
		}

		if err == nil && !st.validTokens {
			t.Fatalf("%v with invalid tokens passed tokenization", st.fileName)
		}

		// files that failed tokenization won't be parsed
		if err != nil {
			return
		}

		err = parseObject(tokens)
		if err != nil && st.validJSON {
			t.Fatalf("valid %v got error: %v", st.fileName, err)
		}

		if err == nil && !st.validJSON {
			t.Fatalf("invalid %v passed", st.fileName)
		}
	}
}

func TestStep1(t *testing.T) {
	subtests := []subtest{
		{fileName: "tests/step1/valid.json", validTokens: true, validJSON: true},
		{fileName: "tests/step1/invalid.json", validTokens: true, validJSON: false},
	}

	run(subtests, t)
}

func TestStep2(t *testing.T) {
	subtests := []subtest{
		{fileName: "tests/step2/valid.json", validTokens: true, validJSON: true},
		{fileName: "tests/step2/valid2.json", validTokens: true, validJSON: true},
		{fileName: "tests/step2/invalid.json", validTokens: true, validJSON: false},
		{fileName: "tests/step2/invalid2.json", validJSON: false},
	}

	run(subtests, t)
}

func TestStep3(t *testing.T) {
	subtests := []subtest{
		{fileName: "tests/step3/valid.json", validTokens: true, validJSON: true},
		{fileName: "tests/step3/invalid.json", validJSON: false},
	}

	run(subtests, t)
}

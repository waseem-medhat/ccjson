package main

import (
	"os"
	"testing"
)

type subtest struct {
	fileName string
	isValid  bool
}

func run(subtests []subtest, t *testing.T) {
	for _, st := range subtests {
		f, err := os.Open(st.fileName)
		if err != nil {
			t.Fatalf("couldn't open file %v", st.fileName)
		}
		defer f.Close()

		err = parseObject(tokenize(f))
		if err != nil && st.isValid {
			t.Fatalf("valid %v got error: %v", st.fileName, err)
		}

		if err == nil && !st.isValid {
			t.Fatalf("invalid %v passed", st.fileName)
		}
	}
}

func TestStep1(t *testing.T) {
	subtests := []subtest{
		{fileName: "tests/step1/valid.json", isValid: true},
		{fileName: "tests/step1/invalid.json", isValid: false},
	}

	run(subtests, t)
}

func TestStep2(t *testing.T) {
	subtests := []subtest{
		{fileName: "tests/step2/valid.json", isValid: true},
		{fileName: "tests/step2/valid2.json", isValid: true},
		{fileName: "tests/step2/invalid.json", isValid: false},
		{fileName: "tests/step2/invalid2.json", isValid: false},
	}

	run(subtests, t)
}

func TestStep3(t *testing.T) {
	subtests := []subtest{
		{fileName: "tests/step3/valid.json", isValid: true},
		{fileName: "tests/step3/invalid.json", isValid: false},
	}

	run(subtests, t)
}

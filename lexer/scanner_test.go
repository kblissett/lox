package lexer

import (
	"reflect"
	"testing"
)

// rightParen
// leftBrace
// rightBrace
// comma
// dot
// minus
// plus
// semicolon
// star

func TestNewLinesAdvanceLineCount(t *testing.T) {
	scanner := Scanner{source: []rune("\n\n\n")}
	scanner.GetTokens()

	want := 4
	if scanner.line != want {
		t.Errorf("Scanner line was not advanced: got %d, wanted %d", scanner.line, want)
	}
}

func TestStringLiterals(t *testing.T) {
	table := []struct {
		source string
		tokens []Token
		errors []error
	}{
		{`"the string"`, []Token{{Kind: stringLiteral, Literal: "the string"}}, nil},
	}

	for _, testCase := range table {
		scanner := Scanner{source: []rune(testCase.source)}
		tokens, errs := scanner.GetTokens()

		if !reflect.DeepEqual(errs, testCase.errors) {
			t.Errorf("Scanner produced incorrect errors: got %v, wanted %v", errs, testCase.errors)
			continue
		}

		if len(tokens) != len(testCase.tokens) {
			t.Errorf("Found the wrong number of tokens: got %d, wanted %d", len(tokens), len(testCase.tokens))
			continue
		}

		if !reflect.DeepEqual(tokens, testCase.tokens) {
			t.Errorf("Failed to parse string literal: got %+v, wanted: %+v", tokens, testCase.tokens)
		}
	}
}

func TestDoesTheTestWork(t *testing.T) {
	table := []struct {
		source string
		tokens []Token
	}{
		{"(", []Token{{Kind: leftParen}}},
		{")", []Token{{Kind: rightParen}}},
		{"{", []Token{{Kind: leftBrace}}},
		{"}", []Token{{Kind: rightBrace}}},
		{",", []Token{{Kind: comma}}},
		{".", []Token{{Kind: dot}}},
		{"-", []Token{{Kind: minus}}},
		{"+", []Token{{Kind: plus}}},
		{";", []Token{{Kind: semicolon}}},
		{"*", []Token{{Kind: star}}},
		{"!", []Token{{Kind: bang}}},
		{"=", []Token{{Kind: equal}}},
		{"<", []Token{{Kind: less}}},
		{">", []Token{{Kind: greater}}},
		{"!=", []Token{{Kind: bangEqual}}},
		{"<=", []Token{{Kind: lessEqual}}},
		{">=", []Token{{Kind: greaterEqual}}},
		{"==", []Token{{Kind: equalEqual}}},
		{"/", []Token{{Kind: slash}}},
		{"// comment", []Token{}},
		{"// comment\n;", []Token{{Kind: semicolon}}},
		{" \t\r\n", []Token{}},
	}

	for _, testCase := range table {
		scanner := Scanner{source: []rune(testCase.source)}
		tokens, _ := scanner.GetTokens()

		if len(tokens) != len(testCase.tokens) {
			t.Errorf("Got the incorrect number of tokens from source %q: got %d, wanted: %d", testCase.source, len(tokens), len(testCase.tokens))
			continue
		}

		if !reflect.DeepEqual(tokens, testCase.tokens) {
			t.Errorf("Tokens from source %q was incorrect: got %v, wanted %v", testCase.source, tokens, testCase.tokens)

		}
	}
}

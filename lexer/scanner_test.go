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
		{`""`, []Token{{Kind: stringLiteral, Literal: ""}}, nil},
		{`"my string`, nil, []error{ScanError{"unterminated string"}}},
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

func TestIdentifiers(t *testing.T) {
	table := map[string]struct {
		source string
		tokens []Token
		errors []error
	}{
		"Can read a simple identifier": {"identifier", []Token{{Kind: identifier, Literal: "identifier"}}, nil},
	}

	for name, tc := range table {
		scanner := Scanner{source: []rune(tc.source)}
		tokens, errs := scanner.GetTokens()

		if !reflect.DeepEqual(errs, tc.errors) {
			t.Errorf("%q: Scanner produced incorrect errors: got %+v, wanted %+v", name, errs, tc.errors)
			continue
		}

		if !reflect.DeepEqual(tokens, tc.tokens) {
			t.Errorf("%q: Failed to parse identifier: got %+v, wanted %+v", name, tokens, tc.tokens)
			continue
		}
	}
}

func TestNumericLiterals(t *testing.T) {
	table := map[string]struct {
		source string
		tokens []Token
		errors []error
	}{
		"Can read a simple numeric literal":   {"1234", []Token{{Kind: number, Literal: "1234"}}, nil},
		"Decimal numbers should not be split": {"12.34", []Token{{Kind: number, Literal: "12.34"}}, nil},
		"Should not create number from method call": {
			"1234.toString()",
			[]Token{
				{Kind: number, Literal: "1234"},
				{Kind: dot, Literal: "."},
				{Kind: identifier, Literal: "toString"},
				{Kind: leftParen, Literal: "("},
				{Kind: rightParen, Literal: ")"},
			}, nil},
	}
	for name, tc := range table {
		scanner := Scanner{source: []rune(tc.source)}
		tokens, errs := scanner.GetTokens()

		if !reflect.DeepEqual(errs, tc.errors) {
			t.Errorf("%s: Scanner produced incorrect errors: got %+v, wanted %+v", name, errs, tc.errors)
			continue
		}

		if !reflect.DeepEqual(tokens, tc.tokens) {
			t.Errorf("%s: Failed to parse identifier: got %+v, wanted %+v", name, tokens, tc.tokens)
			continue
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

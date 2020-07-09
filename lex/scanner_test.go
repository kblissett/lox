package lex

import (
	"reflect"
	"testing"
)

func TestNewLinesAdvanceLineCount(t *testing.T) {
	scanner := Scanner{source: []rune("\n\n\n")}
	scanner.GetTokens()

	want := 4
	if scanner.line != want {
		t.Errorf("Scanner line was not advanced: got %d, wanted %d", scanner.line, want)
	}
}

func TestScanner(t *testing.T) {
	table := map[string]struct {
		source string
		tokens []Token
		errors []error
	}{
		"Can scan left paren":                 {"(", []Token{{Kind: LeftParen, Literal: "("}}, nil},
		"Can scan right paren":                {")", []Token{{Kind: RightParen, Literal: ")"}}, nil},
		"Can scan left brace":                 {"{", []Token{{Kind: LeftBrace, Literal: "{"}}, nil},
		"Can scan right brace":                {"}", []Token{{Kind: RightBrace, Literal: "}"}}, nil},
		"Can scan comma":                      {",", []Token{{Kind: Comma, Literal: ","}}, nil},
		"Can scan dot":                        {".", []Token{{Kind: Dot, Literal: "."}}, nil},
		"Can scan minus":                      {"-", []Token{{Kind: Minus, Literal: "-"}}, nil},
		"Can scan plus":                       {"+", []Token{{Kind: Plus, Literal: "+"}}, nil},
		"Can scan semicolon":                  {";", []Token{{Kind: Semicolon, Literal: ";"}}, nil},
		"Can scan star":                       {"*", []Token{{Kind: Star, Literal: "*"}}, nil},
		"Can scan bang":                       {"!", []Token{{Kind: Bang, Literal: "!"}}, nil},
		"Can scan equal":                      {"=", []Token{{Kind: Equal, Literal: "="}}, nil},
		"Can scan less":                       {"<", []Token{{Kind: Less, Literal: "<"}}, nil},
		"Can scan greater":                    {">", []Token{{Kind: Greater, Literal: ">"}}, nil},
		"Can scan bang equal":                 {"!=", []Token{{Kind: BangEqual, Literal: "!="}}, nil},
		"Can scan less equal":                 {"<=", []Token{{Kind: LessEqual, Literal: "<="}}, nil},
		"Can scan greater equal":              {">=", []Token{{Kind: GreaterEqual, Literal: ">="}}, nil},
		"Can scan equal equal":                {"==", []Token{{Kind: EqualEqual, Literal: "=="}}, nil},
		"Can scan Slash":                      {"/", []Token{{Kind: Slash, Literal: "/"}}, nil},
		"Comments should be ignored":          {"// comment", []Token{}, nil},
		"Can scan tokens after a comment":     {"// comment\n;", []Token{{Kind: Semicolon, Literal: ";"}}, nil},
		"Whitespace should be ignored":        {" \t\r\n", []Token{}, nil},
		"Can read a simple numeric literal":   {"1234", []Token{{Kind: Number, Literal: "1234"}}, nil},
		"Decimal numbers should not be split": {"12.34", []Token{{Kind: Number, Literal: "12.34"}}, nil},
		"Should not create number from method call": {
			"1234.toString()",
			[]Token{
				{Kind: Number, Literal: "1234"},
				{Kind: Dot, Literal: "."},
				{Kind: Identifier, Literal: "toString"},
				{Kind: LeftParen, Literal: "("},
				{Kind: RightParen, Literal: ")"},
			}, nil},
		"Can read a simple identifier":                            {"identifier", []Token{{Kind: Identifier, Literal: "identifier"}}, nil},
		"Can scan a simple string literal":                        {`"the string"`, []Token{{Kind: StringLiteral, Literal: "the string"}}, nil},
		"Can scan an empty string literal":                        {`""`, []Token{{Kind: StringLiteral, Literal: ""}}, nil},
		"Should return error when string literal is unterminated": {`"my string`, nil, []error{ScanError{"unterminated string"}}},
		"Keywords should be recognized":                           {"return", []Token{{Kind: ReturnType, Literal: "return"}}, nil},
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

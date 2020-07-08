package lexer

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
		"Can scan left paren":                 {"(", []Token{{Kind: leftParen, Literal: "("}}, nil},
		"Can scan right paren":                {")", []Token{{Kind: rightParen, Literal: ")"}}, nil},
		"Can scan left brace":                 {"{", []Token{{Kind: leftBrace, Literal: "{"}}, nil},
		"Can scan right brace":                {"}", []Token{{Kind: rightBrace, Literal: "}"}}, nil},
		"Can scan comma":                      {",", []Token{{Kind: comma, Literal: ","}}, nil},
		"Can scan dot":                        {".", []Token{{Kind: dot, Literal: "."}}, nil},
		"Can scan minus":                      {"-", []Token{{Kind: minus, Literal: "-"}}, nil},
		"Can scan plus":                       {"+", []Token{{Kind: plus, Literal: "+"}}, nil},
		"Can scan semicolon":                  {";", []Token{{Kind: semicolon, Literal: ";"}}, nil},
		"Can scan star":                       {"*", []Token{{Kind: star, Literal: "*"}}, nil},
		"Can scan bang":                       {"!", []Token{{Kind: bang, Literal: "!"}}, nil},
		"Can scan equal":                      {"=", []Token{{Kind: equal, Literal: "="}}, nil},
		"Can scan less":                       {"<", []Token{{Kind: less, Literal: "<"}}, nil},
		"Can scan greater":                    {">", []Token{{Kind: greater, Literal: ">"}}, nil},
		"Can scan bang equal":                 {"!=", []Token{{Kind: bangEqual, Literal: "!="}}, nil},
		"Can scan less equal":                 {"<=", []Token{{Kind: lessEqual, Literal: "<="}}, nil},
		"Can scan greater equal":              {">=", []Token{{Kind: greaterEqual, Literal: ">="}}, nil},
		"Can scan equal equal":                {"==", []Token{{Kind: equalEqual, Literal: "=="}}, nil},
		"Can scan slash":                      {"/", []Token{{Kind: slash, Literal: "/"}}, nil},
		"Comments should be ignored":          {"// comment", []Token{}, nil},
		"Can scan tokens after a comment":     {"// comment\n;", []Token{{Kind: semicolon, Literal: ";"}}, nil},
		"Whitespace should be ignored":        {" \t\r\n", []Token{}, nil},
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
		"Can read a simple identifier":                            {"identifier", []Token{{Kind: identifier, Literal: "identifier"}}, nil},
		"Can scan a simple string literal":                        {`"the string"`, []Token{{Kind: stringLiteral, Literal: "the string"}}, nil},
		"Can scan an empty string literal":                        {`""`, []Token{{Kind: stringLiteral, Literal: ""}}, nil},
		"Should return error when string literal is unterminated": {`"my string`, nil, []error{ScanError{"unterminated string"}}},
		"Keywords should be recognized":                           {"return", []Token{{Kind: returnType, Literal: "return"}}, nil},
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

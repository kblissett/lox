package lexer

import (
	"unicode"
)

// ScanError indicates an error during scanning
type ScanError struct {
	text string
}

func (e ScanError) Error() string {
	return e.text
}

// Scanner is a scanner for lox source code.
type Scanner struct {
	source          []rune
	currentPosition int
	line            int
}

func (s *Scanner) match(nextToken rune, ifTrue, ifFalse Token) Token {
	if s.peek() == nextToken {
		s.currentPosition++
		return ifTrue
	}

	return ifFalse
}

func (s *Scanner) peek() rune {
	if s.currentPosition >= len(s.source) {
		return eofRune
	}
	return s.source[s.currentPosition]
}

func (s *Scanner) peekTwo() rune {
	if s.currentPosition+1 >= len(s.source) {
		return eofRune
	}
	return s.source[s.currentPosition+1]
}

func (s *Scanner) advance() rune {
	currentRune := s.peek()
	s.currentPosition++
	return currentRune
}

func (s *Scanner) consumeComment() {
	for c := s.peek(); c != eofRune && c != '\n'; c = s.advance() {
	}
}

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
}

// GetTokens gets the tokens from the source in the scanner.
func (s *Scanner) GetTokens() ([]Token, []error) {
	tokens := []Token{}
	var errors []error
	s.line = 1
	for c := s.advance(); c != eofRune; c = s.advance() {
		switch c {
		case '(':
			tokens = append(tokens, Token{Kind: leftParen, Literal: "("})
		case ')':
			tokens = append(tokens, Token{Kind: rightParen, Literal: ")"})
		case '{':
			tokens = append(tokens, Token{Kind: leftBrace, Literal: "{"})
		case '}':
			tokens = append(tokens, Token{Kind: rightBrace, Literal: "}"})
		case ',':
			tokens = append(tokens, Token{Kind: comma, Literal: ","})
		case '.':
			tokens = append(tokens, Token{Kind: dot, Literal: "."})
		case '-':
			tokens = append(tokens, Token{Kind: minus, Literal: "-"})
		case '+':
			tokens = append(tokens, Token{Kind: plus, Literal: "+"})
		case ';':
			tokens = append(tokens, Token{Kind: semicolon, Literal: ";"})
		case '*':
			tokens = append(tokens, Token{Kind: star, Literal: "*"})
		case '!':
			tokens = append(tokens, s.match('=', Token{Kind: bangEqual, Literal: "!="}, Token{Kind: bang, Literal: "!"}))
		case '=':
			tokens = append(tokens, s.match('=', Token{Kind: equalEqual, Literal: "=="}, Token{Kind: equal, Literal: "="}))
		case '<':
			tokens = append(tokens, s.match('=', Token{Kind: lessEqual, Literal: "<="}, Token{Kind: less, Literal: "<"}))
		case '>':
			tokens = append(tokens, s.match('=', Token{Kind: greaterEqual, Literal: ">="}, Token{Kind: greater, Literal: ">"}))
		case '/':
			if c = s.peek(); c != eofRune && c == '/' {
				s.consumeComment()
			} else {
				tokens = append(tokens, Token{Kind: slash, Literal: "/"})
			}
		case '\n':
			s.line++
		case '"':
			literalChars := []rune{}
			for c = s.advance(); c != eofRune && c != '"'; c = s.advance() {
				literalChars = append(literalChars, c)
			}

			if c == eofRune {
				errors = append(errors, ScanError{"unterminated string"})
			}

			tokens = append(tokens, Token{Kind: stringLiteral, Literal: string(literalChars)})
		default:
			if unicode.IsNumber(c) {
				chars := []rune{c}
				for unicode.IsNumber(s.peek()) || (s.peek() == '.' && unicode.IsNumber(s.peekTwo())) {
					chars = append(chars, s.advance())
				}
				tokens = append(tokens, Token{Kind: number, Literal: string(chars)})
			} else if isAlnum(c) {
				chars := []rune{c}
				for isAlnum(s.peek()) {
					chars = append(chars, s.advance())
				}
				tokens = append(tokens, Token{Kind: identifier, Literal: string(chars)})
			}
		}
	}

	if errors != nil {
		return nil, errors
	}
	return tokens, nil
}

// Token is a lox source code token.
type Token struct {
	Kind    TokenKind
	Literal string
}

// TokenKind is the kind of a token.
type TokenKind string

const eofRune = rune(-1)

const (
	leftParen  TokenKind = "LEFT_PAREN"
	rightParen           = "RIGHT_PAREN"
	leftBrace            = "LEFT_BRACE"
	rightBrace           = "RIGHT_BRACE"
	comma                = "COMMA"
	dot                  = "DOT"
	minus                = "MINUS"
	plus                 = "PLUS"
	semicolon            = "SEMICOLON"
	star                 = "STAR"
	bang                 = "BANG"
	equal                = "EQUAL"
	less                 = "LESS"
	greater              = "GREATER"

	// two character kinds
	bangEqual    = "BANG_EQUAL"
	equalEqual   = "EQUAL_EQUAL"
	lessEqual    = "LESS_EQUAL"
	greaterEqual = "GREATER_EQUAL"
	slash        = "SLASH"

	// literals
	stringLiteral = "STRING_LITERAL"
	number        = "NUMBER_LITERAL"

	// identifier
	identifier = "IDENTIFIER"
)

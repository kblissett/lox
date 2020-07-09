package lex

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
			tokens = append(tokens, Token{Kind: LeftParen, Literal: "("})
		case ')':
			tokens = append(tokens, Token{Kind: RightParen, Literal: ")"})
		case '{':
			tokens = append(tokens, Token{Kind: LeftBrace, Literal: "{"})
		case '}':
			tokens = append(tokens, Token{Kind: RightBrace, Literal: "}"})
		case ',':
			tokens = append(tokens, Token{Kind: Comma, Literal: ","})
		case '.':
			tokens = append(tokens, Token{Kind: Dot, Literal: "."})
		case '-':
			tokens = append(tokens, Token{Kind: Minus, Literal: "-"})
		case '+':
			tokens = append(tokens, Token{Kind: Plus, Literal: "+"})
		case ';':
			tokens = append(tokens, Token{Kind: Semicolon, Literal: ";"})
		case '*':
			tokens = append(tokens, Token{Kind: Star, Literal: "*"})
		case '!':
			tokens = append(tokens, s.match('=', Token{Kind: BangEqual, Literal: "!="}, Token{Kind: Bang, Literal: "!"}))
		case '=':
			tokens = append(tokens, s.match('=', Token{Kind: EqualEqual, Literal: "=="}, Token{Kind: Equal, Literal: "="}))
		case '<':
			tokens = append(tokens, s.match('=', Token{Kind: LessEqual, Literal: "<="}, Token{Kind: Less, Literal: "<"}))
		case '>':
			tokens = append(tokens, s.match('=', Token{Kind: GreaterEqual, Literal: ">="}, Token{Kind: Greater, Literal: ">"}))
		case '/':
			if c = s.peek(); c != eofRune && c == '/' {
				s.consumeComment()
			} else {
				tokens = append(tokens, Token{Kind: Slash, Literal: "/"})
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

			tokens = append(tokens, Token{Kind: StringLiteral, Literal: string(literalChars)})
		default:
			if unicode.IsNumber(c) {
				chars := []rune{c}
				for unicode.IsNumber(s.peek()) || (s.peek() == '.' && unicode.IsNumber(s.peekTwo())) {
					chars = append(chars, s.advance())
				}
				tokens = append(tokens, Token{Kind: Number, Literal: string(chars)})
			} else if isAlnum(c) {
				chars := []rune{c}
				for isAlnum(s.peek()) {
					chars = append(chars, s.advance())
				}
				if tokenType, ok := keywords[string(chars)]; ok {
					tokens = append(tokens, Token{Kind: tokenType, Literal: string(chars)})
				} else {
					tokens = append(tokens, Token{Kind: Identifier, Literal: string(chars)})
				}
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

var keywords = map[string]TokenKind{
	"and":    AndType,
	"class":  ClassType,
	"else":   ElseType,
	"false":  FalseType,
	"for":    ForType,
	"fun":    FunType,
	"if":     IfType,
	"nil":    NilType,
	"or":     OrType,
	"print":  PrintType,
	"return": ReturnType,
	"super":  SuperType,
	"this":   ThisType,
	"true":   TrueType,
	"var":    VarType,
	"while":  WhileType,
}

const (
	LeftParen  TokenKind = "LEFT_PAREN"
	RightParen           = "RIGHT_PAREN"
	LeftBrace            = "LEFT_BRACE"
	RightBrace           = "RIGHT_BRACE"
	Comma                = "COMMA"
	Dot                  = "DOT"
	Minus                = "MINUS"
	Plus                 = "PLUS"
	Semicolon            = "SEMICOLON"
	Star                 = "STAR"
	Bang                 = "BANG"
	Equal                = "EQUAL"
	Less                 = "LESS"
	Greater              = "GREATER"

	// two character kinds
	BangEqual    = "BANG_EQUAL"
	EqualEqual   = "EQUAL_EQUAL"
	LessEqual    = "LESS_EQUAL"
	GreaterEqual = "GREATER_EQUAL"
	Slash        = "SLASH"

	// literals
	StringLiteral = "STRING_LITERAL"
	Number        = "NUMBER_LITERAL"

	// Identifier
	Identifier = "IDENTIFIER"

	// keywords
	AndType    = "AND"
	ClassType  = "CLASS"
	ElseType   = "ELSE"
	FalseType  = "FALSE"
	ForType    = "FOR"
	FunType    = "FUN"
	IfType     = "IF"
	NilType    = "NIL"
	OrType     = "OR"
	PrintType  = "PRINT"
	ReturnType = "RETURN"
	SuperType  = "SUPER"
	ThisType   = "THIS"
	TrueType   = "TRUE"
	VarType    = "VAR"
	WhileType  = "WHILE"
)

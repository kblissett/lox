package lexer

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

func (s *Scanner) match(nextToken rune, ifTrue, ifFalse TokenKind) TokenKind {
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

func (s *Scanner) advance() rune {
	currentRune := s.peek()
	s.currentPosition++
	return currentRune
}

func (s *Scanner) consumeComment() {
	for c := s.peek(); c != eofRune && c != '\n'; c = s.advance() {
	}
}

// GetTokens gets the tokens from the source in the scanner.
func (s *Scanner) GetTokens() ([]Token, []error) {
	tokens := []Token{}
	var errors []error
	s.line = 1
	for c := s.advance(); c != eofRune; c = s.advance() {
		switch c {
		case '(':
			tokens = append(tokens, Token{Kind: leftParen})
		case ')':
			tokens = append(tokens, Token{Kind: rightParen})
		case '{':
			tokens = append(tokens, Token{Kind: leftBrace})
		case '}':
			tokens = append(tokens, Token{Kind: rightBrace})
		case ',':
			tokens = append(tokens, Token{Kind: comma})
		case '.':
			tokens = append(tokens, Token{Kind: dot})
		case '-':
			tokens = append(tokens, Token{Kind: minus})
		case '+':
			tokens = append(tokens, Token{Kind: plus})
		case ';':
			tokens = append(tokens, Token{Kind: semicolon})
		case '*':
			tokens = append(tokens, Token{Kind: star})
		case '!':
			tokens = append(tokens, Token{Kind: s.match('=', bangEqual, bang)})
		case '=':
			tokens = append(tokens, Token{Kind: s.match('=', equalEqual, equal)})
		case '<':
			tokens = append(tokens, Token{Kind: s.match('=', lessEqual, less)})
		case '>':
			tokens = append(tokens, Token{Kind: s.match('=', greaterEqual, greater)})
		case '/':
			if c = s.peek(); c != eofRune && c == '/' {
				s.consumeComment()
			} else {
				tokens = append(tokens, Token{Kind: slash})
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

	// identifier
	identifier = "IDENTIFIER"
)

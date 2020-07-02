package lexer

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
func (s *Scanner) GetTokens() []Token {
	tokens := []Token{}
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

			tokens = append(tokens, Token{Kind: stringLiteral, Literal: string(literalChars)})
		}
	}
	return tokens
}

// Token is a lox source code token.
type Token struct {
	Kind    TokenKind
	Literal string
}

// TokenKind is the kind of a token.
type TokenKind int

const eofRune = rune(-1)

const (
	leftParen TokenKind = iota
	rightParen
	leftBrace
	rightBrace
	comma
	dot
	minus
	plus
	semicolon
	star
	bang
	equal
	less
	greater

	// two character kinds
	bangEqual
	equalEqual
	lessEqual
	greaterEqual
	slash

	// literals
	stringLiteral
)

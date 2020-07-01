package lexer

// Scanner is a scanner for lox source code.
type Scanner struct {
	source          string
	currentPosition int
}

func (s *Scanner) match(nextToken rune, ifTrue, ifFalse TokenKind) TokenKind {
	if s.currentPosition+1 >= len(s.source) {
		return ifFalse
	}

	if []rune(s.source)[s.currentPosition+1] == nextToken {
		s.currentPosition++
		return ifTrue
	}

	return ifFalse
}

// GetTokens gets the tokens from the source in the scanner.
func (s *Scanner) GetTokens() []Token {
	tokens := []Token{}
	for s.currentPosition = 0; s.currentPosition < len([]rune(s.source)); s.currentPosition++ {
		c := []rune(s.source)[s.currentPosition]
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
		}
	}
	return tokens
}

// Token is a lox source code token.
type Token struct {
	Kind TokenKind
}

// TokenKind is the kind of a token.
type TokenKind int

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
)

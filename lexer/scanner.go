package lexer

// Scanner is a scanner for lox source code.
type Scanner struct {
	source string
}

// GetTokens gets the tokens from the source in the scanner.
func (s *Scanner) GetTokens() []Token {
	tokens := []Token{}
	for _, c := range s.source {
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
)

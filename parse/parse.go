package parse

// Expression represents a Lox language expression
type Expression interface {
	String() string
}

// Literal represents a literal value
type Literal struct {
	Value string
}

// String produces a string representation of a literal
func (l Literal) String() string {
	return l.Value
}

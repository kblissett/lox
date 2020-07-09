package parse

import (
	"fmt"

	"github.com/kblissett/lox/lex"
)

// Expression represents a Lox language expression
type Expression interface {
	TreeString() string
}

// Binary represents a Lox binary expression
type Binary struct {
	LeftHandSide  Expression
	Operator      lex.Token
	RightHandSide Expression
}

// TreeString produces a string representation of a binary expression
func (b Binary) TreeString() string {
	return fmt.Sprintf("(%s %s %s)", b.Operator.Literal, b.LeftHandSide.TreeString(), b.RightHandSide.TreeString())
}

// Literal represents a literal value
type Literal struct {
	Value string
}

// TreeString produces a string representation of a literal
func (l Literal) TreeString() string {
	return l.Value
}

// Unary represents a unary expression
type Unary struct {
	Operator      lex.Token
	RightHandSide Expression
}

// TreeString produces a string representation of the tree represented by Unary
func (u Unary) TreeString() string {
	return u.Operator.Literal + u.RightHandSide.TreeString()
}

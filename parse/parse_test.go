package parse

import (
	"testing"

	"github.com/kblissett/lox/lex"
)

func TestPrintTree(t *testing.T) {
	table := map[string]struct {
		expr Expression
		want string
	}{
		"Can print a basic literal expression":  {Literal{"1"}, "1"},
		"Can print a basic unary expression":    {Unary{lex.Token{Kind: lex.Minus, Literal: "-"}, Literal{"1"}}, "-1"},
		"Can print a basic binary expression":   {Binary{Literal{"1"}, lex.Token{Kind: lex.Plus, Literal: "+"}, Literal{"2"}}, "(+ 1 2)"},
		"Can print a basic grouping expression": {Grouping{Literal{"1"}}, "(group 1)"},
	}

	for name, tc := range table {
		got := tc.expr.TreeString()

		if got != tc.want {
			t.Errorf("%s: Produced incorrect tree: Got %q, Wanted %q", name, got, tc.want)
		}
	}
}

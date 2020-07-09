package parse

import (
	"testing"
)

func TestPrintTree(t *testing.T) {
	table := map[string]struct {
		expr Expression
		want string
	}{
		"Can print basic unary expression": {Literal{"1"}, "1"},
	}

	for name, tc := range table {
		got := tc.expr.String()

		if got != tc.want {
			t.Errorf("%s: Produced incorrect tree: Got %q, Wanted %q", name, got, tc.want)
		}
	}
}

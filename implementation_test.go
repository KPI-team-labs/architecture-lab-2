package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type PrefixToInfixSuite struct{}

var _ = Suite(&PrefixToInfixSuite{})

func (s *PrefixToInfixSuite) TestPrefixToInfix(c *C) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"+ 3 4", "(3 + 4)", nil},
		{"* + 2 3 4", "((2 + 3) * 4)", nil},
		{"- / 10 2 3", "((10 / 2) - 3)", nil},
		{"^ 2 + 3 2", "(2 ^ (3 + 2))", nil},
		{"^ - * 1 2 3 / 4 5", "(((1 * 2) - 3) ^ (4 / 5))", nil},
		{"* / 12 3 + 4 5", "((12 / 3) * (4 + 5))", nil},

		{"+ a b", "", fmt.Errorf("letters are not allowed")},
		{"3 4 +", "", fmt.Errorf("invalid prefix expression")},
		{"* 1 2 3 4 5 6 7", "", fmt.Errorf("your prefix expression is not valid")},
		{"$ 1 2", "", fmt.Errorf("your prefix expression is not valid")},
	}

	for _, tc := range testCases {
		result, err := PrefixToInfix(tc.input)
		if err != nil && tc.err != nil {
			c.Check(err.Error(), Equals, tc.err.Error())
		} else {
			c.Check(strings.TrimSpace(result), Equals, tc.expected)
		}
	}
}

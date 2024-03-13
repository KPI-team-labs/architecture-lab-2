package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfix(t *testing.T) {
	res, err := PrefixToInfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

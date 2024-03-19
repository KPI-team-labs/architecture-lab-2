package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (s *MySuite) TestComputeHandler_Compute_Success(c *C) {
	input := strings.NewReader("+ 1 2")
	output := &bytes.Buffer{}
	handler := ComputeHandler{Input: input, Output: output}

	err := handler.Compute()
	c.Assert(err, IsNil)

	expectedOutput := "(1 + 2)"
	c.Assert(output.String(), Equals, expectedOutput)
}

func (s *MySuite) TestComputeHandler_Compute_SyntaxError(c *C) {
	input := strings.NewReader("+ + 1 2")
	output := &bytes.Buffer{}
	handler := ComputeHandler{Input: input, Output: output}

	err := handler.Compute()
	c.Assert(err, NotNil)
}

func (s *MySuite) TestComputeHandler_Compute_EmptyInput(c *C) {
	input := ""
	outputBuffer := &bytes.Buffer{}
	handler := &ComputeHandler{
		Input:  strings.NewReader(input),
		Output: outputBuffer,
	}

	err := handler.Compute()
	c.Assert(err, NotNil)

	expectedOutput := "Syntax error: empty input\n"
	c.Assert(outputBuffer.String(), Equals, expectedOutput)
}

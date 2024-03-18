package lab2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (handler *ComputeHandler) Compute() error {
	reader := bufio.NewReader(handler.Input)
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	result, err := PrefixToInfix(str)
	if err == nil {
		_, err = handler.Output.Write([]byte(result))
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		fmt.Fprintln(handler.Output, err)
	}
	return err
}

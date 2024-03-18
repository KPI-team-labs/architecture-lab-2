package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	lab2 "github.com/KPI-team-labs/architecture-lab-2"
	"os"
)

var (
	inConsoleTag  = flag.String("e", "", "Tag to read the expression from console")
	inputFileTag  = flag.String("f", "", "Tag to read the expression from the input file")
	outputFileTag = flag.String("o", "", "Tag to put the converted expression into the output file")
)

func main() {
	var err error
	var handler lab2.ComputeHandler
	flag.Parse()

	if err := checkTags(inConsoleTag, inputFileTag); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	handler = lab2.ComputeHandler{
		Input:  bytes.NewBufferString(*inConsoleTag),
		Output: os.Stdout,
	}

	if *inConsoleTag == "" && *inputFileTag != "" {
		inFile, err := os.Open(*inputFileTag)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer inFile.Close()

		handler.Input = inFile
	}

	if *outputFileTag != "" {
		outfile, _ := os.OpenFile(*outputFileTag, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
		handler.Output = outfile
		defer outfile.Close()
	}

	err = handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors: %s", err)
	}
}

func checkTags(e, f *string) error {
	if *e == "" && *f == "" {
		return errors.New("you must set or -e tag or -f tag")
	}
	if *e != "" && *f != "" {
		return errors.New("you cannot use -e and -f tags at the same time")
	}
	return nil
}

package main

import (
	"bytes"
	"flag"
	"os"
	"strings"
	"testing"
)

func TestEntriesTextRequired(t *testing.T) {
	// Manipuate the Args to set them up for the testcases
	// After this test we restore the initial args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		name            string
		inputs          map[string]string
		expectedExit    int
		expectedOutputs []string
	}{
		{
			"No inputs",
			map[string]string{},
			1,
			[]string{"Required input 'entries-text' missing or empty"},
		},
		{
			"Just entries-text",
			map[string]string{"entries-text": "foobar"},
			0,
			[]string{},
		},
	}
	for _, tc := range cases {
		// this call is required because otherwise flags panics,
		// if args are set between flag.Parse call
		flag.CommandLine = flag.NewFlagSet(tc.name, flag.ExitOnError)
		// we need a value to set Args[0] to cause flag begins parsing at Args[1]
		args := []string{}
		for name, value := range tc.inputs {
			args = append(args, "-"+name, value)
		}
		os.Args = append([]string{tc.name}, args...)
		var buf bytes.Buffer
		actualExit := realMain(&buf)
		if tc.expectedExit != actualExit {
			t.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				args, tc.expectedExit, actualExit)
		}
		actualOutput := buf.String()
		for _, expected := range tc.expectedOutputs {
			if !strings.Contains(actualOutput, expected) {
				t.Errorf("Wrong output for args: %v, expected %v, got: %v",
					args, expected, actualOutput)
			}
		}
	}
}

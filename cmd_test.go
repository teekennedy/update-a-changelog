package main

import (
	"bytes"
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
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

func TestCli(t *testing.T) {
	assert := assert.New(t)
	// Manipuate the Args to set them up for the testcases
	// After this test we restore the initial args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	for _, tc := range cases {
		// Setup CLI
		// This call is required because otherwise flags panics if args are set
		// between flag.Parse calls
		flag.CommandLine = flag.NewFlagSet(tc.name, flag.ExitOnError)
		// First argument (program name) is unparsed. Fill with test case name
		os.Args = []string{tc.name}
		for name, value := range tc.inputs {
			os.Args = append(os.Args, "-"+name, value)
		}

		var buf bytes.Buffer
		actualExit := realMain(&buf)

		// Assert correct exit code
		assert.Equal(
			tc.expectedExit, actualExit,
			"Wrong exit code for args: %v, expected: %v, got: %v",
			os.Args, tc.expectedExit, actualExit,
		)

		// Assert inputs were assigned correctly
		for name, expected := range tc.inputs {
			actual := *CmdInputs[name].boundVar
			assert.Equal(
				expected, actual,
				"Wrong value assigned to input %v: expected: %v, got: %v. Args: %v",
				name, expected, actual, os.Args,
			)
		}

		// Assert output contains expected messages
		actualOutput := buf.String()
		for _, expected := range tc.expectedOutputs {
			assert.Contains(
				actualOutput, expected,
				"Output does not contain expected message: %v, args: %v",
				expected, os.Args,
			)
		}
	}
}

func TestRunningAsAction(t *testing.T) {
	const envKey = "GITHUB_ACTIONS"
	origEnv, origEnvSet := os.LookupEnv(envKey)
	defer func() {
		if origEnvSet {
			os.Setenv(envKey, origEnv)
		} else {
			os.Unsetenv(envKey)
		}
	}()
	os.Setenv(envKey, "true")
	assert.True(t, RunningAsAction(), "Expected RunningAsAction to be true when %v is true", envKey)
	os.Setenv(envKey, "")
	assert.False(t, RunningAsAction(), "Expected RunningAsAction to be false when %v is empty", envKey)
	os.Unsetenv(envKey)
	assert.False(t, RunningAsAction(), "Expected RunningAsAction to be false when %v is unset", envKey)
}

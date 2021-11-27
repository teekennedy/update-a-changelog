package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(realMain(os.Stdout))
}

// Variables used as inputs in CLI and GitHub Actions parsing
var (
	entriesText   string
	changelogPath string
)

// CmdInputs maps input names to attributes used to populate the input.
// NB: If you make changes to these inputs, be sure to update action.yml as well.
var CmdInputs = map[string]struct {
	boundVar     *string
	description  string
	defaultValue string
	required     bool
}{
	"entries-text": {
		boundVar:     &entriesText,
		description:  "Text to parse as source for new changelog entries",
		defaultValue: "",
		required:     true,
	},
	"changelog-path": {
		boundVar:     &changelogPath,
		description:  "Path to the changelog markdown file to update.",
		defaultValue: "./CHANGELOG.md",
		required:     false,
	},
}

func realMain(out io.Writer) int {
	for name, input := range CmdInputs {
		flag.StringVar(input.boundVar, name, input.defaultValue, input.description)
	}
	flag.Parse()
	for name, input := range CmdInputs {
		if input.required && *input.boundVar == "" {
			fmt.Fprintf(out, "Required input '%v' missing or empty\n", name)
			return 1
		}
	}
	return 0
}

// RunningAsAction returns true if GitHub Actions is running this code.
func RunningAsAction() bool {
	// This is always true when GitHub Actions is running a workflow
	// https://docs.github.com/en/actions/learn-github-actions/environment-variables
	return os.Getenv("GITHUB_ACTIONS") == "true"
}

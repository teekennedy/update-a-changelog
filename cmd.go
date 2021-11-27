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
var entriesText string
var changelogPath string

// PortableInput associates a variable to a CLI flag and GitHub action input
type PortableInput struct {
	boundVar     *string
	name         string
	description  string
	defaultValue string
	required     bool
}

// CmdInputs is an of PortableInput used by this command
// NB: If you make changes to this list, be sure to update action.yml as well.
var CmdInputs = []PortableInput{
	{
		boundVar:     &entriesText,
		name:         "entries-text",
		description:  "Text to parse as source for new changelog entries",
		defaultValue: "",
		required:     true,
	},
	{
		boundVar:     &changelogPath,
		name:         "changelog-path",
		description:  "Path to the changelog markdown file to update.",
		defaultValue: "./CHANGELOG.md",
		required:     false,
	},
}

func realMain(out io.Writer) int {
	for _, input := range CmdInputs {
		flag.StringVar(input.boundVar, input.name, input.defaultValue, input.description)
	}
	flag.Parse()
	for _, input := range CmdInputs {
		if input.required && *input.boundVar == "" {
			fmt.Fprintf(out, "Required input '%v' missing or empty\n", input.name)
			return 1
		}
	}
	return 0
}

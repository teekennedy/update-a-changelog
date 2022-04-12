package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that parser can convert empty content
func testConvertMarkdownEmpty(t *testing.T) {
	buf := ParseEntries([]byte(""))
	// Nothing but the trailing newline
	assert.Equal(t, "\n", buf.String())
}

package godevtools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuffixToFilename(t *testing.T) {
	output := AddSuffixToFilename("foo.bar", "_baz")
	require.Equal(t, "foo_baz.bar", output)
}

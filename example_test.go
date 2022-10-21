package ci_example

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReversing(t *testing.T) {
	s := "hello"
	require.Equal(t, "olleh", ExampleReverse(s))
}

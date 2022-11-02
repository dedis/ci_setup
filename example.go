package ci_example

import (
	"golang.org/x/example/stringutil"
)

// ExampleReverse just for a laugh 
func ExampleReverse(s string) string {
	_ = buhForked()
	return stringutil.Reverse(s)
}

func buhForked() string {
	return "buh"
}

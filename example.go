package ci_example

import (
	"golang.org/x/example/stringutil"
)

// ExampleReverse just for a laugh 
func ExampleReverse(s string) string {
	_ = buh()
	return stringutil.Reverse(s)
}

func buh() string {
	return "buh"
}

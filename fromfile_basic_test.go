package sort_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleSort_fromFile_basic() {
	// sort testdata/unsorted.txt
	gloo.MustRun(
		Sort(gloo.File("testdata/unsorted.txt")),
	)
	// Output:
	// apple
	// banana
	// cherry
}

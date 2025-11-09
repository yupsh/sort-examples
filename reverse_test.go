package sort_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_reverse() {
	// echo "apple\nbanana\ncherry" | sort -r
	gloo.MustRun(
		Sort(Reverse, strings.NewReader("apple\nbanana\ncherry")),
	)
	// Output:
	// cherry
	// banana
	// apple
}

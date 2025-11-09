package sort_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_unique() {
	// echo "apple\nbanana\napple\ncherry\nbanana" | sort -u
	gloo.MustRun(
		Sort(Unique, strings.NewReader("apple\nbanana\napple\ncherry\nbanana")),
	)
	// Output:
	// apple
	// banana
	// cherry
}

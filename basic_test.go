package sort_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_basic() {
	// echo "cherry\napple\nbanana" | sort
	gloo.MustRun(
		Sort(strings.NewReader("cherry\napple\nbanana")),
	)
	// Output:
	// apple
	// banana
	// cherry
}

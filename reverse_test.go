package sort_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_reverse() {
	// echo "apple\nbanana\ncherry" | sort -r
	yup.MustRun(
		Sort(Reverse, strings.NewReader("apple\nbanana\ncherry")),
	)
	// Output:
	// cherry
	// banana
	// apple
}


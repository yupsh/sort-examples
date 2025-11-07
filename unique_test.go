package sort_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_unique() {
	// echo "apple\nbanana\napple\ncherry\nbanana" | sort -u
	yup.MustRun(
		Sort(Unique, strings.NewReader("apple\nbanana\napple\ncherry\nbanana")),
	)
	// Output:
	// apple
	// banana
	// cherry
}


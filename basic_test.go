package sort_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_basic() {
	// echo "cherry\napple\nbanana" | sort
	yup.MustRun(
		Sort(strings.NewReader("cherry\napple\nbanana")),
	)
	// Output:
	// apple
	// banana
	// cherry
}


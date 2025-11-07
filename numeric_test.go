package sort_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sort"
)

func ExampleSort_numeric() {
	// echo "10\n2\n100\n20" | sort -n
	yup.MustRun(
		Sort(Numeric, strings.NewReader("10\n2\n100\n20")),
	)
	// Output:
	// 2
	// 10
	// 20
	// 100
}


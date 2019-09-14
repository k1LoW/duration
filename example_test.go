package duration_test

import (
	"fmt"

	"github.com/k1LoW/duration"
)

func ExampleParse() {
	d, _ := duration.Parse("3 days 4 hours")
	fmt.Printf("%s", d)
	// Output: 76h0m0s
}

package duration_test

import (
	"fmt"
	"log"

	"github.com/k1LoW/duration"
)

func ExampleParse() {
	d, err := duration.Parse("3 days 4 hours")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", d)
	// Output: 76h0m0s
}

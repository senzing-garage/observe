package main

import (
	"testing"
)

func TestMain(test *testing.T) {
	test.Setenv("SENZING_TOOLS_AVOID_SERVING", "true")
	main()
}

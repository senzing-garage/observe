package main

import (
	"testing"
)

/*
 * The unit tests in this file simulate command line invocation.
 */
func TestMain(test *testing.T) {
	_ = test
	test.Setenv("SENZING_TOOLS_AVOID_SERVING", "true")
	main()
}

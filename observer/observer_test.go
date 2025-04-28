package observer_test

import (
	"testing"

	"github.com/senzing-garage/observe/observer"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSimpleObserver_Serve(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := &observer.SimpleObserver{
		AvoidServing: true,
	}
	err := testObject.Serve(ctx)
	require.NoError(test, err)
}

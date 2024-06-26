package observer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSimpleObserver_Serve(test *testing.T) {
	_ = test
	ctx := context.TODO()
	testObject := &SimpleObserver{
		AvoidServing: true,
	}
	err := testObject.Serve(ctx)
	require.NoError(test, err)
}

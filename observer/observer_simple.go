package observer

import (
	"context"
	"fmt"
	"os"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-observing/grpcserver"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SimpleObserver is an ObserverInterface.
type SimpleObserver struct {
	AvoidServing  bool
	Port          int
	ServerOptions []grpc.ServerOption
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The SaySomething method simply prints the 'Something' value in the type-struct.

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/
func (observerImpl *SimpleObserver) Serve(ctx context.Context) error {
	// Create a Subject.
	aSubject := &subject.SimpleSubject{}

	// Register an observer with the Subject.

	anObserver := &observer.RawObserver{
		ID: "observe",
	}

	err := aSubject.RegisterObserver(ctx, anObserver)
	if err != nil {
		return wraperror.Errorf(err, "observer.Serve.RegisterObserver error: %w", err)
	}

	// Run an Observer gRPC service.

	aGrpcServer := &grpcserver.SimpleGrpcServer{
		Port:          observerImpl.Port,
		ServerOptions: observerImpl.ServerOptions,
		Subject:       aSubject,
	}

	fmt.Printf(">>>>>>> SENZING_TOOLS_AVOID_SERVING: %s\n", os.Getenv("SENZING_TOOLS_AVOID_SERVING"))
	fmt.Printf(">>>>>>> observerImpl.AvoidServing %t\n", observerImpl.AvoidServing)

	if !observerImpl.AvoidServing {
		err = aGrpcServer.Serve(ctx)
	}

	fmt.Printf(">>>>>>> observerImpl.AvoidServing done\n")

	return wraperror.Errorf(err, "observer.Serve error: %w", err)
}

package observer

import (
	"context"
	"fmt"

	"github.com/senzing/go-observing/grpcserver"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ObserverImpl is an ObserverInterface.
type ObserverImpl struct {
	Port int
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
func (observerImpl *ObserverImpl) Serve(ctx context.Context) error {
	var err error = nil

	// Create a Subject.

	aSubject := &subject.SubjectImpl{}

	// Register an observer.

	anObserver := &observer.ObserverNull{
		Id: "observe",
	}

	// Run an Observer gRPC service.

	err = aSubject.RegisterObserver(ctx, anObserver)
	if err != nil {
		fmt.Print(err)
	}

	aGrpcServer := &grpcserver.GrpcServerImpl{
		Port:    observerImpl.Port,
		Subject: aSubject,
	}
	aGrpcServer.Serve(ctx)

	return err
}

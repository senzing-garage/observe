package observer

import (
	"context"

	"github.com/senzing-garage/go-observing/grpcserver"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ObserverImpl is an ObserverInterface.
type ObserverImpl struct {
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
func (observerImpl *ObserverImpl) Serve(ctx context.Context) error {
	// Create a Subject.

	aSubject := &subject.SubjectImpl{}

	// Register an observer with the Subject.

	anObserver := &observer.ObserverRaw{
		Id: "observe",
	}

	err := aSubject.RegisterObserver(ctx, anObserver)
	if err != nil {
		return err
	}

	// Run an Observer gRPC service.

	aGrpcServer := &grpcserver.GrpcServerImpl{
		Port:          observerImpl.Port,
		ServerOptions: observerImpl.ServerOptions,
		Subject:       aSubject,
	}
	err = aGrpcServer.Serve(ctx)

	return err
}

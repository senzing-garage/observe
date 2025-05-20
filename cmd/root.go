/*
 */
package cmd

import (
	"context"
	"os"

	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/observe/observer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	Short string = "Aggregate observations"
	Use   string = "observe"
	Long  string = `
Listen for Observer messages over gRPC and print them to STDOUT.
    `
)

var avoidServe = option.ContextVariable{
	Arg:     "avoid-serving",
	Default: option.OsLookupEnvBool("SENZING_TOOLS_AVOID_SERVING", false),
	Envar:   "SENZING_TOOLS_AVOID_SERVING",
	Help:    "Avoid serving.  For testing only. [%s]",
	Type:    optiontype.Bool,
}

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var ContextVariablesForMultiPlatform = []option.ContextVariable{
	option.ObserverGrpcPort,
	option.LogLevel,
	avoidServe,
}

var ContextVariables = append(ContextVariablesForMultiPlatform, ContextVariablesForOsArch...)

// ----------------------------------------------------------------------------
// Command
// ----------------------------------------------------------------------------

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:     Use,
	Short:   Short,
	Long:    Long,
	PreRun:  PreRun,
	RunE:    RunE,
	Version: Version(),
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used in construction of cobra.Command.
func PreRun(cobraCommand *cobra.Command, args []string) {
	cmdhelper.PreRun(cobraCommand, args, Use, ContextVariables)
}

// Used in construction of cobra.Command.
func RunE(_ *cobra.Command, _ []string) error {
	ctx := context.Background()

	// IMPROVE: Support various gRPC server options.

	serverOptions := []grpc.ServerOption{}

	// Create and run gRPC server.

	observer := &observer.SimpleObserver{
		AvoidServing:  viper.GetBool(avoidServe.Arg),
		Port:          viper.GetInt(option.ObserverGrpcPort.Arg),
		ServerOptions: serverOptions,
	}

	if err := observer.Serve(ctx); err != nil {
		return wraperror.Errorf(err, "Serve")
	}

	return nil
}

// Used in construction of cobra.Command.
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	cmdhelper.Init(RootCmd, ContextVariables)
}

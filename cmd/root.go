/*
 */
package cmd

import (
	"context"
	"os"

	"github.com/senzing/observe/observer"
	"github.com/senzing/senzing-tools/cmdhelper"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/help"
	"github.com/senzing/senzing-tools/option"
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

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var ContextInts = []cmdhelper.ContextInt{
	{
		Default: cmdhelper.OsLookupEnvInt(envar.GrpcPort, 8260),
		Envar:   envar.GrpcPort,
		Help:    help.GrpcPort,
		Option:  option.GrpcPort,
	},
}

var ContextStrings = []cmdhelper.ContextString{
	{
		Default: cmdhelper.OsLookupEnvString(envar.LogLevel, "INFO"),
		Envar:   envar.LogLevel,
		Help:    help.LogLevel,
		Option:  option.LogLevel,
	},
}

var ContextVariables = &cmdhelper.ContextVariables{
	Ints:    ContextInts,
	Strings: ContextStrings,
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	cmdhelper.Init(RootCmd, *ContextVariables)
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

// Used in construction of cobra.Command
func PreRun(cobraCommand *cobra.Command, args []string) {
	cmdhelper.PreRun(cobraCommand, args, Use, *ContextVariables)
}

// Used in construction of cobra.Command
func RunE(_ *cobra.Command, _ []string) error {
	var err error = nil
	ctx := context.Background()

	// TODO: Support various gRPC server options.

	serverOptions := []grpc.ServerOption{}

	// Create and run gRPC server.

	observer := &observer.ObserverImpl{
		Port:          viper.GetInt(option.GrpcPort),
		ServerOptions: serverOptions,
	}
	err = observer.Serve(ctx)
	return err
}

// Used in construction of cobra.Command
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

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

/*
 */
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/senzing/observe/observer"
	"github.com/senzing/senzing-tools/constant"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/helper"
	"github.com/senzing/senzing-tools/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	defaultLogLevel         string = "INFO"
	defaultObserverGrpcPort int    = 8260
	Short                   string = "Aggregate observations"
	Use                     string = "observe"
	Long                    string = `
Listen for Observer messages over gRPC and print them to STDOUT.
	`
)

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	RootCmd.Flags().Int("observer-grpc-port", defaultObserverGrpcPort, fmt.Sprintf("Port used to listen to Observer messages over gRPC [%s]", "SENZING_TOOLS_OBSERVER_GRPC_PORT"))
	RootCmd.Flags().String(option.LogLevel, defaultLogLevel, fmt.Sprintf("Log level of TRACE, DEBUG, INFO, WARN, ERROR, FATAL, or PANIC [%s]", envar.LogLevel))
	// RootCmd.Flags().Int(option.ObserverGrpcPort, defaultObserverGrpcPort, fmt.Sprintf("Port used to listen to Observer messages over gRPC [%s]", envar.ObserverGrpcPort))
}

// If a configuration file is present, load it.
func loadConfigurationFile(cobraCommand *cobra.Command) {
	configuration := ""
	configFlag := cobraCommand.Flags().Lookup(option.Configuration)
	if configFlag != nil {
		configuration = configFlag.Value.String()
	}
	if configuration != "" { // Use configuration file specified as a command line option.
		viper.SetConfigFile(configuration)
	} else { // Search for a configuration file.

		// Determine home directory.

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Specify configuration file name.

		viper.SetConfigName("observe")
		viper.SetConfigType("yaml")

		// Define search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Applying configuration file:", viper.ConfigFileUsed())
	}
}

// Configure Viper with user-specified options.
func loadOptions(cobraCommand *cobra.Command) {
	var err error = nil
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	// Ints

	intOptions := map[string]int{
		"observer-grpc-port": defaultObserverGrpcPort,
		// option.ObserverGrpcPort: defaultObserverGrpcPort,
	}
	for optionKey, optionValue := range intOptions {
		viper.SetDefault(optionKey, optionValue)
		err = viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
		if err != nil {
			panic(err)
		}
	}

	// Strings

	stringOptions := map[string]string{
		option.LogLevel: defaultLogLevel,
	}
	for optionKey, optionValue := range stringOptions {
		viper.SetDefault(optionKey, optionValue)
		err = viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
		if err != nil {
			panic(err)
		}
	}
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
	loadConfigurationFile(cobraCommand)
	loadOptions(cobraCommand)
	cobraCommand.SetVersionTemplate(constant.VersionTemplate)
}

// Used in construction of cobra.Command
func RunE(_ *cobra.Command, _ []string) error {
	var err error = nil
	ctx := context.TODO()

	// TODO: Support various gRPC server options.

	serverOptions := []grpc.ServerOption{}

	// Create and run gRPC server.

	observer := &observer.ObserverImpl{
		Port: viper.GetInt("observer-grpc-port"),
		// Port:          viper.GetInt(option.ObserverGrpcPort),
		ServerOptions: serverOptions,
	}
	err = observer.Serve(ctx)
	return err
}

// Used in construction of cobra.Command
func Version() string {
	return helper.MakeVersion(githubVersion, githubIteration)
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

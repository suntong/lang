////////////////////////////////////////////////////////////////////////////
// Porgram: CommandLineCV
// Purpose: Go commandline via cobra & viper demo
// Authors: Tong Sun (c) 2015, All rights reserved
// based on https://github.com/chop-dbhi/origins-dispatch/blob/master/main.go
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const version = "0.1.0"

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

/*

$ CommandLineCV help
HTTP service that consumes events and dispatches them to subscribers.

Usage:
  dispatch [flags]
  dispatch [command]

Available Commands:
  version     Print the version.
  help        Help about any command

Flags:
      --addr="localhost:5002": Address of the service
      --debug=false: Turn on debugging.
      --email-from="noreply@example.com": The from email address.
  -h, --help=false: help for dispatch
      --smtp-addr="localhost:25": Address of the SMTP server
      --smtp-password="": Password to authenticate with the SMTP server
      --smtp-user="": User to authenticate with the SMTP server


Use "dispatch help [command]" for more information about a command.

*/

// The main command describes the service and defaults to printing the
// help message.
var mainCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Event dispatch service.",
	Long:  `HTTP service that consumes events and dispatches them to subscribers.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

// The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the dispatch service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

// Go special automatically executed init function
func init() {
	mainCmd.AddCommand(versionCmd)

	viper.SetEnvPrefix("DISPATCH")
	viper.AutomaticEnv()

	/*

	  When AutomaticEnv called, Viper will check for an environment variable any
	  time a viper.Get request is made. It will apply the following rules. It
	  will check for a environment variable with a name matching the key
	  uppercased and prefixed with the EnvPrefix if set.

	*/

	flags := mainCmd.Flags()

	flags.Bool("debug", false, "Turn on debugging.")
	flags.String("addr", "localhost:5002", "Address of the service")
	flags.String("smtp-addr", "localhost:25", "Address of the SMTP server")
	flags.String("smtp-user", "", "User to authenticate with the SMTP server")
	flags.String("smtp-password", "", "Password to authenticate with the SMTP server")
	flags.String("email-from", "noreply@example.com", "The from email address.")

	viper.BindPFlag("debug", flags.Lookup("debug"))
	viper.BindPFlag("addr", flags.Lookup("addr"))
	viper.BindPFlag("smtp_addr", flags.Lookup("smtp-addr"))
	viper.BindPFlag("smtp_user", flags.Lookup("smtp-user"))
	viper.BindPFlag("smtp_password", flags.Lookup("smtp-password"))
	viper.BindPFlag("email_from", flags.Lookup("email-from"))
}

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	mainCmd.Execute()
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func serve() {
	debug := viper.GetBool("debug")
	addr := viper.GetString("addr")

	fmt.Printf("* Serving on http://%s\n", addr)

	if debug {
		fmt.Println("* Debugging enabled")
	}

}

/*

$ go run CommandLineCV.go
* Serving on http://localhost:5002

$ go run CommandLineCV.go --addr="localhost:5005"
* Serving on http://localhost:5005

$ DISPATCH_ADDR="localhost:6006" go run CommandLineCV.go
* Serving on http://localhost:6006

$ DISPATCH_ADDR="localhost:6006" go run CommandLineCV.go --addr="localhost:5005"
* Serving on http://localhost:5005

$ DISPATCH_DEBUG=True DISPATCH_ADDR="localhost:6006" go run CommandLineCV.go --addr="localhost:5005"
* Serving on http://localhost:5005
* Debugging enabled

*/

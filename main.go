package main

import (
	flags "github.com/jessevdk/go-flags"
	"log"
	"os"
)

// opts is a struct of expected command line flags that are parsed upon
// startup.
var opts struct {
	verbose bool `short:"v" long:"verbose" description:"Show verbose debug information" default:"false"`
}

// args is an array of command line arguments that are parsed upon startup.
var args []string

func main() {
	log.Printf("fried-gophers is live!\nWas given the following command line "+
		"arguments: %v", args)
}

// init is automatically called upon program / package startup, before even
// main() is run.
//
// Perform initialization actions here.
func init() {
	// Parse command line flags
	args = parseFlags()

	// Setup internal logger
	log.SetFlags(0) // Do not log time information
}

// parseFlags reads command line flags into 'data', and returns the remaining
// args.
func parseFlags() (args []string) {
	var err error
	args, err = flags.Parse(&opts)
	if err != nil {
		flagErr := err.(*flags.Error)
		switch flagErr.Type {
		// Do not error out upon a usage error.
		case flags.ErrHelp, flags.ErrUnknownFlag:
			os.Exit(0)
		default:
			log.Fatalf("Could not parse flags: %v", err)
		}
	}
	return
}

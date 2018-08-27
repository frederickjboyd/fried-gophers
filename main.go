package main

import (
	flags "github.com/jessevdk/go-flags"
	"log"
	"os"
)

// opts is a struct of expected command line flags that are parsed upon
// startup.
var opts struct {
	Verbose    bool    `short:"v" long:"verbose" description:"Show verbose debug information"`
	Output     string  `short:"o" long:"output" description:"Rename file output name"`
	Saturation float64 `short:"s" long:"saturation" description:"Adjust saturation of output" default:"5"`
}

// args is an array of command line arguments that are parsed upon startup.
var args []string

func main() {
	args := parseFlags()
	img, format := openImage(args[0])
	if verbose := opts.Verbose; verbose {
		log.Printf("arguments: %s", args)
		log.Printf("output: %s", opts.Output)
		log.Printf("saturation: %f", opts.Saturation)
		log.Printf("input image format: %s", format)
	}
	friedImg := adjustSaturation(img, opts.Saturation)
	// noiseImg := genNoise()
	if outputLength := len(opts.Output); outputLength > 0 {
		writeImage(friedImg, opts.Output+".jpg")
	} else {
		writeImage(friedImg, "deep-fried.jpg")
	}
	// writeImage(noiseImg, "noise.jpg")
}

// init is automatically called upon program / package startup, before even
// main() is run.
//
// Perform initialization actions here.
func init() {
	// Setup internal logger, do not log time data
	log.SetFlags(0)

	// Parse command line flags
	if args = parseFlags(); len(args) == 0 {
		log.Fatal("Expected an file path argument; received none.")
	}
}

// parseFlags reads command line flags into 'data', and returns the remaining
// args.
func parseFlags() (args []string) {
	var err error
	// arguments := []string{
	// 	"-v",
	// 	"-o", "file.exe",
	// 	"-s", "hello",
	// 	"-s", "world",
	// }
	// for i := 0; i < len(arguments); i++ {
	// 	log.Println(arguments[i])
	// }
	// args, err = flags.ParseArgs(&opts, arguments)
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

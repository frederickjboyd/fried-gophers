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
	Directory  string  `short:"d" long:"dir" description:"Directory to write image to (writes to target image directory by default)"`
	Output     string  `short:"o" long:"output" description:"Rename file output name" default:"deep-fried"`
	Saturation float64 `short:"s" long:"saturation" description:"Adjust saturation of output" default:"50"`
	Quality    int     `short:"q" long:"quality" description:"Adjust quality of output" default:"5"`
}

// args is an array of command line arguments that are parsed upon startup.
var args []string

func main() {
	args := parseFlags()
	img, format := openImage(args[0])
	if opts.Verbose {
		log.Printf("target image: %s", args[0])
		log.Printf("image format: %s", format)
		log.Printf("output file name: %s", opts.Output)
		log.Printf("saturation: %f", opts.Saturation)
		log.Printf("quality: %d", opts.Quality)
	}
	friedImg := adjustSaturation(img, opts.Saturation)
	writeImage(friedImg, opts.Output+".jpg", opts.Directory, opts.Quality)

	// noiseImg := genNoise()
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

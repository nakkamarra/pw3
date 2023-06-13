package arguments

import (
	"flag"

	"github.com/nakkamarra/pw3/algorithm"
)

// Arguments represents the potential arguments passed to this program
// via the command line.
type Arguments struct {
	Algorithm algorithm.Algorithm
	Base64    bool
	Help      bool
	Length    uint
}

// Bind reads in an array of arguments and attempts to parse and store
// them in an instance of Arguments.
func Bind() *Arguments {
	var alg = flag.String("alg", "sha256", "Algorithm to use; i.e sha256, sha512, md5")
	var base64 = flag.Bool("base64", false, "Encode output as base64")
	var help = flag.Bool("help", false, "Print out helpful information about the program")
	var length = flag.Uint("length", 0, "Specify length of output")
	flag.Parse()
	args := new(Arguments)
	args.Algorithm = algorithm.Algorithm(*alg)
	args.Base64 = *base64
	args.Length = *length
	args.Help = *help
	return args
}

// PrintDefaults calls to flag to print the defaults
func PrintDefaults() {
	flag.PrintDefaults()
}

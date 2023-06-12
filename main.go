package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"golang.org/x/term"
	"os"
	"reflect"
	"syscall"
)

func main() {
	argv := os.Args[1:]
	arguments := Bind(argv)
	if arguments.Help == true {
		fmt.Fprintf(os.Stderr, "PW3: a tool for generating encrypted passphrases\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n\t pw3 <flags...>\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n\n")
		fmt.Fprintf(os.Stderr, "%s\n", getFlags(*arguments))
		fmt.Fprintln(os.Stderr, "Try using --help for more info")
		os.Exit(1)
	}
	if arguments.Error != nil {
		fmt.Fprintln(os.Stdout, arguments.Error)
	}

	fmt.Fprint(os.Stderr, "Enter passphrase: ")
	bytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't read passphrase")
		os.Exit(1)
	}

	var sum []byte
	switch arguments.Algorithm {
	case MD5:
		sum = md5.Sum(bytes)[:]
	case SHA512:
		sum = sha512.Sum512(bytes)[:]
	case SHA256:
	default:
		sum = sha256.Sum256(bytes)[:]
	}

	if arguments.Base64 == true {
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
		base64.StdEncoding.Encode(dst, sum)
		bytes = dst
	}

	fmt.Fprintln(os.Stdout)
	fmt.Fprintf(os.Stdout, "%x\n", sum)
}

type Algorithm string

const (
	SHA256 Algorithm = "sha256" // Default
	SHA512           = "sha512"
	MD5              = "md5"
)

// Arguments represents the potential arguments passed to this program
// via the command line.
type Arguments struct {
	Algorithm string `short:"-a" long:"--alg" desc:"Algorithm to use; i.e sha256, sha512, md5"`
	Base64    bool   `short:"-b" long:"--base64" desc:"Encode output as base64"`
	Help      bool   `short:"-h" long:"--help" desc:"Print out helpful information about the program"`
	Length    uint   `short:"-l" long:"--long" desc:"Specify length / trim range; i.e --length=20 or -l 10"`
	Error     error
}

// Bind reads in an array of arguments and attempts to parse and store
// them in an instance of Arguments.
func Bind(arguments []string) *Arguments {
	args := new(Arguments)
	for index, arg := range arguments {
		switch arg {
		case "-a":
		case "--alg":
			args.Algorithm = arg
		case "-b":
		case "--base64":
			args.Base64 = true
		case "-l":
		case "--length":
			args.Length = arg
		case "-h":
		case "--help":
			args.Help = true
		default:
			args.Error = fmt.Errorf("Unknown argument provided: %s\n", arg)
		}
	}
	return args
}

// getFlags uses reflection to read the custom struct tags used by this program
// and return them formatted as a string.
func getFlags(argv interface{}) string {
	output := ""
	a := reflect.TypeOf(argv)
	for i := 0; i < a.NumField(); i++ {
		field := a.Field(i)

		short := field.Tag.Get("short")
		long := field.Tag.Get("long")
		description := field.Tag.Get("desc")
		if short != "" && long != "" {
			output += fmt.Sprintf("\t%s,%s\t%s\n", short, long, description)
		}
	}

	return output
}

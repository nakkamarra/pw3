package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	argv := os.Args[1:]
	arguments := Bind(argv)
	if len(argv) == 0 {
		fmt.Fprintln(os.Stderr, "No argument...")
		fmt.Fprintln(os.Stderr, "Try using --help for more info")
		os.Exit(1)
	}
	if arguments.Help == true {
		fmt.Fprintf(os.Stdout, "PW3: a tool for generating encrypted passphrases\n\n")
		fmt.Fprintf(os.Stdout, "Usage:\n\n\t pw3 <flags...> [passphrase]\n\n")
		fmt.Fprintf(os.Stdout, "Flags:\n\n")
		fmt.Fprintf(os.Stdout, getFlags(*arguments))
		os.Exit(1)
	}

}

type Arguments struct {
	Base64  bool   `short:"-b" long:"--base64" desc:"Encode output as base64"`
	Help    bool   `short:"-h" long:"--help" desc:"Print out helpful information about the program"`
	Length  string `short:"-l" long:"--long" desc:"Specify length / trim range; i.e --length=20 or -l 10"`
	Payload string
}

func Bind(arguments []string) *Arguments {
	args := new(Arguments)
	for _, arg := range arguments {
		switch arg {
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
			args.Payload = arg
		}
	}
	return args
}

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

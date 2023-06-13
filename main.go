package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/term"
	"os"
	"syscall"

	"github.com/nakkamarra/pw3/algorithm"
	"github.com/nakkamarra/pw3/arguments"
)

func main() {
	args := arguments.Bind()
	if args.Help == true {
		fmt.Fprintf(os.Stderr, "PW3: a tool for generating encrypted passphrases\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n   pw3 <flags...>\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n\n")
		arguments.PrintDefaults()
		os.Exit(1)
	}

	fmt.Fprint(os.Stderr, "Enter passphrase: ")
	bytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't read passphrase")
		os.Exit(1)
	}

	sum := algorithm.GetSum(args.Algorithm, bytes)

	if args.Base64 == true {
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
		base64.StdEncoding.Encode(dst, sum)
		sum = dst
	}

	if args.Length != 0 {
		length := args.Length
		sumLength := uint(len(sum))
		if length > sumLength {
			length = sumLength
		}
		sum = sum[:length]
	}

	fmt.Fprintln(os.Stdout)
	fmt.Fprintf(os.Stdout, "%x\n", sum)
}

package main

import (
	"bufio"
	"encoding/hex"
	"encoding/base64"
	"golang.org/x/term"
	"os"
	"syscall"

	"github.com/nakkamarra/pw3/algorithm"
	"github.com/nakkamarra/pw3/arguments"
)

func main() {
	stdout := bufio.NewWriter(os.Stdout)
	stderr := bufio.NewWriter(os.Stderr)

	args := arguments.Bind()
	if args.Help == true {
		stderr.WriteString("PW3: a tool for generating encrypted passphrases\n\n")
		stderr.WriteString("Usage:\n\n   pw3 <flags...>\n\n")
		stderr.WriteString("Flags:\n\n")
		stderr.Flush()
		arguments.PrintDefaults()
		os.Exit(1)
	}

	stderr.WriteString("Enter passphrase: ")
	stderr.Flush()
	bytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		stderr.WriteString("Couldn't read passphrase\n")
		stderr.Flush()
		os.Exit(1)
	}

	sum := algorithm.GetSum(args.Algorithm, bytes)
	hexString := hex.EncodeToString(sum)

	if args.Base64 == true {
		bytes := []byte(hexString)
		length := base64.StdEncoding.EncodedLen(len(bytes))
		dst := make([]byte, length) 
		base64.StdEncoding.Encode(dst, bytes)
		hexString = string(dst)
	}

	if args.Length != 0 {
		length := args.Length
		sumLength := uint(len(hexString))
		if length > sumLength {
			length = sumLength
		}
		hexString = hexString[:length]
	}

	stdout.WriteString("\n")
	stdout.WriteString(hexString + "\n")
	stdout.Flush()
}

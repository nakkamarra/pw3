# pw3

A simple cli for writing hashes from input passphrases.

## Installation

```bash
go install github.com/nakkamarra/pw3@latest
```

## Usage

For standard (sha256 by default):
```bash
pw3
```

For limiting the number of bytes written (length of 10 = 20 hex digits):
```bash
pw3 --length=10
```

For base64-encoded MD5:
```bash
pw3 --alg=md5 --base64
```

For more info:
```bash
pw3 --help
```

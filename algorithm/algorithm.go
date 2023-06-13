package algorithm

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
)

type Algorithm string

const (
	SHA256 Algorithm = "sha256" // Default
	SHA512           = "sha512"
	MD5              = "md5"
)

// GetSum takes an algorithm type and a slice of bytes
// and returns the result of calling Sum on the bytes
// for the given algorithm. Default is SHA256.
func GetSum(alg Algorithm, bytes []byte) []byte {
	switch alg {
	case MD5:
		digest := md5.Sum(bytes)
		return digest[:]
	case SHA512:
		digest := sha512.Sum512(bytes)
		return digest[:]
	case SHA256:
		fallthrough
	default:
		digest := sha256.Sum256(bytes)
		return digest[:]
	}
}

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Setting flag
	var (
		hash = flag.String("hash", "sha256", "This flag is use hash protocol.")
	)
	flag.Parse()

	// stdin
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()

	// input to hash
	sha := makeHash(buf.Text(), *hash)
	fmt.Printf("%x\n", binary.BigEndian.Uint64(sha))
}

func makeHash(value string, hash string) []byte {
	var sha []byte
	if hash == "sha256" {
		s := makeHashSHA256("a")
		sha = s[:]
	} else if hash == "sha384" {
		s := makeHashSHA384("a")
		sha = s[:]
	} else if hash == "sha512" {
		s := makeHashSHA512("a")
		sha = s[:]
	}
	return sha
}

func makeHashSHA256(value string) [32]byte {
	return sha256.Sum256([]byte(value))
}

func makeHashSHA384(value string) [48]byte {
	return sha512.Sum384([]byte(value))
}

func makeHashSHA512(value string) [64]byte {
	return sha512.Sum512([]byte(value))
}

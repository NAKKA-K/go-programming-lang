package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	fmt.Println(countEqualBitSHA(s1, s2))
}

func countEqualBitSHA(sha1 [32]byte, sha2 [32]byte) int {
	andBit := binary.BigEndian.Uint32(sha1[:]) & binary.BigEndian.Uint32(sha2[:])
	var cnt int
	for ; andBit > 0; cnt++ {
		andBit &= andBit - 1
	}
	return cnt
}

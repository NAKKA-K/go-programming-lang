package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func compressionSpace(slice []byte) []byte {
	_, size := utf8.DecodeLastRune(slice)
	for i := 0; i < len(slice)-size; {
		r1, size1 := utf8.DecodeRune(slice[i:])
		r2, size2 := utf8.DecodeRune(slice[i+size1:])
		fmt.Println(string(r1), string(r2))
		if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
			slice = append(slice[:i+size1], slice[i+size1+size2:]...)
			continue
		}
		i += size1
	}
	return slice
}

func main() {
	bytes := []byte("あい　　　　　　　　お")
	bytes = compressionSpace(bytes)
	fmt.Println(string(bytes))
}

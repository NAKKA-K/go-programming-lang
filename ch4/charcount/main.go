package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Unicode char counts
	var utflen [utf8.UTFMax + 1]int // utf8 encoding len
	invalid := 0                    // invalid utf8 char num

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { // 不正文字の判定
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	// 表示処理
	fmt.Printf("rune]\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

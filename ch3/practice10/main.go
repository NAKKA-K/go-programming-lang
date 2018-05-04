package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n > 3 && n%3 != 0 {
		buf.WriteString(s[:n%3] + ",")
		s = s[n%3:]
	}
	for len(s) > 3 {
		buf.WriteString(s[:3] + ",")
		s = s[3:]
	}
	buf.WriteString(s)
	return buf.String()
}

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("12"))
}

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	floatIndex := strings.LastIndex(s, ".")
	var floatString string
	if floatIndex != -1 {
		floatString = s[floatIndex:]
		s = s[:floatIndex]
	}

	n := len(s)
	if n > 3 && n%3 != 0 {
		buf.WriteString(s[:n%3] + ",")
		s = s[n%3:]
	}
	for len(s) > 3 {
		buf.WriteString(s[:3] + ",")
		s = s[3:]
	}
	buf.WriteString(s + floatString)
	return buf.String()
}

func main() {
	fmt.Println(comma("12345.6789"))
	fmt.Println(comma("12345.678"))
	fmt.Println(comma("12"))
	fmt.Println(comma("12.1234"))
}

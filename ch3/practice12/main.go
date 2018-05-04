package main

import (
	"fmt"
	"strings"
)

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if strings.LastIndex(s1, string(s2[i])) == -1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("abc", "bac"))
	fmt.Println(isAnagram("ab", "bac"))
	fmt.Println(isAnagram("abc", "bc"))
	fmt.Println(isAnagram("", ""))
}

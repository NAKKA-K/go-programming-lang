package main

import "fmt"

func nonempty(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	fmt.Println(nonempty([]string{"", "aaa", "age"}))
}

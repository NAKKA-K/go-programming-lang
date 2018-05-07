package main

import "fmt"

func nonempty(strings []string) []string {
	var out []string
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	fmt.Println(nonempty([]string{"", "aaa", "age"}))
}

package main

import "fmt"

func removeOverlap(slice []string) []string {
	var out []string
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			continue
		}
		out = append(out, slice[i])
	}
	out = append(out, slice[len(slice)-1])
	return out
}

func main() {
	slice := []string{"aa", "aa", "aa", "aa", "aa", "bb", "cc", " ", "rr"}
	slice = removeOverlap(slice)
	for _, v := range slice {
		fmt.Printf("%v\n", v)
	}
}

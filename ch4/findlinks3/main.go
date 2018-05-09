package main

import (
	"fmt"
	"go-programming-lang/ch4/links"
	"log"
	"os"
	"strings"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	if strings.HasSuffix(url, "gz") || strings.HasSuffix(url, "msi") || strings.HasSuffix(url, "zip") || strings.HasSuffix(url, "pkg") {
		return nil
	}
	fmt.Println(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}

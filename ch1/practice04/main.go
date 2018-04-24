package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
        printCounts("Stdin", counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            printCounts(f.Name(), counts)
            f.Close()

            counts = make(map[string]int)
        }
    }

}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}

func printCounts(fileName string, counts map[string]int) {
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s: %d\t%s\n", fileName, n, line)
        }
    }
}
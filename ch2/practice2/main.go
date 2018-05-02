package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/tempconv"
)

func main() {
	nums := os.Args[1:]

	if len(nums) == 0 {
		fmt.Print("Input: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nums = strings.Split(scanner.Text(), " ")
	}

	for _, arg := range nums {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

package main

import (
	"fmt"
	"time"
)

func PopCount(x uint64) int {
	var bit int
	for bit = 0; x != 0; bit++ {
		x &= x - 1
	}
	return bit
}

func main() {
	start := time.Now()
	fmt.Println(PopCount(1000000))
	sec := time.Since(start).Seconds()
	fmt.Printf("%fs\n", sec) //=> 0.000028s
}

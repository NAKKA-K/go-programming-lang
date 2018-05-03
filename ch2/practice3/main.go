package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var bit int
	var i uint64
	for i = 0; i < 8; i++ {
		bit += int(pc[byte(x>>(i*8))])
	}
	return bit
}

func main() {
	start := time.Now()
	fmt.Println(PopCount(1000000))
	sec := time.Since(start).Seconds()
	fmt.Printf("%fs\n", sec) //=> 0.000028s

	startL := time.Now()
	fmt.Println(PopCountLoop(1000000))
	secL := time.Since(startL).Seconds()
	fmt.Printf("%fs\n", secL) //=> 0.000002s
}

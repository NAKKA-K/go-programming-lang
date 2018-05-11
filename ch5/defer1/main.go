package main

import (
	"fmt"
)

func main() {
	f(3)
	fmt.Println("end main")
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

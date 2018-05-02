package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BlilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%gC", c) }

func main() {
	c := FToC(212.0)
	fmt.Println(c.String()) //=> 100C
	fmt.Printf("%v\n", c)   //=> 100C
	fmt.Printf("%s\n", c)   //=> 100C
	fmt.Printf("%d\n", c)   //=> %!d(main.Celsius=100)
	fmt.Println(c)          //=> 100C
	fmt.Printf("%g\n", c)   //=> 100
	fmt.Println(float64(c)) //=> 100
}

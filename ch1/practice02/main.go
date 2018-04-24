// echo program: echoを修正して、個々の引数のインデックスと値の組を1行ごとに表示
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    s, sep := "", "\n"
    for i, arg := range os.Args {
        s += strconv.Itoa(i) + " " + arg + sep
    }
    fmt.Print(s)
}

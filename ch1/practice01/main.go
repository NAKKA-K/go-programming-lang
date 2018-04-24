// echo program: echoを修正して、起動したコマンド名も表示する
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args, " "))
}
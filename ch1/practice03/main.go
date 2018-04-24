package main

import (
	"os"
    "time"
    "fmt"
    "strings"
)

func main() {
    // 最初の処理速度は圧倒的に早いが、その後すごい勢いて遅くなる
    start := time.Now()
    var s, sep string
    for i := 1; i <= len(os.Args[1:]); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
    fmt.Printf("Old = %f\n", time.Since(start).Seconds())

    fmt.Println("===")

    // 一定な処理速度を保つため、負荷が低い間は少し遅い
    start = time.Now()
    fmt.Println(strings.Join(os.Args[1:], " "))    
    fmt.Printf("New = %f\n", time.Since(start).Seconds())
}
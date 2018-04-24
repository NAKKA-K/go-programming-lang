package main

import (
	"bufio"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    file, err := os.Open("to.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
    sc := bufio.NewScanner(file)

    start := time.Now()
    ch := make(chan string)
    cnt := 0
    for sc.Scan() {
        cnt++
        go fetch(sc.Text(), ch) //start goroutine
    }
    file.Close()

    for i := 0; i < cnt; i++ {
        fmt.Println(<-ch) // get from channel
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
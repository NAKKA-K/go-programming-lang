package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) //start goroutine
    }

    file, err := os.OpenFile("./tmp.out", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetchall: %s\n", err)
        os.Exit(1)
    }
    fmt.Fprintf(file, "\n%v\n", time.Now()) // print time stamp
    for range os.Args[1:] {
        //fmt.Println(<-ch)
        msg := <-ch // get from channel
        fmt.Printf("%s\n", msg)
        fmt.Fprintf(file, "%s\n", msg)
    }
    file.Close()
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
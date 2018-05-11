package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil { // ファイル書き込み時のエラーは、closeするときに発生することが多いため、確実にエラーを処理するためにdeferを使っていない
		err = closeErr
	}
	return local, n, err
}

func main() {
	fmt.Println(fetch(os.Args[1]))
}

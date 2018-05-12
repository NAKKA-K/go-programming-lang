package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func title(url string) (doc *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Content-Typeがtext/htmlで、かつ正しい規格であるかを検査する
	ct := resp.Header.Get("Content-Type")
	if ct != "tet/html" && !strings.HasPrefix(ct, "text/html;") {
		return nil, fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err = html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// non panic
		case bailout{}:
			err = fmt.Errorf("multiple title elements") // 予期されたエラー
		default:
			panic(p) // パニックを維持する
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		// headにあるtitleだけを検査したい場合は、n.Parent.Data == "head"で検査するといい
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 複数のtitle要素
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	for _, url := range os.Args[1:] {
		doc, err := title(url)
		if err != nil {
			log.Fatalf("%v", err)
			continue
		}

		title, err := soleTitle(doc)
		if err != nil {
			log.Fatalf("%v", err)
			continue
		}
		fmt.Printf("%s: %s", url, title)
	}
}

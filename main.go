package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	path := flag.String("p", "http://google.com", "path default to 'google.com' if not set")
	flag.Parse()

	pathString := string(*path)
	url := pathString
	if strings.HasPrefix(pathString, "/") {
		url = "http://google.com"+pathString
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error reading URL: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("error reading HTML: %v\n", err)
	}
	
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(doc)
}
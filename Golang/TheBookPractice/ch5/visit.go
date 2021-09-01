package ch5

import (
	"golang.org/x/net/html"
)

func visit(links []string, root *html.Node) []string {
	if root.Type == html.ElementNode && root.Data == "a" {
		for _, a := range root.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

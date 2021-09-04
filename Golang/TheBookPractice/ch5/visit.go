package ch5

import (
	"fmt"

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

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

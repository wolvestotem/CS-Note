package ch5

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func Input(s []string) {
	for _, url := range s {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks:%v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func Input2(s []string) {
	for _, url := range s {
		err := findLinks2(url)
		if err != nil {
			fmt.Printf("findlinks:%v\n", err)
		}
	}
}

// findLinks2 函数值实现
func findLinks2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

// Extract 匿名函数+函数值
func Extract(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	fmt.Println("No problem")
	var links []string
	visitNode := func(root *html.Node) {
		if root.Type == html.ElementNode && root.Data == "a" {
			for _, a := range root.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	for _, link := range links {
		fmt.Println(link)
	}
	return nil
}

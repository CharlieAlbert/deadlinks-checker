package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(htmlContent string) ([]string, error) {
	var links []string
	reader := strings.NewReader(htmlContent)

	doc, err := html.Parse(reader)

	if err != nil {
		return nil, err
	}

	var crawl func(*html.Node)
	crawl = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawl(c)
		}
	}

	crawl(doc)
	return links, nil
}
package main

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
)

func FetchAndExtract(ctx context.Context, pageURL string) (status, title string, links []string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", pageURL, nil)
	if err != nil {
		return "request_error", "", nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "fetch_error", "", nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("http_%d", resp.StatusCode), "", nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "parse_error", "", nil, err
	}
	title = extractTitle(doc)
	links = extractLinks(doc, pageURL)
	return "ok", title, links, nil
}

func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t := extractTitle(c)
		if t != "" {
			return t
		}
	}
	return ""
}

func extractLinks(n *html.Node, base string) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := strings.TrimSpace(attr.Val)
					if href == "" || strings.HasPrefix(href, "#") {
						continue
					}
					u, err := url.Parse(href)
					if err != nil {
						continue
					}
					baseURL, _ := url.Parse(base)
					link := baseURL.ResolveReference(u).String()
					links = append(links, link)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return links
}

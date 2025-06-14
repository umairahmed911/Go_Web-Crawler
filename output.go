package main

import (
	"fmt"
)

func PrintResult(res CrawlResult) {
	fmt.Printf("[Depth %d] %s | %s | %s\n", res.Depth, res.Status, res.URL, res.Title)
}

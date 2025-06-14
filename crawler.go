package main

import (
	"context"
	"sync"
)

type Crawler struct {
	rateLimiter *RateLimiter
	maxDepth   int
	visited    sync.Map // map[string]struct{}
	wg         sync.WaitGroup
}

type CrawlResult struct {
	URL    string
	Depth  int
	Status string
	Title  string
}

func NewCrawler(rateLimit, maxDepth int) *Crawler {
	return &Crawler{
		rateLimiter: NewRateLimiter(rateLimit),
		maxDepth:   maxDepth,
	}
}

func (c *Crawler) Crawl(ctx context.Context, url string, depth int, results chan<- CrawlResult) {
	if depth > c.maxDepth {
		return
	}
	if _, loaded := c.visited.LoadOrStore(url, struct{}{}); loaded {
		return
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.rateLimiter.Wait(ctx)
		status, title, links, err := FetchAndExtract(ctx, url)
		results <- CrawlResult{URL: url, Depth: depth, Status: status, Title: title}
		if err != nil || depth == c.maxDepth {
			return
		}
		for _, link := range links {
			c.Crawl(ctx, link, depth+1, results)
		}
	}()
}

func (c *Crawler) Wait() {
	c.wg.Wait()
}

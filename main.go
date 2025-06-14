package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"sync"
)

func main() {
	baseURL := flag.String("url", "", "Base URL to start crawling")
	maxDepth := flag.Int("depth", 2, "Maximum crawl depth")
	rateLimit := flag.Int("rate", 5, "Requests per second")
	xlsFile := flag.String("xls", "", "Export results to XLS file (optional)")
	flag.Parse()

	if *baseURL == "" {
		fmt.Print("Enter the URL to crawl: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("No URL entered. Exiting.")
			os.Exit(1)
		}
		*baseURL = input
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nReceived interrupt. Shutting down gracefully...")
		cancel()
	}()

	results := make(chan CrawlResult)
	var collected []CrawlResult

// Start goroutine to print results
var printWg sync.WaitGroup
printWg.Add(1)
go func() {
	defer printWg.Done()
	for res := range results {
		PrintResult(res)
		if *xlsFile != "" {
			collected = append(collected, res)
		}
	}
}()

crawler := NewCrawler(*rateLimit, *maxDepth)
crawler.Crawl(ctx, *baseURL, 0, results)

// Wait for crawling to finish
crawler.Wait()
close(results)
printWg.Wait()

	if *xlsFile != "" {
		xlsName := *xlsFile
		// Always enforce .xlsx extension
		if len(xlsName) > 4 && xlsName[len(xlsName)-4:] == ".xls" {
			xlsName = xlsName[:len(xlsName)-4] + ".xlsx"
		} else if len(xlsName) < 5 || xlsName[len(xlsName)-5:] != ".xlsx" {
			xlsName = xlsName + ".xlsx"
		}
		fmt.Printf("Exporting to file: %s\n", xlsName)

		err := ExportResultsToXLS(collected, xlsName)
		if err != nil {
			fmt.Printf("Failed to export XLSX: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", xlsName)
		}
	}
}

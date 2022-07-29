package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	url := bufio.NewScanner(os.Stdin)
	url.Scan()
	u := url.Text()

	hostName, err := extractHostnameFromURL(u)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return
	}

	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(".*(\\.|\\/\\/)" + strings.ReplaceAll(hostName, ".", "\\.") + "((#|\\/|\\?).*)?"),
		),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// c.OnHTML("script[src]", func(e *colly.HTMLElement) {
	// 	printResults(e.Attr("src"), "script", results, e)
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(u)
}

func extractHostnameFromURL(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}

func printResults(link string, sourceName string, results chan string, e *colly.HTMLElement) {
	result := e.Request.AbsoluteURL(link)

	if result != "" {
		result = "[" + sourceName + "] " + result
	}

	results <- result
}

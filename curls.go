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

	allowedDomains := []string{hostName}

	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomains...),
		colly.URLFilters(
			regexp.MustCompile(".*(\\.|\\/\\/)"+strings.ReplaceAll(hostName, ".", "\\.")+"((#|\\/|\\?).*)?"),
		),
		colly.MaxDepth(2),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		fmt.Printf(e.Attr("src"), "script", e)
	})

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

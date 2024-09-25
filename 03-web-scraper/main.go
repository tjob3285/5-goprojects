package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var startURL = "https://scrape-me.dreamsofcode.io"
var checkedURLs = []string{}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go checkAndExtractLinks(startURL, &wg)
	wg.Wait()
}

// extractLinks parses the HTML and extracts all links.
func extractLinks(resp *http.Response) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

// isExcludedDomain checks if the URL is from a specified domain to exclude.
func isExcludedDomain(link string) bool {
	excludedDomains := []string{
		"youtube.com",
		"youtu.be",
		"facebook.com",
		"instagram.com",
		"twitter.com",
	}

	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}

	for _, domain := range excludedDomains {
		if strings.Contains(parsedURL.Hostname(), domain) {
			return true
		}
	}
	return false
}

func checkAndExtractLinks(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	fmt.Printf("URL: %s, Status Code: %d\n", link, resp.StatusCode)
	checkedURLs = append(checkedURLs, link)
	defer resp.Body.Close()

	subLinks, err := extractLinks(resp)
	if err != nil {
		fmt.Printf("Error extracting links: %v\n", err)
		return
	}

	if len(subLinks) == 0 {
		return
	}

	var subWg sync.WaitGroup
	subWg.Wait()
	for i := 0; i < len(subLinks); i++ {
		if isExcludedDomain(subLinks[i]) {
			subLinks[i] = subLinks[len(subLinks)-1]
		}

		if contains(checkedURLs, startURL+subLinks[i]) {
			subLinks[i] = subLinks[len(subLinks)-1]
		} else {
			subWg.Add(1)
			go checkAndExtractLinks(startURL+subLinks[i], &subWg)
		}
	}
	subWg.Wait()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

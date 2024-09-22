package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

/*
// checkURL checks if the URL is valid and returns its status code.
func checkURL(link string) (int, error) {
	resp, err := http.Get(link)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
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

// checkAndExtractLinks checks a URL and extracts its sublinks.
func checkAndExtractLinks(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	if isExcludedDomain(link) {
		fmt.Printf("Skipping excluded domain: %s\n", link)
		return
	}

	statusCode, err := checkURL(link)
	if err != nil {
		fmt.Printf("Error checking URL: %v\n", err)
		return
	}

	fmt.Printf("URL: %s, Status Code: %d\n", link, statusCode)

	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	subLinks, err := extractLinks(resp)
	if err != nil {
		fmt.Printf("Error extracting links: %v\n", err)
		return
	}

	var subWg sync.WaitGroup
	for _, subLink := range subLinks {
		subWg.Add(1)
		go checkAndExtractLinks(link+subLink, &subWg)
	}
	subWg.Wait()
}

// main function to run the scraper.
func main() {
	/*if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	startURL := os.Args[1]
	startURL := "https://scrape-me.dreamsofcode.io"
	var wg sync.WaitGroup
	wg.Add(1)
	go checkAndExtractLinks(startURL, &wg)
	wg.Wait()
}*/

var startURL = "https://scrape-me.dreamsofcode.io"

func main() {
	checkAndExtractLinks(startURL)
}

// checkURL checks if the URL is valid and returns its status code.
func checkURL(link string) (int, error) {
	resp, err := http.Get(link)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
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

func checkAndExtractLinks(link string) {
	statusCode, err := checkURL(link)
	if err != nil {
		fmt.Printf("Error checking URL: %v\n", err)
		return
	}

	fmt.Printf("URL: %s, Status Code: %d\n", link, statusCode)

	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	subLinks, err := extractLinks(resp)
	if err != nil {
		fmt.Printf("Error extracting links: %v\n", err)
		return
	}

	fmt.Printf("SubLinks: %s\n", subLinks)
	if len(subLinks) == 0 {
		return
	}

	for i := 0; i < len(subLinks); i++ {
		if isExcludedDomain(subLinks[i]) {
			fmt.Printf("Skipping excluded domain: %s\n", subLinks[i])
			subLinks[i] = subLinks[len(subLinks)-1]
		}

		checkAndExtractLinks(startURL + subLinks[i])
	}
}

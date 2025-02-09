package pratices

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// CrawlQueue to manage visited URLs
type CrawlQueue struct {
	visited map[string]bool
	mu      sync.Mutex
}

// NewCrawlQueue initializes a new crawl queue
func NewCrawlQueue() *CrawlQueue {
	return &CrawlQueue{visited: make(map[string]bool)}
}

// Add marks a URL as visited and returns whether it was new
func (c *CrawlQueue) Add(u string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.visited[u] {
		return false
	}
	c.visited[u] = true
	return true
}

// Fetch and parse links from a given URL
func fetchLinks(baseURL string) ([]string, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	links := []string{}
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if strings.HasPrefix(link, "/") {
							parsedURL, _ := url.Parse(baseURL)
							link = parsedURL.Scheme + "://" + parsedURL.Host + link
						}
						links = append(links, link)
					}
				}
			}
		}
	}
}

// Crawl function
func crawl(startURL string, depth int, queue *CrawlQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 || !queue.Add(startURL) {
		return
	}

	fmt.Println("Crawling:", startURL)
	links, err := fetchLinks(startURL)
	if err != nil {
		fmt.Println("Error fetching:", err)
		return
	}

	for _, link := range links {
		wg.Add(1)
		go crawl(link, depth-1, queue, wg)
	}
}

// Main function
func RunWebCrawler() {
	startURL := "https://example.com"
	depth := 2
	queue := NewCrawlQueue()
	var wg sync.WaitGroup

	wg.Add(1)
	go crawl(startURL, depth, queue, &wg)
	wg.Wait()
}

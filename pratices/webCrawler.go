package pratices

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// crawlQueue to manage visited URLs
type crawlQueue struct {
	visited map[string]bool // Maps a URL (string) to a boolean (true if visited)
	mu      sync.Mutex      // Mutex to prevent race conditions in concurrent access
}

// newCrawlQueue initializes a new crawl queue
func newCrawlQueue() *crawlQueue {
	return &crawlQueue{visited: make(map[string]bool)}
}

// Add marks a URL as visited and returns whether it was new
func (c *crawlQueue) Add(u string) bool {
	c.mu.Lock()         // Lock the mutex to ensure exclusive access to `visited`
	defer c.mu.Unlock() // Unlock at the end of function execution

	if c.visited[u] { // Check if the URL is already visited
		return false // If yes, return false (URL is already processed)
	}

	c.visited[u] = true // Mark URL as visited
	return true         // Return true (URL was newly added)
}

// Fetch and parse links from a given URL
// fetchLinks fetches and parses the links (href attributes) from the given URL (baseURL)
func fetchLinks(baseURL string) ([]string, error) {
	// Send an HTTP GET request to the baseURL and store the response in resp.
	// If there's an error (e.g., the URL is unreachable), return the error.
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}

	// Defer the closing of the response body (to ensure it gets closed when the function returns).
	// The close function is wrapped in an anonymous function so that we can handle any potential error from closing.
	defer func(Body io.ReadCloser) {
		// Attempt to close the response body.
		err := Body.Close()
		// If there was an error closing the body, log it and terminate the program.
		if err != nil {
			log.Fatal("Error on Body.Close()")
		}
	}(resp.Body)

	// Initialize an empty slice to store the links we will find.
	links := []string{}

	// Create a new HTML tokenizer that will allow us to parse the HTML from the response body.
	tokenizer := html.NewTokenizer(resp.Body)

	// Start a loop that processes the tokens in the HTML document one by one.
	for {
		// Get the next token from the tokenizer.
		tokenType := tokenizer.Next()

		// Switch on the type of token we encountered.
		switch tokenType {
		// If we encounter an error token or reach the end of the document, return the links found so far.
		case html.ErrorToken:
			return links, nil
		// If we encounter a start tag (e.g., <a>), process it.
		case html.StartTagToken:
			token := tokenizer.Token() // Retrieve the full token for the start tag.

			// If the tag is an anchor tag (<a>), check its attributes.
			if token.Data == "a" {
				// Loop through the attributes of the anchor tag.
				for _, attr := range token.Attr {
					// If the attribute is "href" (i.e., the URL), capture its value.
					if attr.Key == "href" {
						link := attr.Val // The value of the href attribute is the link.

						// If the link is a relative URL (starts with "/"), convert it to an absolute URL.
						if strings.HasPrefix(link, "/") {
							// Parse the baseURL into a URL object.
							parsedURL, _ := url.Parse(baseURL)
							// Combine the base URL's scheme and host with the relative link to form a full URL.
							link = parsedURL.Scheme + "://" + parsedURL.Host + link
						}

						// Append the found link (absolute or relative) to the links slice.
						links = append(links, link)
					}
				}
			}
		}
	}
}

// Crawl function
func crawl(startURL string, depth int, queue *crawlQueue, wg *sync.WaitGroup) {
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
// RunWebCrawler starts the crawling process for a given URL
func RunWebCrawler() {
	// The starting URL to begin the crawl from
	startURL := "https://www.example.ai"

	// The depth of the crawl: how many levels deep we want to crawl
	// Depth 2 means it will crawl the startURL and links found on that page (i.e., first 2 levels)
	depth := 3

	// Initialize a new crawlQueue to track visited URLs
	// This queue will keep track of which URLs have already been visited to prevent revisiting them
	queue := newCrawlQueue()

	// Create a WaitGroup to manage goroutines (for concurrent crawling)
	var wg sync.WaitGroup

	// Add 1 to the WaitGroup counter to indicate we're starting 1 goroutine (the initial crawl)
	wg.Add(1)

	// Start the crawl process in a new goroutine.
	// `go crawl(startURL, depth, queue, &wg)` initiates the crawl concurrently,
	// which means the crawl won't block the main thread, and other tasks can run in parallel if needed.
	go crawl(startURL, depth, queue, &wg)

	// Wait for all goroutines (in this case, the initial crawl goroutine) to complete
	// This ensures that the program won't exit until the crawling process finishes
	wg.Wait()
}

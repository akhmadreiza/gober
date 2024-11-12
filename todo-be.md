# Gober BE step-by-step

Building a concurrent web scraper using Go is a great way to practice goroutines and channels. Here’s a step-by-step guide to help you get started:

### Step-by-Step Guide to Implement a Concurrent Web Scraper in Go

#### 1. **Set up your project**
   - Create a new Go project.
   - Initialize your Go module if needed using `go mod init concurrent-web-scraper`.

```bash
go mod init concurrent-web-scraper
```

#### 2. **Install necessary libraries**
   Go has a built-in `net/http` package to make HTTP requests. You may also need libraries like `goquery` for parsing HTML content, which you can install as follows:

```bash
go get -u github.com/PuerkitoBio/goquery
```

#### 3. **Basic HTTP request in Go**
   Start by writing a function that fetches a webpage's content using `net/http`. This will give you the raw HTML response to work with.

```go
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	url := "https://example.com"
	body, err := fetchURL(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	fmt.Println("Fetched content:", body[:100]) // Print first 100 characters of body
}
```

#### 4. **Use `goquery` to parse the HTML**
   Once you have the page content, use `goquery` to extract specific data (e.g., titles, links, or other information).

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func fetchTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Extract the page title
	title := doc.Find("title").Text()
	return title, nil
}

func main() {
	url := "https://example.com"
	title, err := fetchTitle(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Page title:", title)
}
```

#### 5. **Implement concurrency using Goroutines and Channels**
   Now that you have the scraper working for a single page, use goroutines and channels to scrape multiple websites concurrently.

```go
package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Fetch the page title for a given URL
func fetchTitle(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching URL %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("Error: status code %d on %s", resp.StatusCode, url)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error parsing URL %s: %v", url, err)
		return
	}

	title := doc.Find("title").Text()
	ch <- fmt.Sprintf("Title of %s: %s", url, title)
}

func main() {
	// List of URLs to scrape
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
	}

	// Create a channel to receive titles
	ch := make(chan string)

	// Use a WaitGroup to ensure all goroutines complete
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchTitle(url, ch, &wg)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Print out results from channel
	for result := range ch {
		fmt.Println(result)
	}
}
```

**Explanation**:
- We create a `fetchTitle` function that retrieves the page title from a given URL.
- The main function spins up a goroutine for each URL in the list and uses a channel to collect the titles as they are fetched.
- A `sync.WaitGroup` is used to ensure that the main goroutine waits until all scraping goroutines finish before closing the channel.
  
#### 6. **Handle errors and edge cases**
   Make sure you handle common errors such as:
   - Invalid URLs.
   - Timeout issues (you can use `http.Client` with timeouts instead of `http.Get`).
   - Handling non-200 status codes properly.
   
```go
client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Get(url)
```

#### 7. **Optimizations: Use Buffered Channels**
   If you have many URLs to scrape, using buffered channels can improve performance by allowing the program to process results without waiting for all goroutines to finish.

```go
ch := make(chan string, len(urls)) // Buffered channel
```

#### 8. **Further improvements**
   - **Rate limiting**: Prevent overwhelming a server by adding delays or restricting the number of concurrent requests.
   - **Retry logic**: Add retries for failed requests due to temporary network issues.
   - **Scrape other elements**: Instead of titles, you can scrape other page elements like links, headings, or specific content.
   
### Conclusion
This step-by-step guide shows how you can build a basic concurrent web scraper using Go. You can extend this project further to scrape multiple types of data, implement rate-limiting, or scrape pages from paginated results. 

Let me know if you need help with any of these additional features!
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	visited map[string]bool
	mux     sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher, notify chan<- int) {
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer func() { notify <- 0 }()
	c.markVisited(url)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	notifies := make([]chan int, 0)
	for _, u := range urls {
		if c.isVisited(u) {
			continue
		}
		n := make(chan int)
		go c.Crawl(u, depth-1, fetcher, n)
		notifies = append(notifies, n)
	}
	for _, n := range notifies {
		<-n
	}
}

func (c *Crawler) isVisited(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.visited[url]
}

func (c *Crawler) markVisited(url string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.visited[url] = true
}

func main() {
	notify := make(chan int)
	c := Crawler{visited: make(map[string]bool)}
	go c.Crawl("http://golang.org/", 4, fetcher, notify)
	<-notify
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type BookInfo struct {
	title: string
	author: string
	five: bool
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("article.product_pod star-rating Five e.children", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Other selectors to try:
		// article.product_pod star-rating Five
		// e.children
		// Previously used selector: .Five , a

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://books.toscrape.com/")
}

package scrapeBooks

import (
	"github.com/gocolly/colly"
)

type BookData struct {
	title string: ""
	author string: ""
	price float32: 0
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: http://books.toscrape.com
		colly.AllowedDomains("hhttp://books.toscrape.com"),
		colly.MaxDepth(2)
	)
}

// resource: http://go-colly.org/docs/examples/basic/
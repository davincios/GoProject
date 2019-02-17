package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func scrape(website string) {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains(website))

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	// Correct URL from input.
	s := fmt.Sprintf("https://%s/sitemap.xml", website)

	// Start the collector
	c.Visit(s)

	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println("\t", url)
	}
	fmt.Println("Collected", len(knownUrls), "URLs")

}

func main() {
	scrape("www.flymble.com")
}

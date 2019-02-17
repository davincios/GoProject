package main

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

type PageVariables struct {
	Date string
	Time string
}

func scrape(website string, w http.ResponseWriter) {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}

	// Create a Collector and specify allowed domains.
	c := colly.NewCollector(colly.AllowedDomains(website))

	// specify a dictionary to store the links in.
	//p := &pageInfo{Links: make(map[string]int)}

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	// Correct URL format.
	s := fmt.Sprintf("https://%s/sitemap.xml", website)

	// Start the collector
	c.Visit(s)

	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println("\t", url)
		// This code allows for the line break to be added
		fmt.Fprintln(w, "\t", url)

	}
	fmt.Println("Collected", len(knownUrls), "URLs")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	scrape("www.flymble.com", w)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

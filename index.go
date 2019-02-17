package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

type deal struct {
	DealURL string
	Title   string
}

func scrape(website string, w http.ResponseWriter) {
	// Array containing all the known URLs in a sitemap
	deals := []deal{}

	// Create a Collector and specify allowed domains.
	c := colly.NewCollector(colly.AllowedDomains(website))

	// Get title of each deal
	c.OnHTML("header h2", func(e *colly.HTMLElement) {
		line := strings.Join(strings.Split(e.ChildText("a"), "Show deal:")[1:], "")

		temp := deal{}
		temp.Title = line
		deals = append(deals, temp)
		// knownUrls = append(knownUrls, line)

	})

	// Get the link to each deal
	c.OnHTML("header h2 a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		temp := deal{}
		temp.DealURL = url
		deals = append(deals, temp)

	})

	// Correct URL format.
	s := fmt.Sprintf("https://%s", website)

	// Start the collector
	c.Visit(s)

	// Print the found URLS
	fmt.Println("All known Deals:")
	for _, deal := range deals {
		// Create correct string with line break
		// This code allows for the line break to be added
		fmt.Fprintln(w, "\t", deal.Title)
		fmt.Fprintln(w, "\t", deal.DealURL)
		fmt.Fprintln(w, "\t", "------------")

	}
	fmt.Println("Collected", len(deals), "Elements")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	scrape("www.holidaypirates.com", w)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

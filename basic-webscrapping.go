package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com", "wiki.hackerspaces.org"),
	)

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("span.text"))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visisted", r.Request.URL)
	})

	c.Visit("https://quotes.toscrape.com/")
}

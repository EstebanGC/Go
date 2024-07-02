package main 

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com", "www.toquotes.toscrape.com")
	)
	

	c.OnHTML("[itemprop='headline']", func(e *colly.HTMLElement) {
		titulo := e.Text
		fmt.Println("Titulo: ", titulo)
	})

	c.OnHTML(".story-body__inner", func(e *colly.HTMLElement) {
		parrafo := e.Text
		fmt.Println("Primer Párrafo: ", parrafo)
	})

	c.Visit("https://quotes.toscrape.com/")
}
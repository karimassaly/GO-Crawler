package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("twitter.com"),
	)

	c.OnHTML("div", func(e *colly.HTMLElement) {
		tweet := e.ChildAttr("article", "role")

		fmt.Printf("tweet found: %q -> %s\n", e.Text, tweet)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://twitter.com/realdonaldtrump")
}

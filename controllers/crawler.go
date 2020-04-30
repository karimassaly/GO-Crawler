package controllers

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

var i int = 0

func Crawler(twitter string) {
	c := colly.NewCollector(
		colly.AllowedDomains("twitter.com"),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.AbsoluteURL(e.Attr("href"))
		i++
		fmt.Println(i)
		ParserTweets(e.Text, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		PageWriter(true, "Request failed with responde: User does not exist")
	})

	Url := "https://twitter.com/" + twitter
	c.Visit(Url)
}

func ParserTweets(text, link string) {

	if i < 50 || i > 178 {
		return
	}
	if text == "" {
		return
	}
	if string(text[0]) >= "0" && string(text[0]) <= "9" {
		fmt.Printf("tweet found: %q -> %s \n", text, link)
	}

}

package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func main() {
	// create new collector from colly
	collector := colly.NewCollector()

	// find all link in the page
	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		srcUrl := element.Request.URL

		// if link is relative, resolve it
		if !strings.HasPrefix(link, "http") {
			link = srcUrl.String() + link
		}
		fmt.Println(link)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	err := collector.Visit("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done")
}

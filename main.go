package main

import (
	"UrlExtractor/cmd"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

func main() {
	// parse command line arguments
	executeParam, err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("The URL to be crawled is : " + executeParam.Url)

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

	err = collector.Visit(executeParam.Url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done")
}

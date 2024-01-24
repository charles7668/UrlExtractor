package main

import (
	"UrlExtractor/cmd"
	"UrlExtractor/fileutil"
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
	if len(executeParam.FileToWrite) > 0 {
		if !fileutil.CheckPathValid(executeParam.FileToWrite) {
			fmt.Println("File path is invalid")
			os.Exit(1)
		}
		fmt.Println("Write output to file : " + executeParam.FileToWrite)
	}

	var links []string
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
		links = append(links, link)
		fmt.Println(link)
	})

	// visit the url
	err = collector.Visit(executeParam.Url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write all links to file
	if len(executeParam.FileToWrite) > 0 {
		file, err := os.Create(executeParam.FileToWrite)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(file)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, link := range links {
				_, err := fmt.Fprintln(file, link)
				if err != nil {
					break
				}
			}
		}
	}

	fmt.Println("Done")
}

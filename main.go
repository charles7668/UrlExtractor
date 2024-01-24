package main

import (
	"UrlExtractor/cmd"
	"UrlExtractor/container"
	"UrlExtractor/fileutil"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
	"time"
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
	fmt.Println("Depth of crawling : ", executeParam.Depth)

	var record = make(map[string]bool)
	var queue container.Queue
	queue.Enqueue(executeParam.Url)
	// create new collector from colly
	collector := colly.NewCollector()

	for i := 0; i < executeParam.Depth; i++ {
		queueSize := queue.Size()
		for j := 0; j < queueSize; j++ {
			visitUrl := queue.Dequeue().(string)
			// find all link in the page
			collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
				link := element.Attr("href")
				srcUrl := element.Request.URL

				// if link is relative, resolve it
				if !strings.HasPrefix(link, "http") {
					link = srcUrl.String() + link
				}
				if record[link] {
					return
				}
				record[link] = true
				queue.Enqueue(link)
				fmt.Println(link)
			})

			// visit the url
			err = collector.Visit(visitUrl)
			if err != nil {
				fmt.Println("url : ", visitUrl, "error : ", err)
				continue
			}

			// wait for 50ms to avoid being blocked
			time.Sleep(50)
		}
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
			for link := range record {
				_, err := fmt.Fprintln(file, link)
				if err != nil {
					break
				}
			}
		}
	}

	fmt.Println("Done")
}

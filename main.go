// scraper.go project main.go
package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	fmt.Printf("Fetching contents\n")
	doc, err := goquery.NewDocument("http://iknowthepilot.com.au")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.blugfull div.post > h1.title a[rel=bookmark]").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		href, _ := item.Attr("href")
		fmt.Printf("Post #%d: %s %s\n", index, title, href)
	})

}

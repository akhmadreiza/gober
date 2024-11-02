package main

import (
	"fmt"

	"github.com/akhmadreiza/gober/parsers"
	"github.com/akhmadreiza/gober/scraper"
)

func main() {
	scraper, err := getScraper("detik")
	if err != nil {
		fmt.Println("Error when getting scraper:", err)
		return
	}

	articles, err := scraper.Search("prabowo")
	if err != nil {
		fmt.Println("Error scrap URL:", err)
		return
	}

	for _, article := range articles {
		fmt.Printf("Scraped Article: URL=%s, Title=%s\n", article.URL, article.Title)
	}

	popArticles, err := scraper.Popular()
	if err != nil {
		fmt.Println("Error getting popular news:", err)
		return
	}
	for _, popArticle := range popArticles {
		fmt.Printf("Popular Article: URL=%s, Title=%s\n", popArticle.URL, popArticle.Title)
	}

}

func getScraper(website string) (scraper.NewsScraper, error) {
	if website == "detik" {
		return parsers.DetikScraper{}, nil
	}
	return nil, fmt.Errorf("scrape %v not supported", website)
}

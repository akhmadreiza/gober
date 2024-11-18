package utils

import (
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/akhmadreiza/gober/models"
	"github.com/gin-gonic/gin"
)

func FetchListArticles(f func(doc *goquery.Document, c *gin.Context) []models.Article, urls []string, c *gin.Context) []models.Article {
	// Create a channel to receive Articles
	ch := make(chan []models.Article)

	// Use a WaitGroup to ensure all goroutines complete
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchListArticlesRoutine(url, ch, &wg, c, f)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	var listArticles []models.Article
	for result := range ch {
		listArticles = append(listArticles, result...)
	}
	return listArticles
}

func fetchListArticlesRoutine(url string, ch chan []models.Article, waitGroup *sync.WaitGroup, ginContext *gin.Context, f func(doc *goquery.Document, c *gin.Context) []models.Article) {
	//call waitGroup.Done at the end of method
	defer waitGroup.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- []models.Article{}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- []models.Article{}
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		ch <- []models.Article{}
		return
	}

	ch <- f(doc, ginContext)
}

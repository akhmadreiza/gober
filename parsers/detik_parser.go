package parsers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/akhmadreiza/gober/models"
	"github.com/gin-gonic/gin"
)

type DetikScraper struct{}

func (detik DetikScraper) Detail(url string) (models.Article, error) {
	log.Println("accessing", url)
	resp, err := http.Get(url)
	if err != nil {
		return models.Article{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.Article{}, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return models.Article{}, err
	}

	title := doc.Find("h1.detail__title").Text()
	author := doc.Find("div.detail__author").Text()
	articleDate := doc.Find("div.detail__date").Text()

	article := models.Article{}
	article.URL = url
	article.Title = title
	article.Author = author
	article.Date = articleDate

	var content string
	doc.Find("div.detail__body-text.itp_bodycontent").Children().Each(func(i int, s *goquery.Selection) {
		if s.Is("p") || s.Is("h3") {
			content = content + s.Text() + "\n"
		}
		// if s.Find("p") != nil {
		// 	content = content + s.Find("p").Text() + "\\n"
		// }
		// if s.Find("h3") != nil {
		// 	content = content + s.Find("h3").Text() + "\\n"
		// }
	})
	article.Content = content

	return article, nil
}

func (detik DetikScraper) Search(keyword string, ginContext *gin.Context) ([]models.Article, error) {
	searchUrl := fmt.Sprintf("https://www.detik.com/search/searchall?query=%v&page=1&result_type=latest", keyword)
	log.Println("accessing", searchUrl)
	resp, err := http.Get(searchUrl)
	if err != nil {
		return []models.Article{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []models.Article{}, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []models.Article{}, err
	}

	return fetchListArticles(doc, ginContext), nil
}

func (detik DetikScraper) Popular(ginContext *gin.Context) ([]models.Article, error) {
	popUrls := []string{
		"https://www.detik.com/terpopuler/news",
		"https://www.detik.com/terpopuler/finance",
		"https://www.detik.com/terpopuler/hot",
		"https://www.detik.com/terpopuler/inet",
		"https://www.detik.com/terpopuler/sport",
		"https://www.detik.com/terpopuler/oto",
		"https://www.detik.com/terpopuler/travel",
		"https://www.detik.com/terpopuler/sepakbola",
		"https://www.detik.com/terpopuler/food",
		"https://www.detik.com/terpopuler/health",
		"https://www.detik.com/terpopuler/wolipop",
		"https://www.detik.com/terpopuler/edu",
	}

	// Create a channel to receive Articles
	ch := make(chan []models.Article)

	// Use a WaitGroup to ensure all goroutines complete
	var wg sync.WaitGroup

	for _, url := range popUrls {
		wg.Add(1)
		go fetchListArticlesRoutine(url, ch, &wg, ginContext)
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
	return listArticles, nil
}

func fetchListArticlesRoutine(url string, ch chan []models.Article, waitGroup *sync.WaitGroup, ginContext *gin.Context) {
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

	ch <- fetchListArticles(doc, ginContext)
}

func fetchListArticles(doc *goquery.Document, c *gin.Context) []models.Article {
	var listArticles []models.Article
	doc.Find("article.list-content__item").Each(func(i int, s *goquery.Selection) {
		article := models.Article{}
		media := s.Find("h3.media__title")

		var resultUrl string
		articleUrl, attrExists := media.Find("a").Attr("href")
		if attrExists {
			resultUrl = articleUrl
		}

		articleTitle := media.Find("a").Text()

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}
		article.URL = scheme + "://" + c.Request.Host + "/article?detailUrl=" + url.QueryEscape(resultUrl)
		article.SourceUrl = resultUrl
		article.Title = articleTitle

		listArticles = append(listArticles, article)
	})
	return listArticles
}

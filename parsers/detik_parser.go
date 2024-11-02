package parsers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/akhmadreiza/gober/models"
)

type DetikScraper struct{}

func (detik DetikScraper) Search(keyword string) ([]models.Article, error) {
	searchUrl := fmt.Sprintf("https://www.detik.com/search/searchall?query=%v&page=1&result_type=latest", keyword)
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

	return fetchListArticles(doc), nil
}

func (detik DetikScraper) Popular() ([]models.Article, error) {
	popularUrl := "https://www.detik.com/terpopuler/news"
	resp, err := http.Get(popularUrl)
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

	return fetchListArticles(doc), nil
}

func fetchListArticles(doc *goquery.Document) []models.Article {
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

		article.URL = resultUrl
		article.Title = articleTitle

		listArticles = append(listArticles, article)
	})
	return listArticles
}

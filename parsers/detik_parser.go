package parsers

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/akhmadreiza/gober/models"
	"github.com/akhmadreiza/gober/utils"
	"github.com/gin-gonic/gin"
)

type DetikScraper struct {
	Client utils.HTTPClient
	Utils  utils.ScrapeUtils
	Cache  utils.CacheOps
}

func (detik DetikScraper) Detail(detailUrl string, c *gin.Context) (models.Article, error) {
	if cachedData, found := detik.Cache.Get("detik:" + detailUrl); found {
		return cachedData.(models.Article), nil
	}

	resp, err := detik.Client.Get(detailUrl)
	if err != nil {
		return models.Article{}, err
	}

	if resp.Status != 200 {
		return models.Article{}, fmt.Errorf("error: status code %d", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Body))
	if err != nil {
		return models.Article{}, err
	}

	title := doc.Find("h1.detail__title").Text()
	author := doc.Find("div.detail__author").Text()
	articleDate := doc.Find("div.detail__date").Text()
	imageUrl, _ := doc.Find("div.detail__media").Find("img").Attr("src")

	article := models.Article{}
	article.URL = detailUrl
	article.Title = title
	article.Author = author
	article.Date = articleDate
	article.ImgUrl = imageUrl

	html, err := doc.Find("div.detail__body-text.itp_bodycontent").Html()

	if err != nil {
		return models.Article{}, err
	}

	article.Content = html

	if html == "" {
		html, err := doc.Find("div.detail__body-text").Html()

		if err != nil {
			return models.Article{}, err
		}

		article.Content = html
	}

	detik.Cache.Set("detik:"+article.URL, article, 5*time.Minute)
	return article, nil
}

func (detik DetikScraper) Search(keyword string, ginContext *gin.Context) ([]models.Article, error) {
	searchUrl := fmt.Sprintf("https://www.detik.com/search/searchall?query=%v&page=1&result_type=latest", keyword)
	resp, err := detik.Client.Get(searchUrl)
	if err != nil {
		return []models.Article{}, err
	}

	if resp.Status != 200 {
		return []models.Article{}, fmt.Errorf("error: status code %d", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Body))
	if err != nil {
		return []models.Article{}, err
	}

	return fetchArticlesDetik(doc, ginContext), nil
}

func (detik DetikScraper) Popular(ginContext *gin.Context) ([]models.Article, error) {
	if cachedData, found := detik.Cache.Get("detik:popular"); found {
		log.Print("cache detik:popular found. return data from cache.")
		return cachedData.([]models.Article), nil
	}

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
		"https://www.detik.com/terpopuler/edu",
	}

	result := detik.Utils.FetchListArticles(fetchArticlesDetik, popUrls, ginContext)

	log.Printf("Detik articles: %v", len(result))
	if len(result) > 0 {
		detik.Cache.Set("detik:popular", result, 5*time.Minute)
	}

	return result, nil
}

func fetchArticlesDetik(doc *goquery.Document, c *gin.Context) []models.Article {
	var listArticles []models.Article
	doc.Find("article.list-content__item").Each(func(i int, s *goquery.Selection) {
		article := models.Article{}
		media := s.Find("h3.media__title")
		img, imgExists := s.Find("div.media__image").Find("img").Attr("src")

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
		article.URL = scheme + "://" + c.Request.Host + "/article?detailUrl=" + url.QueryEscape(resultUrl+"?single=1")
		article.SourceUrl = resultUrl + "?single=1"
		article.Title = articleTitle
		if imgExists {
			article.ImgUrl = img
		}

		parsedUrl, err := url.Parse(resultUrl)
		if err == nil {
			article.ShortDesc = parsedUrl.Host
		}

		article.Date = s.Find("div.media__date").Find("span").AttrOr("title", "-")

		listArticles = append(listArticles, article)
	})
	return listArticles
}

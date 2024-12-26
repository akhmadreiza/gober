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

type KompasScraper struct {
	Client utils.HTTPClient
	Utils  utils.ScrapeUtils
	Cache  utils.CacheOps
}

func (k KompasScraper) Search(keyword string, g *gin.Context) ([]models.Article, error) {
	return []models.Article{}, fmt.Errorf("KompasScraper Search is not supported")
}

func (k KompasScraper) Popular(c *gin.Context) ([]models.Article, error) {
	if cachedData, found := k.Cache.Get("kompas:popular"); found {
		log.Print("cache kompas:popular found. return data from cache.")
		return cachedData.([]models.Article), nil
	}

	popUrls := []string{
		"https://indeks.kompas.com/terpopuler",
		"https://indeks.kompas.com/headline",
	}

	result := k.Utils.FetchListArticles(fetchArticlesKompas, popUrls, c)

	log.Printf("Kompas articles: %v", len(result))
	if len(result) > 0 {
		k.Cache.Set("kompas:popular", result, 5*time.Minute)
	}

	return result, nil
}

func (k KompasScraper) Detail(url string, c *gin.Context) (models.Article, error) {
	if cachedData, found := k.Cache.Get("kompas:" + url); found {
		return cachedData.(models.Article), nil
	}

	resp, err := k.Client.Get(url)
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

	title := doc.Find("h1.read__title").Text()
	author := doc.Find("div.credit-title").Text()
	articleDate := doc.Find("div.read__time").Text()
	imageUrl, _ := doc.Find("div.photo__wrap").Find("img").Attr("src")

	article := models.Article{}
	article.URL = url
	article.Title = title
	article.Author = author
	article.Date = articleDate
	article.ImgUrl = imageUrl

	html, _ := doc.Find("div.read__content").Html()
	article.Content = html

	k.Cache.Set("kompas:"+article.URL, article, 5*time.Minute)
	return article, nil
}

func fetchArticlesKompas(doc *goquery.Document, c *gin.Context) []models.Article {
	var listArticles []models.Article
	doc.Find("div.articleItem").Each(func(i int, s *goquery.Selection) {
		article := models.Article{}

		var resultUrl string
		articleUrl, attrExists := s.Find("a.article-link").Attr("href")
		if attrExists {
			resultUrl = articleUrl
		}
		article.SourceUrl = resultUrl + "?page=all"

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}
		article.URL = scheme + "://" + c.Request.Host + "/article?detailUrl=" + url.QueryEscape(resultUrl)

		parsedUrl, err := url.Parse(resultUrl + "?page=all")
		if err == nil {
			article.ShortDesc = parsedUrl.Host

			//extract date from kompas url
			su := strings.Split(parsedUrl.Path, "/")
			at := su[len(su)-2]
			ah := at[0:2]
			am := at[2:4]
			article.Date = s.Find("div.articlePost-date").Text() + ", " + ah + ":" + am + " WIB"
		}

		img, imgExists := s.Find("div.articleItem-img").Find("img").Attr("src")
		if imgExists {
			article.ImgUrl = img
		}

		article.Title = s.Find("h2.articleTitle").Text()

		listArticles = append(listArticles, article)
	})

	return listArticles
}

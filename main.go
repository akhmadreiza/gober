package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akhmadreiza/gober/models"
	"github.com/akhmadreiza/gober/parsers"
	"github.com/akhmadreiza/gober/scraper"
	"github.com/gin-gonic/gin"
)

type GoberResp struct {
	Status   string           `json:"status"`
	Count    int              `json:"count"`
	Articles []models.Article `json:"articles"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[GOBER] ")
	initRouter()
}

func initRouter() {
	router := gin.Default()
	router.GET("/articles/popular", getPopularArticle)
	router.GET("/articles", searchArticle)
	router.GET("/article", articleDetail)
	router.Run(":8080")
}

func articleDetail(ginContext *gin.Context) {
	website := ginContext.DefaultQuery("source", "detik")
	detailUrl := ginContext.Query("detailUrl")

	log.Println("source:", website, "detail url:", detailUrl)

	if detailUrl == "" {
		ginContext.IndentedJSON(http.StatusBadRequest, gin.H{
			"desc":   "param detailUrl is not exists or is empty",
			"status": "Failed",
		})
		return
	}

	scraper, err := getScraper(website)
	if err != nil {
		fmt.Println("Error when getting scraper:", err)
		ginContext.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
		return
	}

	article, err := scraper.Detail(detailUrl)
	if err != nil {
		fmt.Println("Error scrap URL:", err)
		return
	}

	var listArticles []models.Article
	listArticles = append(listArticles, article)
	resp := GoberResp{
		Status:   "Success",
		Count:    1,
		Articles: listArticles,
	}

	ginContext.IndentedJSON(http.StatusOK, resp)
}

func searchArticle(ginContext *gin.Context) {
	website := ginContext.DefaultQuery("source", "detik")
	searchKey := ginContext.Query("q")

	log.Println("source:", website, "search key:", searchKey)

	if searchKey == "" {
		ginContext.IndentedJSON(http.StatusBadRequest, gin.H{
			"desc":   "param q is not exists or is empty",
			"status": "Failed",
		})
		return
	}

	scraper, err := getScraper(website)
	if err != nil {
		fmt.Println("Error when getting scraper:", err)
		ginContext.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
		return
	}

	articles, err := scraper.Search(searchKey, ginContext)
	if err != nil {
		fmt.Println("Error scrap URL:", err)
		return
	}

	resp := GoberResp{
		Status:   "Success",
		Count:    len(articles),
		Articles: articles,
	}

	ginContext.IndentedJSON(http.StatusOK, resp)
}

func getPopularArticle(ginContext *gin.Context) {
	website := ginContext.DefaultQuery("source", "detik")
	log.Println("source:", website)

	scraper, err := getScraper(website)
	if err != nil {
		fmt.Println("Error when getting scraper:", err)
		ginContext.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
		return
	}

	popArticles, err := scraper.Popular(ginContext)
	if err != nil {
		fmt.Println("Error getting popular news:", err)
		return
	}

	resp := GoberResp{
		Status:   "Success",
		Count:    len(popArticles),
		Articles: popArticles,
	}

	ginContext.IndentedJSON(http.StatusOK, resp)

}

func getScraper(website string) (scraper.NewsScraper, error) {
	if website == "detik" {
		return parsers.DetikScraper{}, nil
	} else if website == "kompas" {
		return parsers.KompasScraper{}, nil
	}
	return nil, fmt.Errorf("scrape %v not supported", website)
}

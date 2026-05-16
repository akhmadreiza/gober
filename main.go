package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/akhmadreiza/gober/models"
	"github.com/akhmadreiza/gober/parsers"
	"github.com/akhmadreiza/gober/scraper"
	"github.com/akhmadreiza/gober/utils"
	"github.com/gin-gonic/gin"
)

type GoberResp struct {
	Status   string           `json:"status"`
	Count    int              `json:"count"`
	Articles []models.Article `json:"articles"`
}

var httpClient *utils.RealHTTPClient
var scrapeUtils utils.ScrapeUtils
var cache *utils.Cache

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[GOBER] ")

	httpClient = utils.NewHTTPClient()
	scrapeUtils = utils.NewScrapeUtils(*httpClient)
	cache = utils.NewCache()

	initRouter()
}

func initRouter() {
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	router := gin.Default()
	router.Use(utils.RateLimitMiddleware())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.Static("/static", "./static")
	router.NoRoute(serveStatic)
	router.GET("/articles/popular", getPopularArticle)
	router.GET("/articles", searchArticle)
	router.GET("/article", articleDetail)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()
	log.Printf("server started on :%s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("received %v, shutting down gracefully...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}
	log.Println("server stopped")
}

// isAllowedURL rejects requests that don't target a known news domain,
// preventing SSRF via the detailUrl parameter.
func isAllowedURL(rawURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Host == "" {
		return false
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}
	host := parsed.Hostname()
	return host == "detik.com" || strings.HasSuffix(host, ".detik.com") ||
		host == "kompas.com" || strings.HasSuffix(host, ".kompas.com")
}

func serveStatic(c *gin.Context) {
	staticPath := "./static/index.html"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	c.File(staticPath)
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

	if !isAllowedURL(detailUrl) {
		log.Printf("blocked disallowed detailUrl: %s", detailUrl)
		ginContext.IndentedJSON(http.StatusBadRequest, gin.H{
			"desc":   "detailUrl points to a disallowed domain",
			"status": "Failed",
		})
		return
	}

	scraper, err := getScraper(website)
	if err != nil {
		log.Printf("Error when getting scraper: %v", err)
		ginContext.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
		return
	}

	article, err := scraper.Detail(detailUrl, ginContext)
	if err != nil {
		log.Printf("Error scrap URL: %v", err)
		ginContext.IndentedJSON(http.StatusInternalServerError, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
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
		log.Printf("Error when getting scraper: %v", err)
		ginContext.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
		return
	}

	articles, err := scraper.Search(searchKey, ginContext)
	if err != nil {
		log.Printf("Error scrap URL: %v", err)
		ginContext.IndentedJSON(http.StatusInternalServerError, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
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
		log.Printf("Error when getting scraper: %v", err)
		ginContext.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
		return
	}

	popArticles, err := scraper.Popular(ginContext)
	if err != nil {
		log.Printf("Error getting popular news: %v", err)
		ginContext.IndentedJSON(http.StatusInternalServerError, gin.H{
			"desc":   err.Error(),
			"status": "Failed",
		})
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
		return parsers.DetikScraper{Client: httpClient, Utils: scrapeUtils, Cache: cache}, nil
	} else if website == "kompas" {
		return parsers.KompasScraper{Client: httpClient, Utils: scrapeUtils, Cache: cache}, nil
	}
	return nil, fmt.Errorf("scrape %v not supported", website)
}

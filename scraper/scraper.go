package scraper

import (
	"github.com/akhmadreiza/gober/models"
	"github.com/gin-gonic/gin"
)

type NewsScraper interface {
	Search(keyword string, ginContext *gin.Context) ([]models.Article, error)
	Popular(ginContext *gin.Context) ([]models.Article, error)
	Detail(url string) (models.Article, error)
}

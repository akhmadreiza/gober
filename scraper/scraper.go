package scraper

import "github.com/akhmadreiza/gober/models"

type NewsScraper interface {
	Search(keyword string) ([]models.Article, error)
	Popular() ([]models.Article, error)
}

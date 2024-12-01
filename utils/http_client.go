package utils

import (
	"github.com/akhmadreiza/gober/models"
)

type HTTPClient interface {
	Get(url string) (models.ScraperResponse, error)
}

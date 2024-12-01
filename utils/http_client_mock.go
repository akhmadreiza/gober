package utils

import "github.com/akhmadreiza/gober/models"

type HttpClientMock struct {
	Response models.ScraperResponse
	Err      error
}

func (m HttpClientMock) Get(url string) (models.ScraperResponse, error) {
	return m.Response, m.Err
}

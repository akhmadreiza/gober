package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/akhmadreiza/gober/models"
)

type RealHTTPClient struct {
	Client *http.Client
}

func NewHTTPClient() *RealHTTPClient {
	return &RealHTTPClient{
		Client: &http.Client{},
	}
}

func (h RealHTTPClient) Get(url string) (sr models.ScraperResponse, err error) {
	resp, err := h.Client.Get(url)
	if err != nil {
		return models.ScraperResponse{}, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ScraperResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return models.ScraperResponse{
		Body:   string(body),
		Status: resp.StatusCode,
		WebUrl: url,
	}, nil
}

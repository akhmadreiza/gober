package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/akhmadreiza/gober/models"
)

type RealHTTPClient struct {
	Client *http.Client
}

func NewHTTPClient() *RealHTTPClient {
	return &RealHTTPClient{
		Client: &http.Client{Timeout: 15 * time.Second},
	}
}

const userAgent = "Gober/1.0 (+https://github.com/akhmadreiza/gober)"

func (h RealHTTPClient) Get(rawURL string) (sr models.ScraperResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return models.ScraperResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := h.Client.Do(req)
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
		WebUrl: rawURL,
	}, nil
}

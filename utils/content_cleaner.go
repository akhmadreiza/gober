package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var defaultRemoveSelectors = []string{
	"script",
	"style",
	"iframe",
	"noscript",
	`[class*="ads"]`,
}

func CleanContent(s *goquery.Selection, extraSelectors ...string) string {
	selectors := append(defaultRemoveSelectors, extraSelectors...)
	for _, sel := range selectors {
		s.Find(sel).Remove()
	}
	html, _ := s.Html()
	return strings.TrimSpace(html)
}

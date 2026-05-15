package utils

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var detikImgWidth = regexp.MustCompile(`w=\d+`)
var kompasDimension = regexp.MustCompile(`/\d+x\d+/`)

// EnhanceImageURL rewrites thumbnail CDN URLs to request a larger image.
func EnhanceImageURL(imgURL string) string {
	if strings.Contains(imgURL, "akcdn.detik.net.id") {
		return detikImgWidth.ReplaceAllString(imgURL, "w=800")
	}
	if strings.Contains(imgURL, "asset.kompas.com") {
		return kompasDimension.ReplaceAllLiteralString(imgURL, "/460x306/")
	}
	return imgURL
}

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

// RewriteContentLinks rewrites internal news links (detik.com, kompas.com)
// to point to Gober's own /detail route, keeping readers on the app.
func RewriteContentLinks(s *goquery.Selection) {
	s.Find("a[href]").Each(func(_ int, a *goquery.Selection) {
		href, _ := a.Attr("href")
		parsed, err := url.Parse(href)
		if err != nil || parsed.Host == "" {
			return
		}

		var source string
		switch {
		case strings.Contains(parsed.Host, "detik.com"):
			source = "detik"
			if !strings.Contains(href, "single=1") {
				if parsed.RawQuery != "" {
					href += "&single=1"
				} else {
					href += "?single=1"
				}
			}
		case strings.Contains(parsed.Host, "kompas.com"):
			source = "kompas"
			if !strings.Contains(href, "page=all") {
				if parsed.RawQuery != "" {
					href += "&page=all"
				} else {
					href += "?page=all"
				}
			}
		default:
			return
		}

		a.SetAttr("href", "/detail?source="+source+"&detailUrl="+url.QueryEscape(href))
		a.RemoveAttr("onclick")
		a.RemoveAttr("target")
	})
}

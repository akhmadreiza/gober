package utils_test

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/akhmadreiza/gober/utils"
	"github.com/stretchr/testify/assert"
)

func TestCleanContentRemovesDefaultSelectors(t *testing.T) {
	html := `<div>
		<p>Article paragraph one.</p>
		<script>googletag.cmd.push(function() {});</script>
		<style>@import url("ads.css");</style>
		<iframe src="https://survey.example.com"></iframe>
		<noscript>Please enable JavaScript</noscript>
		<div class="ads-scrollpage-container">Ad content</div>
		<span class="ads-on-body"></span>
		<p>Article paragraph two.</p>
	</div>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	result := utils.CleanContent(doc.Find("div").First())

	assert.NotContains(t, result, "<script>")
	assert.NotContains(t, result, "<style>")
	assert.NotContains(t, result, "<iframe>")
	assert.NotContains(t, result, "<noscript>")
	assert.NotContains(t, result, "ads-scrollpage-container")
	assert.NotContains(t, result, "ads-on-body")
	assert.Contains(t, result, "Article paragraph one.")
	assert.Contains(t, result, "Article paragraph two.")
}

func TestCleanContentRemovesExtraSelectors(t *testing.T) {
	html := `<div>
		<p>Article content.</p>
		<div class="parallaxindetail">Parallax ad</div>
		<div class="kompasidRec">Recommendation widget</div>
		<div class="staticdetail_container">Static ad</div>
		<p class="para_caption">ADVERTISEMENT</p>
	</div>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	result := utils.CleanContent(doc.Find("div").First(),
		".parallaxindetail",
		".kompasidRec",
		".staticdetail_container",
		".para_caption",
	)

	assert.NotContains(t, result, "parallaxindetail")
	assert.NotContains(t, result, "kompasidRec")
	assert.NotContains(t, result, "staticdetail_container")
	assert.NotContains(t, result, "ADVERTISEMENT")
	assert.Contains(t, result, "Article content.")
}

func TestRewriteContentLinksDetik(t *testing.T) {
	html := `<div>
		<p>Baca juga: <a href="https://news.detik.com/berita/d-123/some-article" onclick="_pt(this)" target="_blank">Some Article</a></p>
		<p>External: <a href="https://example.com/page">External Link</a></p>
	</div>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	sel := doc.Find("div").First()
	utils.RewriteContentLinks(sel)

	result, _ := sel.Html()
	assert.Contains(t, result, `/detail?source=detik&amp;detailUrl=`)
	assert.Contains(t, result, `single%3D1`) // "single=1" URL-encoded inside detailUrl
	assert.NotContains(t, result, `onclick`)
	assert.NotContains(t, result, `target=`)
	assert.Contains(t, result, `href="https://example.com/page"`) // external unchanged
}

func TestRewriteContentLinksKompas(t *testing.T) {
	html := `<div>
		<p>Baca juga: <a href="https://nasional.kompas.com/read/2024/01/01/some-article">Some Article</a></p>
	</div>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	sel := doc.Find("div").First()
	utils.RewriteContentLinks(sel)

	result, _ := sel.Html()
	assert.Contains(t, result, `/detail?source=kompas&amp;detailUrl=`)
	assert.Contains(t, result, `page%3Dall`) // "page=all" URL-encoded inside detailUrl
}

func TestCleanContentPreservesArticleMarkup(t *testing.T) {
	html := `<div>
		<p><strong>Bold text</strong> and <em>italic text</em>.</p>
		<p><a href="https://example.com">A link</a></p>
		<img src="https://example.com/image.jpg" />
	</div>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	result := utils.CleanContent(doc.Find("div").First())

	assert.Contains(t, result, "<strong>Bold text</strong>")
	assert.Contains(t, result, "<em>italic text</em>")
	assert.Contains(t, result, `<a href="https://example.com">A link</a>`)
	assert.Contains(t, result, `<img src="https://example.com/image.jpg"/>`)
}

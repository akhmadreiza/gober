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

package parsers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/akhmadreiza/gober/models"
	"github.com/akhmadreiza/gober/parsers"
	"github.com/akhmadreiza/gober/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ginContext = createTestGinContext()

func TestDetailDetik(t *testing.T) {

	//prepare data
	mockHTML := `
	<html>
		<h1 class="detail__title">Test Title</h1>
		<div class="detail__author">Test Author</div>
		<div class="detail__date">2024-11-29</div>
		<div class="detail__media"><img src="https://example.com/image.jpg" /></div>
		<div class="detail__body-text itp_bodycontent">Test Content</div>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	cache := utils.NewCache()
	scraper := parsers.DetikScraper{Client: mockClient, Utils: util, Cache: cache}
	result, err := scraper.Detail("https://detik.com", ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, "Test Title", result.Title)
	assert.Equal(t, "Test Author", result.Author)
	assert.Equal(t, "2024-11-29", result.Date)
	assert.Equal(t, "https://example.com/image.jpg", result.ImgUrl)
	assert.Equal(t, "Test Content", result.Content)
}

func TestSearchDetik(t *testing.T) {
	//prepare data
	mockHTML := `
	<html>
		<article class="list-content__item">
			<div class="media media--left media--image-radius block-link">
				<div class="media__image">
					<!-- <span class="media__label bg-blue font-xs">LABEL</span> -->
					<a href="https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak&quot;, &quot;artikel newsfeed&quot;)">
						<span class="ratiobox ratiobox--4-3 lqd" style="background-image: url(&quot;https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&amp;q=90&quot;);">
								<img src="https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&amp;q=90" alt="Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak" title="Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak" class="" style="display: none;">					                </span>
					</a>
				</div>
				<div class="media__text">
															<h3 class="media__title">
						<a href="https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak&quot;, &quot;artikel newsfeed&quot;)">Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak</a>
					</h3>
					<div class="media__date">detikNews | <span d-time="1733070610" title="Minggu, 01 Des 2024 23:30 WIB">1 jam yang lalu</span></div>
				</div>
			</div>
		</article>
		<article class="list-content__item">
			<div class="media media--left media--image-radius block-link">
				<div class="media__image">
					<!-- <span class="media__label bg-blue font-xs">LABEL</span> -->
					<a href="https://news.detik.com/berita/d-7666189/polisi-ungkap-kondisi-ibu-korban-penusukan-abg-di-cilandak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Ungkap Kondisi Ibu Korban Penusukan ABG di Cilandak&quot;, &quot;artikel newsfeed&quot;)">
						<span class="ratiobox ratiobox--4-3 lqd" style="background-image: url(&quot;https://akcdn.detik.net.id/community/media/visual/2024/12/01/polisi-ungkap-kondisi-abg-pembunuh-ayah-ibu-di-cilandak_43.jpeg?w=220&amp;q=90&quot;);">
								<img src="https://akcdn.detik.net.id/community/media/visual/2024/12/01/polisi-ungkap-kondisi-abg-pembunuh-ayah-ibu-di-cilandak_43.jpeg?w=220&amp;q=90" alt="Polisi Ungkap Kondisi Ibu Korban Penusukan ABG di Cilandak" title="Polisi Ungkap Kondisi Ibu Korban Penusukan ABG di Cilandak" class="" style="display: none;">					                </span>
					</a>
				</div>
				<div class="media__text">
															<h3 class="media__title">
						<a href="https://news.detik.com/berita/d-7666189/polisi-ungkap-kondisi-ibu-korban-penusukan-abg-di-cilandak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Ungkap Kondisi Ibu Korban Penusukan ABG di Cilandak&quot;, &quot;artikel newsfeed&quot;)">Polisi Ungkap Kondisi Ibu Korban Penusukan ABG di Cilandak</a>
					</h3>
					<div class="media__date">detikNews | <span d-time="1733072746" title="Senin, 02 Des 2024 00:05 WIB">1 jam yang lalu</span></div>
				</div>
			</div>
		</article>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	cache := utils.NewCache()
	scraper := parsers.DetikScraper{Client: mockClient, Utils: util, Cache: cache}
	result, err := scraper.Search("anak abah", ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak", result[0].Title)
	assert.Equal(t, "", result[0].Author)
	assert.Equal(t, "Minggu, 01 Des 2024 23:30 WIB", result[0].Date)
	assert.Equal(t, "https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&q=90", result[0].ImgUrl)
}

func TestPopularDetikNoResult(t *testing.T) {
	//prepare data
	mockHTML := `
	<html></html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	cacheItems := utils.CacheItemsMock{
		Data:  nil,
		Found: false,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.DetikScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))

	//make sure no data inserted to cache if result length is 0
	_, ok := cache.Items.Data.([]models.Article)
	assert.False(t, ok)
}

func TestPopularDetikNoCache(t *testing.T) {
	//prepare data
	mockHTML := `
	<html>
		<article class="list-content__item">
			<div class="media media--left media--image-radius block-link">
				<div class="media__image">
					<!-- <span class="media__label bg-blue font-xs">LABEL</span> -->
					<a href="https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak&quot;, &quot;artikel newsfeed&quot;)">
						<span class="ratiobox ratiobox--4-3 lqd" style="background-image: url(&quot;https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&amp;q=90&quot;);">
								<img src="https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&amp;q=90" alt="Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak" title="Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak" class="" style="display: none;">					                </span>
					</a>
				</div>
				<div class="media__text">
															<h3 class="media__title">
						<a href="https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak&quot;, &quot;artikel newsfeed&quot;)">Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak</a>
					</h3>
					<div class="media__date">detikNews | <span d-time="1733070610" title="Minggu, 01 Des 2024 23:30 WIB">1 jam yang lalu</span></div>
				</div>
			</div>
		</article>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	cacheItems := utils.CacheItemsMock{
		Data:  nil,
		Found: false,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.DetikScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 11, len(res)) //11 is sum of popUrls slice in detik_parser

	//make sure data inserted to cache
	resdata, ok := cache.Items.Data.([]models.Article)
	assert.True(t, ok)
	assert.NotNil(t, resdata)
	assert.Equal(t, 11, len(resdata))
}

func TestPopularDetikWithCache(t *testing.T) {
	//prepare data
	mockHTML := `
	<html>
		<article class="list-content__item">
			<div class="media media--left media--image-radius block-link">
				<div class="media__image">
					<!-- <span class="media__label bg-blue font-xs">LABEL</span> -->
					<a href="https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak&quot;, &quot;artikel newsfeed&quot;)">
						<span class="ratiobox ratiobox--4-3 lqd" style="background-image: url(&quot;https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&amp;q=90&quot;);">
								<img src="https://akcdn.detik.net.id/community/media/visual/2024/12/01/wisatawan-pakai-pelat-palsu-polri-di-puncak_43.jpeg?w=220&amp;q=90" alt="Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak" title="Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak" class="" style="display: none;">					                </span>
					</a>
				</div>
				<div class="media__text">
															<h3 class="media__title">
						<a href="https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak" class="media__link" onclick="_pt(this, &quot;newsfeed&quot;, &quot;Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak&quot;, &quot;artikel newsfeed&quot;)">Polisi Tindak Wisatawan Pakai Pelat Palsu Polri Demi Lolos Gage di Puncak</a>
					</h3>
					<div class="media__date">detikNews | <span d-time="1733070610" title="Minggu, 01 Des 2024 23:30 WIB">1 jam yang lalu</span></div>
				</div>
			</div>
		</article>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	var a []models.Article
	a = append(a, models.Article{
		URL: "https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak",
	})
	cacheItems := utils.CacheItemsMock{
		Data:  a,
		Found: true,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.DetikScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))

	//make sure data inserted to cache
	resdata, ok := cache.Items.Data.([]models.Article)
	assert.True(t, ok)
	assert.NotNil(t, resdata)
	assert.Equal(t, 1, len(resdata))
}

func TestSearchKompas(t *testing.T) {
	//prepare data
	mockHTML := `<html></html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	cache := utils.NewCache()
	scraper := parsers.KompasScraper{Client: mockClient, Utils: util, Cache: cache}
	_, err := scraper.Search("anak abah", ginContext)

	//assertions
	assert.EqualError(t, err, "KompasScraper Search is not supported")
}

func TestDetailKompas(t *testing.T) {

	//prepare data
	mockHTML := `
	<html>
		<h1 class="read__title">Test Title</h1>
		<div class="credit-title">Test Author</div>
		<div class="read__time">2024-11-29</div>
		<div class="photo__wrap"><img src="https://example.com/image.jpg" /></div>
		<div class="read__content">Test Content</div>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	cache := utils.NewCache()
	scraper := parsers.KompasScraper{Client: mockClient, Utils: util, Cache: cache}
	result, err := scraper.Detail("https://kompas.com", ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, "Test Title", result.Title)
	assert.Equal(t, "Test Author", result.Author)
	assert.Equal(t, "2024-11-29", result.Date)
	assert.Equal(t, "https://example.com/image.jpg", result.ImgUrl)
	assert.Equal(t, "Test Content", result.Content)
}

func TestPopularKompasNoResult(t *testing.T) {
	//prepare data
	mockHTML := `
	<html></html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	cacheItems := utils.CacheItemsMock{
		Data:  nil,
		Found: false,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.KompasScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))

	//make sure no data inserted to cache if result length is 0
	_, ok := cache.Items.Data.([]models.Article)
	assert.False(t, ok)
}

func TestPopularKompasNoCache(t *testing.T) {
	//prepare data
	mockHTML := `
	<html>
		<div class="articleList -list">
			<div class="articleItem">
				<a class="article-link" href="https://nasional.kompas.com/read/2024/12/26/12015311/donny-tri-istiomah-tersangka-akui-pernah-dititipi-uang-ratusan-juta-oleh">
					<div class="articleItem-wrap">
						<div class="articleItem-img">
							<img src="https://asset.kompas.com/crops/jPMzGqXbL3W2mnqLOP_quLuKL8M=/0x0:580x387/230x153/data/photo/2024/12/24/676aa403e48f6.jpg" alt="">
						</div>
						<div class="articleItem-box">
							<h2 class="articleTitle">Donny Tri Istiomah Tersangka, Akui Pernah Dititipi Uang Ratusan Juta oleh Harun Masiku</h2>
							<div class="articlePost">
								<ul>
									<li>
										<div class="articlePost-subtitle">NEWS</div>
									</li>
									<li>
										<div class="articlePost-date">26/12/2024</div>
									</li>
								</ul>
							</div>
						</div>
					</div>
				</a>
			</div>
		</div>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	cacheItems := utils.CacheItemsMock{
		Data:  nil,
		Found: false,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.KompasScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(res)) //2 is sum of popUrls slice in kompas_parser

	//make sure data inserted to cache
	resdata, ok := cache.Items.Data.([]models.Article)
	assert.True(t, ok)
	assert.NotNil(t, resdata)
	assert.Equal(t, 2, len(resdata))
}

func TestPopularKompasWithCache(t *testing.T) {
	//prepare data
	mockHTML := `
	<html>
		<div class="articleList -list">
			<div class="articleItem">
				<a class="article-link" href="https://nasional.kompas.com/read/2024/12/26/12015311/donny-tri-istiomah-tersangka-akui-pernah-dititipi-uang-ratusan-juta-oleh">
					<div class="articleItem-wrap">
						<div class="articleItem-img">
							<img src="https://asset.kompas.com/crops/jPMzGqXbL3W2mnqLOP_quLuKL8M=/0x0:580x387/230x153/data/photo/2024/12/24/676aa403e48f6.jpg" alt="">
						</div>
						<div class="articleItem-box">
							<h2 class="articleTitle">Donny Tri Istiomah Tersangka, Akui Pernah Dititipi Uang Ratusan Juta oleh Harun Masiku</h2>
							<div class="articlePost">
								<ul>
									<li>
										<div class="articlePost-subtitle">NEWS</div>
									</li>
									<li>
										<div class="articlePost-date">26/12/2024</div>
									</li>
								</ul>
							</div>
						</div>
					</div>
				</a>
			</div>
		</div>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	var a []models.Article
	a = append(a, models.Article{
		URL: "https://news.detik.com/berita/d-7666179/polisi-tindak-wisatawan-pakai-pelat-palsu-polri-demi-lolos-gage-di-puncak",
	})
	cacheItems := utils.CacheItemsMock{
		Data:  a,
		Found: true,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.KompasScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))

	//make sure data inserted to cache
	resdata, ok := cache.Items.Data.([]models.Article)
	assert.True(t, ok)
	assert.NotNil(t, resdata)
	assert.Equal(t, 1, len(resdata))
}

func TestPopularKompasNoArticleTimeFromURL(t *testing.T) {
	//prepare data
	mockHTML := `
	<html>
		<div class="articleList -list">
			<div class="articleItem">
				<a class="article-link" href="https://jeo.kompas.com/mencari-harimau-jawa-antara-ada-dan-tiada">
					<div class="articleItem-wrap">
						<div class="articleItem-img">
							<img src="https://asset.kompas.com/crops/TCuq_DNPXldyvIS6MJVs9R-AF3w=/71x0:885x543/230x153/data/photo/2025/01/01/677527db4db98.jpg" alt="">
						</div>
						<div class="articleItem-box">
							<h2 class="articleTitle">Mencari Harimau Jawa, Antara Ada dan Tiada</h2>
							<div class="articlePost">
								<ul>
									<li>
										<div class="articlePost-subtitle">JEO</div>
									</li>
									<li>
										<div class="articlePost-date">02/01/2025</div>
									</li>
								</ul>
							</div>
						</div>
					</div>
				</a>
			</div>
		</div>
	</html>`

	//mock
	mockClient := utils.HttpClientMock{
		Response: models.ScraperResponse{
			Body:   mockHTML,
			Status: 200,
		},
	}
	cacheItems := utils.CacheItemsMock{
		Data:  nil,
		Found: false,
	}
	cache := utils.CacheMock{
		Items: cacheItems,
	}

	//do test
	util := utils.NewScrapeUtils(mockClient)
	scraper := parsers.KompasScraper{Client: mockClient, Utils: util, Cache: &cache}
	res, err := scraper.Popular(ginContext)

	//assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(res)) //2 is sum of popUrls slice in kompas_parser

	//make sure data inserted to cache
	resdata, ok := cache.Items.Data.([]models.Article)
	assert.True(t, ok)
	assert.NotNil(t, resdata)
	assert.Equal(t, 2, len(resdata))
}

func createTestGinContext() *gin.Context {
	// Create a ResponseRecorder (a test HTTP response writer)
	w := httptest.NewRecorder()
	// Create a new Gin context
	c, _ := gin.CreateTestContext(w)

	// Simulate an HTTP request for context purposes
	req := httptest.NewRequest("GET", "/", nil)
	c.Request = req

	return c
}

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What is Gober

Gober (Go Berita) is a monorepo news aggregator with a Go backend and Vue.js frontend. The backend scrapes Indonesian news sites (detik.com, kompas.com) and exposes REST APIs. The Vue.js frontend consumes those APIs and is served as static files from the Go server in production.

## Commands

### Backend (Go)
```bash
go run main.go          # Run backend server on :8080
go test ./...           # Run all tests
go test ./parsers/...   # Run tests in a specific package
go build .              # Build binary
go mod tidy             # Sync dependencies
```

### Frontend (Vue.js — run from `web/` directory)
```bash
npm run serve   # Dev server on :8081 (proxies API to Go backend at localhost:8080)
npm run build   # Build for production
npm run lint    # Lint
```

### Full build & deploy
```bash
make build-fe   # Build frontend → moves dist to ../static
make build-be   # Build Go binary
make build-gober       # Both of the above
make install-gober     # Build + install to /opt/gober/
```

## Architecture

### Request flow
1. Browser → Gin router (`main.go`)
2. Route handler calls `getScraper(source)` which returns a `NewsScraper` for the requested site
3. Scraper checks in-memory `Cache` (5 min TTL); on miss, calls `HTTPClient.Get()` and parses HTML with goquery
4. `ScrapeUtils.FetchListArticles` fans out popular article fetches concurrently via goroutines + channels

### Key interfaces (dependency injection enables mocking in tests)
- `scraper.NewsScraper` — `Search`, `Popular`, `Detail` methods; implemented by `DetikScraper` and `KompasScraper` in `parsers/`
- `utils.HTTPClient` — wraps `http.Get`; `RealHTTPClient` for prod, `HttpClientMock` for tests
- `utils.CacheOps` — `Get`/`Set`; `Cache` (sync.RWMutex in-memory) for prod, `CacheMock` for tests

### Content cleaning pipeline
Every `Detail()` implementation must run two steps on the content selection before returning:

```go
utils.RewriteContentLinks(content)   // rewrites detik/kompas hrefs → /detail?source=...
article.Content = utils.CleanContent(content, /* site-specific selectors */)
```

`CleanContent` (`utils/content_cleaner.go`) removes `script`, `style`, `iframe`, `noscript`, and any element whose class contains `"ads"` by default. Extra site-specific selectors (ad container class names, widget divs) are passed as variadic args.

`RewriteContentLinks` rewrites `<a href>` tags pointing to detik.com or kompas.com to `/detail?source=SOURCE&detailUrl=ENCODED_URL`, appending `?single=1` (Detik) or `?page=all` (Kompas), and strips `onclick` tracking attributes.

### Adding a new news source
1. Create `parsers/<site>_parser.go` implementing `scraper.NewsScraper`
2. Register the source string in `getScraper()` in `main.go`
3. In the `Detail()` method, call `RewriteContentLinks` then `CleanContent` with site-specific ad selectors (inspect real article HTML to find them)
4. Add the new source domain to `RewriteContentLinks` in `utils/content_cleaner.go`
5. Add test cases to `parsers/parser_test.go` using mock HTTP client and cache

### Frontend ↔ Backend integration
In production, the frontend is built into `static/` and served by Gin (`router.Static("/static", "./static")`). All unmatched routes fall back to `static/index.html` (SPA mode).

In development, `vue.config.js` proxies `/articles/popular` and `/article` to `VUE_APP_GOBER_API_URL` (set in `web/.env.development`, defaults to `http://localhost:8080`). The Vue dev server runs on `:8081` because `:8080` is taken by Go.

### API endpoints
| Endpoint | Query params | Notes |
|---|---|---|
| `GET /articles/popular` | `source=detik\|kompas` | Returns popular articles; cached 5 min |
| `GET /articles` | `source=`, `q=keyword` | Search (detik only; kompas not supported) |
| `GET /article` | `source=`, `detailUrl=` | Fetch full article detail; cached 5 min |

The `detailUrl` passed to `/article` is URL-encoded. Article URLs in list responses are already pre-encoded proxied URLs (e.g., `http://host/article?detailUrl=...`).

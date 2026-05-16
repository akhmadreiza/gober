<template>
  <div class="static-page">

    <nav class="top-nav">
      <router-link to="/" class="nav-back">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 12H5M12 5l-7 7 7 7"/></svg>
        Gober
      </router-link>
    </nav>

    <main class="page-container">

      <header class="page-header">
        <p class="page-kicker">About</p>
        <h1 class="page-title">How Gober Works</h1>
        <div class="ornament-row">
          <span class="orn-line"></span>
          <span class="orn-diamond">◆</span>
          <span class="orn-line"></span>
        </div>
        <p class="page-intro">
          Gober is a free, open-source Indonesian news reader that strips away advertisements,
          pop-ups, and distractions — serving only the article you came to read.
        </p>
      </header>

      <section class="page-section">
        <h2 class="section-title">The Problem It Solves</h2>
        <p>
          Reading news on Indonesian media websites today often means navigating through
          layers of banner ads, auto-playing videos, cookie banners, and paywalls that
          exist between you and the story. Gober is built on a simple premise: <em>the news
          should be easy to read.</em>
        </p>
        <p>
          Gober does not host, own, or produce any news content. It acts solely as a
          reading interface — fetching articles from their original sources and presenting
          them in a clean, typography-first layout.
        </p>
      </section>

      <section class="page-section">
        <h2 class="section-title">Gober vs. Ad-block Extensions</h2>
        <p>
          The honest answer is: <strong>if you already use a good ad-blocker, it solves
          most of the same problems.</strong> Extensions like uBlock Origin are free,
          mature, and highly effective at removing ads, trackers, and auto-playing videos
          from Indonesian news sites — often better than Gober at handling dynamically
          injected ads that vary with each page load.
        </p>
        <p>
          If you are comfortable installing a browser extension and you mostly read news
          on a desktop, an ad-blocker is probably the simpler, more reliable choice.
          There is no need to go through a separate reading interface.
        </p>

        <div class="compare-grid">
          <div class="compare-col">
            <h3 class="compare-heading where-gober">Where Gober has an edge</h3>
            <ul class="compare-list">
              <li>
                <strong>No extension required.</strong> Works on any browser and device
                — including iOS Safari, which does not support conventional extensions,
                and shared or locked-down computers where you cannot install software.
              </li>
              <li>
                <strong>Consistent reading layout.</strong> Every article is rendered
                in the same clean typography regardless of the original site's design.
                No layout shifts, no cookie banners that survived the blocker, no
                "please disable your ad-blocker" nag screens.
              </li>
              <li>
                <strong>Server-side cleaning.</strong> Because Gober strips ads before
                the HTML ever reaches your browser, there is nothing for anti-adblock
                scripts to detect. Sites that actively fight ad-blockers behave normally
                on Gober.
              </li>
              <li>
                <strong>Lighter page loads.</strong> Only the editorial content is
                transferred to your browser — not the dozens of third-party scripts,
                fonts, and trackers that come with the original page.
              </li>
            </ul>
          </div>
          <div class="compare-col">
            <h3 class="compare-heading where-adblock">Where an ad-blocker wins</h3>
            <ul class="compare-list">
              <li>
                <strong>Broader coverage.</strong> A good ad-blocker works across every
                website you visit, not just the sources Gober supports.
              </li>
              <li>
                <strong>Always up to date.</strong> Filter lists like EasyList are
                updated daily by large communities. Gober's parsers may break silently
                when a news site redesigns its HTML structure.
              </li>
              <li>
                <strong>Richer media support.</strong> Ad-blockers still show embedded
                videos, galleries, and interactive graphics. Gober strips most of these
                since they typically come bundled with tracking and ad code.
              </li>
              <li>
                <strong>No dependency on a third-party server.</strong> An ad-blocker
                works offline and does not route your reading through someone else's
                backend. Gober is open source and self-hostable, but if this hosted
                instance goes down, it is unavailable.
              </li>
            </ul>
          </div>
        </div>

        <p class="compare-summary">
          In short, Gober and ad-blockers are complementary tools with different
          trade-offs. Gober is best suited for mobile browsers, situations where
          extensions are not an option, or readers who simply prefer a unified
          reading environment. It is not trying to replace ad-blockers — it is
          a different kind of solution to a shared problem.
        </p>
      </section>

      <section class="page-section">
        <h2 class="section-title">How a Request Flows</h2>
        <p>
          When you open Gober and browse articles, here is what happens behind the scenes:
        </p>
        <ol class="flow-list">
          <li>
            <span class="step-num">01</span>
            <div>
              <strong>You visit Gober.</strong> The Vue.js frontend loads in your browser
              and immediately requests the list of popular articles from the Gober backend.
            </div>
          </li>
          <li>
            <span class="step-num">02</span>
            <div>
              <strong>The backend checks its cache.</strong> Gober keeps a short-lived
              in-memory cache (5-minute TTL). If the article list was recently fetched,
              it is returned instantly without hitting the original site again.
            </div>
          </li>
          <li>
            <span class="step-num">03</span>
            <div>
              <strong>On a cache miss, Gober fetches the source website.</strong> The Go
              backend makes HTTP requests to the original news site (e.g. detik.com or
              kompas.com) — the same request your browser would make — and reads the
              HTML response.
            </div>
          </li>
          <li>
            <span class="step-num">04</span>
            <div>
              <strong>The HTML is parsed and cleaned.</strong> Using <em>goquery</em>
              (a Go HTML parsing library modelled after jQuery), Gober extracts the
              article title, author, date, image, and body content. Script tags, style
              tags, iframes, ad containers, and tracking elements are then removed
              from the content.
            </div>
          </li>
          <li>
            <span class="step-num">05</span>
            <div>
              <strong>Internal links are rewritten.</strong> Any "Baca Juga" (read also)
              links inside an article that point back to the original news site are
              rewritten to go through Gober instead, so you stay in the same clean
              reading experience.
            </div>
          </li>
          <li>
            <span class="step-num">06</span>
            <div>
              <strong>The clean article is returned to your browser.</strong> The frontend
              renders the article using Playfair Display and Lora — the same typefaces
              used in high-quality print journalism — in a single-column layout optimised
              for reading.
            </div>
          </li>
        </ol>
      </section>

      <section class="page-section">
        <h2 class="section-title">The Technology</h2>
        <p>
          Gober is a small monorepo combining a Go backend and a Vue.js frontend,
          deployed as a single binary serving both the API and the static frontend files.
        </p>
        <div class="tech-grid">
          <div class="tech-card">
            <h3>Backend</h3>
            <ul>
              <li>Go 1.23</li>
              <li>Gin — HTTP router and server</li>
              <li>goquery — HTML scraping and cleaning</li>
              <li>In-memory cache with RWMutex</li>
            </ul>
          </div>
          <div class="tech-card">
            <h3>Frontend</h3>
            <ul>
              <li>Vue.js 3</li>
              <li>Vue Router</li>
              <li>Axios — API requests</li>
              <li>Playfair Display &amp; Lora — typography</li>
            </ul>
          </div>
        </div>
      </section>

      <section class="page-section">
        <h2 class="section-title">Currently Supported Sources</h2>
        <p>
          Gober currently supports popular and article detail fetching for:
        </p>
        <ul class="source-list">
          <li><strong>Detik.com</strong> — news, finance, entertainment, sport, tech, travel, food, health, and education sections</li>
          <li><strong>Kompas.com</strong> — headline and most-read sections</li>
        </ul>
        <p>
          Support for additional Indonesian news sources is planned. Each new source
          requires a custom HTML parser to handle its specific page structure and
          ad-injection patterns.
        </p>
      </section>

      <section class="page-section">
        <h2 class="section-title">Open Source</h2>
        <p>
          Gober is fully open source. You can read the code, report issues, suggest
          improvements, or contribute new news source parsers on GitHub.
        </p>
        <a
          href="https://github.com/akhmadreiza/gober"
          target="_blank"
          rel="noopener noreferrer"
          class="github-link"
        >
          <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z"/></svg>
          github.com/akhmadreiza/gober
        </a>
      </section>

    </main>

    <footer class="page-footer">
      <span class="footer-ornament">◆</span>
      <p class="footer-text">Gober — Go Berita. Baca tanpa gangguan iklan.</p>
    </footer>

  </div>
</template>

<script>
export default {
  name: 'HowItWorks',
};
</script>

<style scoped>
.static-page {
  background: var(--paper);
  min-height: 100vh;
  font-family: 'Lora', Georgia, serif;
  color: var(--ink);
}

.top-nav {
  padding: 16px 24px;
  border-bottom: 1px solid var(--border);
  position: sticky;
  top: 0;
  background: var(--nav-bg);
  backdrop-filter: blur(10px);
  z-index: 50;
}

.nav-back {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-family: 'Playfair Display', serif;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--ink-muted);
  text-decoration: none;
  transition: color 0.2s;
}
.nav-back:hover { color: var(--accent); }

.page-container {
  max-width: 720px;
  margin: 0 auto;
  padding: 56px 24px 72px;
}

.page-header {
  text-align: center;
  margin-bottom: 56px;
}

.page-kicker {
  font-family: 'Playfair Display', serif;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: var(--accent);
  margin: 0 0 12px;
}

.page-title {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: clamp(2rem, 6vw, 3.2rem);
  font-weight: 900;
  letter-spacing: 0.02em;
  color: var(--ink);
  margin: 0 0 20px;
  line-height: 1.1;
}

.ornament-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 24px;
}
.orn-line { width: 60px; height: 1px; background: var(--border); }
.orn-diamond { font-size: 0.5rem; color: var(--accent); }

.page-intro {
  font-size: 1.05rem;
  line-height: 1.75;
  color: var(--ink-muted);
  font-style: italic;
  max-width: 560px;
  margin: 0 auto;
}

.page-section {
  margin-bottom: 52px;
}

.section-title {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--ink);
  margin: 0 0 16px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--border);
  letter-spacing: 0.02em;
}

p {
  font-size: 0.97rem;
  line-height: 1.85;
  color: var(--ink);
  margin: 0 0 16px;
}

.flow-list {
  list-style: none;
  padding: 0;
  margin: 20px 0 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.flow-list li {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

.step-num {
  font-family: 'Playfair Display', serif;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: var(--accent);
  min-width: 28px;
  padding-top: 3px;
}

.flow-list div {
  font-size: 0.95rem;
  line-height: 1.75;
  color: var(--ink);
}

.flow-list strong { color: var(--ink); }

.tech-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-top: 16px;
}

@media (max-width: 520px) {
  .tech-grid { grid-template-columns: 1fr; }
}

.tech-card {
  background: var(--paper-warm);
  border: 1px solid var(--border);
  padding: 20px 24px;
  border-radius: 3px;
}

.tech-card h3 {
  font-family: 'Playfair Display', serif;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--accent);
  margin: 0 0 12px;
}

.tech-card ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.tech-card li {
  font-size: 0.88rem;
  line-height: 1.8;
  color: var(--ink-muted);
  padding-left: 14px;
  position: relative;
}

.tech-card li::before {
  content: '—';
  position: absolute;
  left: 0;
  color: var(--ink-faint);
}

.compare-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin: 20px 0 20px;
}

@media (max-width: 560px) {
  .compare-grid { grid-template-columns: 1fr; }
}

.compare-col {
  background: var(--paper-warm);
  border: 1px solid var(--border);
  border-radius: 3px;
  padding: 20px 22px;
}

.compare-heading {
  font-family: 'Playfair Display', serif;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  margin: 0 0 14px;
}

.where-gober  { color: var(--accent); }
.where-adblock { color: var(--ink-muted); }

.compare-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.compare-list li {
  font-size: 0.88rem;
  line-height: 1.7;
  color: var(--ink);
  padding-left: 16px;
  position: relative;
}

.compare-list li::before {
  content: '—';
  position: absolute;
  left: 0;
  color: var(--ink-faint);
}

.compare-summary {
  font-style: italic;
  color: var(--ink-muted);
  font-size: 0.93rem;
  line-height: 1.8;
  border-left: 2px solid var(--accent);
  padding-left: 16px;
  margin: 0;
}

.source-list {
  padding-left: 20px;
  margin: 12px 0 16px;
}

.source-list li {
  font-size: 0.95rem;
  line-height: 1.8;
  color: var(--ink);
  margin-bottom: 6px;
}

.github-link {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  font-family: 'Playfair Display', serif;
  font-size: 0.82rem;
  font-weight: 700;
  letter-spacing: 0.1em;
  text-transform: uppercase;
  color: var(--ink);
  text-decoration: none;
  border: 1px solid var(--ink);
  padding: 12px 20px;
  border-radius: 2px;
  margin-top: 8px;
  transition: background 0.2s, color 0.2s;
}

.github-link:hover {
  background: var(--ink);
  color: var(--paper);
}

.page-footer {
  border-top: 1px solid var(--border);
  text-align: center;
  padding: 28px 24px;
  color: var(--ink-faint);
}

.footer-ornament {
  display: block;
  font-size: 0.5rem;
  margin-bottom: 10px;
  color: var(--accent);
}

.footer-text {
  font-family: 'Lora', serif;
  font-style: italic;
  font-size: 0.8rem;
  margin: 0;
  letter-spacing: 0.03em;
}
</style>

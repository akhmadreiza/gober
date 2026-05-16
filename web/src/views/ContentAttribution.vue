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
        <p class="page-kicker">Legal &amp; Attribution</p>
        <h1 class="page-title">Content Attribution &amp; Publisher Notice</h1>
        <div class="ornament-row">
          <span class="orn-line"></span>
          <span class="orn-diamond">◆</span>
          <span class="orn-line"></span>
        </div>
        <p class="page-intro">
          Gober does not own, produce, or claim any rights to the news articles
          displayed on this website. All content belongs entirely to its original publishers.
        </p>
      </header>

      <section class="page-section">
        <h2 class="section-title">1. Ownership of Content</h2>
        <p>
          All news articles, headlines, photographs, images, videos, and other media
          content accessible through Gober are the exclusive intellectual property of
          their respective original publishers. Gober makes no claim — expressed or
          implied — to ownership of any content that originates from third-party news
          sources.
        </p>
        <p>
          The publishers currently served by Gober include, but are not limited to:
        </p>
        <ul class="publisher-list">
          <li>
            <strong>PT Agrakom / Detik Network</strong> — operator of
            <a href="https://www.detik.com" target="_blank" rel="noopener noreferrer">detik.com</a>
            and its affiliated sub-domains (news.detik.com, finance.detik.com,
            sport.detik.com, health.detik.com, inet.detik.com, travel.detik.com,
            food.detik.com, otomotif.detik.com, and others)
          </li>
          <li>
            <strong>PT Kompas Cyber Media</strong> — operator of
            <a href="https://www.kompas.com" target="_blank" rel="noopener noreferrer">kompas.com</a>
            and its affiliated sub-domains (nasional.kompas.com, regional.kompas.com,
            money.kompas.com, bola.kompas.com, health.kompas.com, and others)
          </li>
        </ul>
        <p>
          The full copyright, database rights, and any other intellectual property rights
          in the content remain with the respective publishers at all times.
        </p>
      </section>

      <section class="page-section">
        <h2 class="section-title">2. What Gober Does</h2>
        <p>
          Gober is a client-side reading tool. It operates in the following manner:
        </p>
        <ul class="policy-list">
          <li>
            Gober fetches publicly accessible HTML pages from the original publisher's
            website — the same request that any standard web browser would make.
          </li>
          <li>
            The HTML is parsed to extract the article's editorial content: title, author
            name, publication date, cover image, and body text.
          </li>
          <li>
            Non-editorial elements — including but not limited to advertisements, tracking
            scripts, analytics pixels, survey overlays, push-notification prompts, and
            third-party widgets — are removed before the content is presented to the reader.
          </li>
          <li>
            The original article URL and source attribution are preserved and displayed
            to the user. Every article page on Gober links directly back to the original
            publisher's version of the article.
          </li>
          <li>
            Gober does not cache content indefinitely. The in-memory cache has a
            maximum time-to-live of five (5) minutes, after which the content is
            re-fetched directly from the publisher's server.
          </li>
          <li>
            Gober does not store article content in a persistent database or archive.
            No long-term copy of any article is retained.
          </li>
        </ul>
      </section>

      <section class="page-section">
        <h2 class="section-title">3. What Gober Does Not Do</h2>
        <p>
          To be explicit about the scope and intent of this service:
        </p>
        <ul class="policy-list">
          <li>
            Gober does <strong>not</strong> reproduce or mirror articles in full on a
            separate domain for the purpose of replacing traffic to the original publisher.
          </li>
          <li>
            Gober does <strong>not</strong> monetise the content in any way. There are
            no advertisements, affiliate links, sponsored placements, or paid
            subscriptions associated with this service.
          </li>
          <li>
            Gober does <strong>not</strong> claim editorial responsibility for the
            accuracy, completeness, or timeliness of any article content. All editorial
            responsibility rests entirely with the original publisher.
          </li>
          <li>
            Gober does <strong>not</strong> index or republish content in a manner
            intended to compete with the publisher's own search engine visibility.
          </li>
          <li>
            Gober does <strong>not</strong> circumvent any technical protection measures,
            paywalls, or login requirements on the publisher's website.
          </li>
        </ul>
      </section>

      <section class="page-section">
        <h2 class="section-title">4. Purpose and Audience</h2>
        <p>
          Gober is a personal, non-commercial project built and maintained by an individual
          developer. It is intended for readers who prefer a distraction-free reading
          experience and is not affiliated with, endorsed by, or sponsored by any of the
          news publishers whose content it aggregates.
        </p>
        <p>
          The service operates in good faith under the principle that publicly accessible
          news articles may be fetched and presented to a reader in a different visual
          format, provided that: the content is not exploited commercially, attribution
          to the original publisher is maintained, and the reader is directed back to the
          original source.
        </p>
        <p>
          Gober respects the journalistic work of Indonesian reporters and editorial teams.
          Our goal is to remove the friction of reading — not to diminish the value of the
          work produced by these publications.
        </p>
      </section>

      <section class="page-section">
        <h2 class="section-title">5. Notice to Publishers</h2>
        <p>
          If you are a representative of a news organisation whose content appears on
          Gober and you have concerns about how your content is being presented, we
          welcome direct communication.
        </p>
        <p>
          We are willing to:
        </p>
        <ul class="policy-list">
          <li>Remove your publication from Gober entirely upon request</li>
          <li>Adjust how your content is attributed or linked</li>
          <li>Add a visible "Read on [Publisher]" call-to-action within every article</li>
          <li>Discuss any other accommodation that aligns with your content policy</li>
        </ul>
        <p>
          To make a takedown request or raise a concern, please open an issue on the
          project's GitHub repository or contact the developer directly through the
          repository's contact information.
        </p>
        <a
          href="https://github.com/akhmadreiza/gober/issues"
          target="_blank"
          rel="noopener noreferrer"
          class="contact-link"
        >
          <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z"/></svg>
          Open an Issue on GitHub
        </a>
      </section>

      <section class="page-section">
        <h2 class="section-title">6. Changes to This Notice</h2>
        <p>
          This notice may be updated from time to time to reflect changes in the
          scope of Gober or in response to feedback from publishers. The most current
          version is always available at this URL. Material changes will be noted in
          the project's GitHub commit history.
        </p>
        <p class="last-updated">
          Last updated: May 2026
        </p>
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
  name: 'ContentAttribution',
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
  font-size: clamp(1.6rem, 5vw, 2.6rem);
  font-weight: 900;
  letter-spacing: 0.02em;
  color: var(--ink);
  margin: 0 0 20px;
  line-height: 1.15;
}

.ornament-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 24px;
}
.orn-line  { width: 60px; height: 1px; background: var(--border); }
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
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--ink);
  margin: 0 0 16px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--border);
  letter-spacing: 0.02em;
}

p {
  font-size: 0.95rem;
  line-height: 1.85;
  color: var(--ink);
  margin: 0 0 14px;
}

.publisher-list,
.policy-list {
  padding-left: 0;
  margin: 12px 0 16px;
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.publisher-list li,
.policy-list li {
  font-size: 0.94rem;
  line-height: 1.8;
  color: var(--ink);
  padding-left: 20px;
  position: relative;
}

.publisher-list li::before,
.policy-list li::before {
  content: '—';
  position: absolute;
  left: 0;
  color: var(--accent);
  font-weight: 700;
}

.publisher-list a,
.policy-list a {
  color: var(--accent);
  text-decoration: none;
}
.publisher-list a:hover,
.policy-list a:hover {
  text-decoration: underline;
}

.last-updated {
  font-size: 0.82rem;
  color: var(--ink-faint);
  font-style: italic;
  margin-top: 8px;
}

.contact-link {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  font-family: 'Playfair Display', serif;
  font-size: 0.8rem;
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

.contact-link:hover {
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

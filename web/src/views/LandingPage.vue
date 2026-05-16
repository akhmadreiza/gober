<template>
  <div class="landing">

    <!-- Masthead -->
    <header class="masthead">
      <div class="masthead-inner">
        <div class="masthead-top-bar">
          <span class="masthead-date">{{ currentDate }}</span>
          <span class="masthead-tagline">Berita Bebas Iklan</span>
        </div>
        <div class="masthead-rule"></div>
        <h1 class="masthead-title">GOBER</h1>
        <p class="masthead-sub">Go Berita</p>
        <div class="masthead-rule"></div>
      </div>
    </header>

    <!-- Source tabs -->
    <nav class="source-nav">
      <div class="source-nav-inner">

        <!-- About dropdown (far left) -->
        <div class="about-wrap">
          <button class="source-tab about-trigger">
            About
            <svg class="about-chevron" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"/></svg>
          </button>
          <div class="about-dropdown">
            <router-link to="/about/how-it-works" class="dropdown-link">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/></svg>
              How It Works
            </router-link>
            <div class="dropdown-separator"></div>
            <router-link to="/about/content-attribution" class="dropdown-link">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M14.83 14.83a4 4 0 1 1 0-5.66"/></svg>
              Content Attribution
            </router-link>
          </div>
        </div>

        <!-- Divider between About and source tabs -->
        <div class="nav-sep"></div>

        <button
          v-for="site in websites"
          :key="site.name"
          :class="['source-tab', { active: activeSource === site.name }]"
          @click="setActiveSource(site.name)"
        >
          {{ site.displayName }}
        </button>
        <div class="tab-rule"></div>
      </div>
    </nav>

    <!-- Main content -->
    <main class="content-area">

      <!-- Skeleton loading -->
      <div v-if="isLoading" class="skeleton-layout">
        <div class="sk-hero-wrap">
          <div class="sk sk-hero-img"></div>
          <div class="sk sk-h1"></div>
          <div class="sk sk-h1 w70"></div>
          <div class="sk sk-date"></div>
        </div>
        <div class="sk-grid">
          <div v-for="n in 6" :key="n" class="sk-card-wrap">
            <div class="sk sk-card-img"></div>
            <div class="sk sk-h2"></div>
            <div class="sk sk-h2 w60"></div>
            <div class="sk sk-date"></div>
          </div>
        </div>
      </div>

      <!-- Articles -->
      <div v-else-if="visibleArticles.length">

        <!-- Hero: first article -->
        <a :href="articleHref(visibleArticles[0])" class="hero-card">
          <div class="hero-img-wrap">
            <img
              v-if="visibleArticles[0].img_url"
              :src="visibleArticles[0].img_url"
              :alt="visibleArticles[0].title"
              class="hero-img"
              @error="fallbackImg"
            />
            <div v-else class="hero-img-empty"></div>
            <span class="hero-badge">{{ activeSource }}</span>
          </div>
          <div class="hero-body">
            <h2 class="hero-title">{{ visibleArticles[0].title }}</h2>
            <time class="hero-time" v-if="visibleArticles[0].timestamp">{{ visibleArticles[0].timestamp }}</time>
          </div>
        </a>

        <!-- Divider between hero and grid -->
        <div class="section-divider">
          <span class="divider-label">Berita Lainnya</span>
          <span class="divider-line"></span>
        </div>

        <!-- Card grid -->
        <div class="cards-grid">
          <a
            v-for="(article, i) in visibleArticles.slice(1)"
            :key="article.source_url || i"
            :href="articleHref(article)"
            class="article-card"
          >
            <div class="card-img-wrap">
              <img
                v-if="article.img_url"
                :src="article.img_url"
                :alt="article.title"
                class="card-img"
                @error="fallbackImg"
              />
              <div v-else class="card-img-empty"></div>
            </div>
            <div class="card-body">
              <h3 class="card-title">{{ article.title }}</h3>
              <time class="card-time" v-if="article.timestamp">{{ article.timestamp }}</time>
            </div>
          </a>
        </div>

        <!-- Load more -->
        <div class="load-more-wrap" v-if="currentWebsite.visibleCount < currentWebsite.articles.length">
          <button class="load-more-btn" @click="loadMore">
            <span class="load-more-line"></span>
            <span class="load-more-text">
              Lihat Semua
              <em>({{ currentWebsite.articles.length - currentWebsite.visibleCount }} berita lagi)</em>
            </span>
            <span class="load-more-line"></span>
          </button>
        </div>
      </div>

      <div v-else class="empty-state">
        <p>Tidak ada artikel tersedia.</p>
      </div>
    </main>

    <footer class="site-footer">
      <span class="footer-ornament">◆</span>
      <p class="footer-text">Gober — Go Berita. Baca tanpa gangguan iklan.</p>
    </footer>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'LandingPage',
  data() {
    return {
      websites: [
        { name: 'detik',  displayName: 'Detik',  articles: [], visibleCount: 10 },
        { name: 'kompas', displayName: 'Kompas', articles: [], visibleCount: 10 },
      ],
      activeSource: this.$route.query.source || 'detik',
      isLoading: true,
    };
  },
  computed: {
    currentWebsite() {
      return this.websites.find(s => s.name === this.activeSource) || {};
    },
    visibleArticles() {
      const site = this.currentWebsite;
      return site.articles ? site.articles.slice(0, site.visibleCount) : [];
    },
    currentDate() {
      return new Date().toLocaleDateString('id-ID', {
        weekday: 'long', year: 'numeric', month: 'long', day: 'numeric',
      });
    },
  },
  methods: {
    articleHref(article) {
      return `/detail?source=${this.activeSource}&detailUrl=${encodeURIComponent(article.source_url)}`;
    },
    setActiveSource(source) {
      this.activeSource = source;
      this.$router.replace({ query: { source } });
    },
    loadMore() {
      const site = this.currentWebsite;
      if (site) site.visibleCount = site.articles.length;
    },
    fallbackImg(event) {
      const img = event.target;
      img.onerror = null;
      img.src = img.src
        .replace(/w=\d+/, 'w=220')
        .replace(/\/\d+x\d+\//, '/230x153/');
    },
    async fetchArticles() {
      try {
        await Promise.all(
          this.websites.map(site =>
            axios.get(`/articles/popular?source=${site.name}`).then(res => {
              if (res.data.status === 'Success' && Array.isArray(res.data.articles)) {
                site.articles = res.data.articles;
              }
            })
          )
        );
      } catch (e) {
        console.error('Error fetching articles:', e);
      } finally {
        this.isLoading = false;
      }
    },
  },
  async created() {
    await this.fetchArticles();
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,600;0,700;0,900;1,400&family=Lora:ital,wght@0,400;0,500;1,400&display=swap');

:root {
  --ink:        #1C1917;
  --ink-muted:  #78716C;
  --ink-faint:  #A8A29E;
  --paper:      #FAFAF7;
  --paper-warm: #F5F0E8;
  --accent:     #B45309;
  --border:     #E7E5E4;
  --nav-bg:     rgba(250, 250, 247, 0.95);
}
</style>

<style scoped>
/* ── Base ───────────────────────────────────────── */
.landing {
  background: var(--paper);
  min-height: 100vh;
  font-family: 'Lora', Georgia, serif;
  color: var(--ink);
}

/* ── Masthead ───────────────────────────────────── */
.masthead {
  padding: 28px 24px 0;
  text-align: center;
  animation: fadeDown 0.6s ease both;
}

@keyframes fadeDown {
  from { opacity: 0; transform: translateY(-12px); }
  to   { opacity: 1; transform: translateY(0); }
}

.masthead-inner {
  max-width: 860px;
  margin: 0 auto;
}

.masthead-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.72rem;
  font-style: italic;
  color: var(--ink-muted);
  letter-spacing: 0.04em;
  margin-bottom: 10px;
  font-family: 'Lora', serif;
}

.masthead-rule {
  height: 1px;
  background: var(--ink);
  margin: 10px 0;
}

.masthead-rule:last-child {
  margin-bottom: 0;
}

.masthead-title {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: clamp(3.2rem, 10vw, 6.5rem);
  font-weight: 900;
  letter-spacing: 0.18em;
  color: var(--ink);
  margin: 8px 0 4px;
  line-height: 1;
}

.masthead-sub {
  font-family: 'Lora', serif;
  font-style: italic;
  font-size: 0.9rem;
  color: var(--ink-muted);
  margin: 0 0 10px;
  letter-spacing: 0.12em;
}

/* ── Source tabs ────────────────────────────────── */
.source-nav {
  position: sticky;
  top: 0;
  z-index: 50;
  background: var(--nav-bg);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--border);
}

.source-nav-inner {
  max-width: 860px;
  margin: 0 auto;
  padding: 0 24px;
  display: flex;
  align-items: stretch;
  gap: 0;
  position: relative;
}

.source-tab {
  font-family: 'Playfair Display', serif;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: var(--ink-muted);
  background: none;
  border: none;
  padding: 14px 20px;
  cursor: pointer;
  position: relative;
  text-transform: uppercase;
  transition: color 0.2s;
}

.source-tab::after {
  content: '';
  position: absolute;
  bottom: 0; left: 0; right: 0;
  height: 2px;
  background: var(--accent);
  transform: scaleX(0);
  transition: transform 0.25s ease;
}

.source-tab:hover {
  color: var(--ink);
}

.source-tab.active {
  color: var(--ink);
}

.source-tab.active::after {
  transform: scaleX(1);
}

.tab-rule {
  display: none;
}

/* ── Content area ───────────────────────────────── */
.content-area {
  max-width: 860px;
  margin: 0 auto;
  padding: 36px 24px 60px;
}

/* ── Skeleton ───────────────────────────────────── */
.skeleton-layout { }

.sk-hero-wrap {
  margin-bottom: 40px;
}

.sk-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
}

@media (min-width: 640px) {
  .sk-grid { grid-template-columns: repeat(3, 1fr); }
}

.sk {
  background: linear-gradient(90deg, var(--border) 25%, var(--paper-warm) 50%, var(--border) 75%);
  background-size: 200% 100%;
  animation: shimmer 1.6s infinite;
  border-radius: 3px;
  margin-bottom: 10px;
  height: 13px;
}

@keyframes shimmer {
  0%   { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

.sk-hero-img { height: 380px; border-radius: 4px; margin-bottom: 16px; }
.sk-h1       { height: 28px; width: 100%; }
.sk-h1.w70   { width: 70%; }
.sk-h2       { height: 15px; width: 100%; }
.sk-h2.w60   { width: 60%; }
.sk-date     { height: 10px; width: 120px; margin-top: 4px; }
.sk-card-img { height: 160px; border-radius: 4px; margin-bottom: 12px; }

/* ── Hero card ──────────────────────────────────── */
.hero-card {
  display: block;
  text-decoration: none;
  color: inherit;
  margin-bottom: 36px;
  animation: fadeUp 0.5s ease both;
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(16px); }
  to   { opacity: 1; transform: translateY(0); }
}

.hero-img-wrap {
  position: relative;
  overflow: hidden;
  border-radius: 4px;
  margin-bottom: 16px;
}

.hero-img {
  width: 100%;
  height: 380px;
  object-fit: cover;
  display: block;
  transition: transform 0.5s ease;
}

.hero-card:hover .hero-img {
  transform: scale(1.02);
}

.hero-img-empty {
  width: 100%;
  height: 380px;
  background: var(--paper-warm);
}

.hero-badge {
  position: absolute;
  bottom: 12px; left: 12px;
  font-family: 'Playfair Display', serif;
  font-size: 0.65rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  background: var(--accent);
  color: #fff;
  padding: 4px 10px;
  border-radius: 2px;
}

.hero-body { }

.hero-title {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: clamp(1.5rem, 4vw, 2rem);
  font-weight: 700;
  line-height: 1.25;
  color: var(--ink);
  margin: 0 0 10px;
  transition: color 0.2s;
}

.hero-card:hover .hero-title {
  color: var(--accent);
}

.hero-time {
  font-family: 'Lora', serif;
  font-style: italic;
  font-size: 0.82rem;
  color: var(--ink-muted);
}

/* ── Section divider ────────────────────────────── */
.section-divider {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 28px;
}

.divider-label {
  font-family: 'Playfair Display', serif;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.16em;
  text-transform: uppercase;
  color: var(--ink-muted);
  white-space: nowrap;
}

.divider-line {
  flex: 1;
  height: 1px;
  background: var(--border);
}

/* ── Cards grid ─────────────────────────────────── */
.cards-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 28px 20px;
  margin-bottom: 40px;
}

@media (min-width: 640px) {
  .cards-grid { grid-template-columns: repeat(3, 1fr); }
}

.article-card {
  display: block;
  text-decoration: none;
  color: inherit;
  animation: fadeUp 0.5s ease both;
}

.article-card:nth-child(1) { animation-delay: 0.04s; }
.article-card:nth-child(2) { animation-delay: 0.08s; }
.article-card:nth-child(3) { animation-delay: 0.12s; }
.article-card:nth-child(4) { animation-delay: 0.16s; }
.article-card:nth-child(5) { animation-delay: 0.20s; }
.article-card:nth-child(6) { animation-delay: 0.24s; }
.article-card:nth-child(7) { animation-delay: 0.28s; }
.article-card:nth-child(8) { animation-delay: 0.32s; }

.card-img-wrap {
  overflow: hidden;
  border-radius: 3px;
  margin-bottom: 12px;
  background: var(--paper-warm);
}

.card-img {
  width: 100%;
  height: 160px;
  object-fit: cover;
  display: block;
  transition: transform 0.4s ease;
}

.article-card:hover .card-img {
  transform: scale(1.04);
}

.card-img-empty {
  width: 100%;
  height: 160px;
  background: var(--paper-warm);
}

.card-body { }

.card-title {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: 0.95rem;
  font-weight: 700;
  line-height: 1.4;
  color: var(--ink);
  margin: 0 0 8px;
  transition: color 0.2s;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-card:hover .card-title {
  color: var(--accent);
}

.card-time {
  font-family: 'Lora', serif;
  font-style: italic;
  font-size: 0.75rem;
  color: var(--ink-faint);
}

/* ── Load more ──────────────────────────────────── */
.load-more-wrap {
  text-align: center;
  margin-top: 8px;
}

.load-more-btn {
  display: inline-flex;
  align-items: center;
  gap: 16px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 12px 0;
  font-family: 'Playfair Display', serif;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--ink-muted);
  transition: color 0.2s;
}

.load-more-btn:hover {
  color: var(--accent);
}

.load-more-line {
  display: block;
  width: 48px;
  height: 1px;
  background: currentColor;
  transition: width 0.3s ease;
}

.load-more-btn:hover .load-more-line {
  width: 64px;
}

.load-more-text em {
  font-weight: 400;
  font-style: italic;
  letter-spacing: 0;
  text-transform: none;
  color: var(--ink-faint);
  margin-left: 4px;
  font-size: 0.85em;
}

/* ── Empty state ────────────────────────────────── */
.empty-state {
  text-align: center;
  padding: 60px 0;
  font-style: italic;
  color: var(--ink-muted);
}

/* ── Footer ─────────────────────────────────────── */
.site-footer {
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

/* ── About dropdown ─────────────────────────────── */
.about-wrap {
  position: relative;
}

.about-trigger {
  display: inline-flex;
  align-items: center;
  gap: 5px;
}

.about-chevron {
  transition: transform 0.2s ease;
  opacity: 0.6;
}

.about-wrap:hover .about-chevron {
  transform: rotate(180deg);
  opacity: 1;
}

.about-dropdown {
  position: absolute;
  top: calc(100% + 1px);
  left: 0;
  background: var(--paper);
  border: 1px solid var(--border);
  border-top: 2px solid var(--accent);
  min-width: 192px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.08);
  opacity: 0;
  visibility: hidden;
  transform: translateY(-6px);
  transition: opacity 0.18s ease, transform 0.18s ease, visibility 0.18s;
  z-index: 200;
  border-radius: 0 0 3px 3px;
}

.about-wrap:hover .about-dropdown {
  opacity: 1;
  visibility: visible;
  transform: translateY(0);
}

.dropdown-link {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 13px 16px;
  font-family: 'Playfair Display', serif;
  font-size: 0.74rem;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--ink-muted);
  text-decoration: none;
  transition: color 0.15s, background 0.15s;
  white-space: nowrap;
}

.dropdown-link svg {
  flex-shrink: 0;
  opacity: 0.5;
  transition: opacity 0.15s;
}

.dropdown-link:hover {
  color: var(--ink);
  background: var(--paper-warm);
}

.dropdown-link:hover svg {
  opacity: 1;
}

.dropdown-separator {
  height: 1px;
  background: var(--border);
  margin: 0 16px;
}

.nav-sep {
  width: 1px;
  background: var(--border);
  margin: 10px 4px;
  align-self: stretch;
}

/* ── Mobile ─────────────────────────────────────── */
@media (max-width: 480px) {
  .masthead-top-bar {
    flex-direction: column;
    gap: 2px;
    text-align: center;
  }

  .hero-img, .hero-img-empty {
    height: 240px;
  }

  .card-img, .card-img-empty {
    height: 130px;
  }

  .cards-grid {
    gap: 20px 12px;
  }
}
</style>

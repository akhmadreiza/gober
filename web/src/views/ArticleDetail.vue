<template>
  <div class="reader-page" :class="{ loaded: isLoaded }">

    <div class="progress-bar" :style="{ width: readingProgress + '%' }"></div>

    <nav class="reader-nav" :class="{ visible: navVisible }">
      <a href="/" class="nav-back">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 12H5M12 5l-7 7 7 7"/></svg>
        Gober
      </a>
      <span class="nav-title" v-if="article">{{ article.title }}</span>
    </nav>

    <main class="reader-container">

      <div v-if="isLoading" class="skeleton-wrapper">
        <div class="sk sk-back"></div>
        <div class="sk sk-title-1"></div>
        <div class="sk sk-title-2"></div>
        <div class="sk sk-title-3"></div>
        <div class="sk sk-meta"></div>
        <div class="sk sk-image"></div>
        <div class="sk sk-p"></div>
        <div class="sk sk-p sk-short"></div>
        <div class="sk sk-p"></div>
        <div class="sk sk-p sk-medium"></div>
        <div class="sk sk-p"></div>
      </div>

      <article v-else-if="article" class="article">
        <a href="/" class="article-back">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 12H5M12 5l-7 7 7 7"/></svg>
          Kembali ke Beranda
        </a>

        <header class="article-header">
          <h1 class="article-title">{{ article.title }}</h1>

          <div class="article-ornament">
            <span class="ornament-line"></span>
            <span class="ornament-diamond">◆</span>
            <span class="ornament-line"></span>
          </div>

          <div class="article-meta">
            <span class="meta-author" v-if="article.author">{{ article.author }}</span>
            <span class="meta-dot" v-if="article.author && article.timestamp"></span>
            <time class="meta-date" v-if="article.timestamp">{{ article.timestamp }}</time>
          </div>
        </header>

        <figure class="article-hero" v-if="article.img_url">
          <img :src="article.img_url" :alt="article.title" class="hero-img" @error="fallbackImg" />
        </figure>

        <div class="article-body" v-html="article.content"></div>

        <footer class="article-footer">
          <div class="footer-rule">
            <span class="ornament-diamond small">◆</span>
          </div>
          <a href="/" class="footer-back">Baca berita lainnya di Gober →</a>
        </footer>
      </article>

      <div v-else class="error-state">
        <p class="error-msg">Artikel tidak dapat dimuat.</p>
        <a href="/" class="error-link">← Kembali ke Beranda</a>
      </div>

    </main>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ArticleDetail',
  data() {
    return {
      article: null,
      isLoading: true,
      isLoaded: false,
      readingProgress: 0,
      navVisible: false,
    };
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll, { passive: true });
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    fallbackImg(event) {
      const img = event.target;
      img.onerror = null;
      img.src = img.src
        .replace(/w=\d+/, 'w=220')
        .replace(/\/\d+x\d+\//, '/230x153/');
    },
    handleScroll() {
      const scrollTop = window.scrollY;
      const docHeight = document.documentElement.scrollHeight - window.innerHeight;
      this.readingProgress = docHeight > 0 ? (scrollTop / docHeight) * 100 : 0;
      this.navVisible = scrollTop > 240;
    },
  },
  async created() {
    const detailUrl = this.$route.query.detailUrl;
    const sourceWebsite = this.$route.query.source;
    if (!detailUrl) {
      this.isLoading = false;
      return;
    }
    try {
      const response = await axios.get(`/article?source=${sourceWebsite}&detailUrl=${detailUrl}`);
      if (response.data.status === 'Success' && response.data.articles.length) {
        this.article = response.data.articles[0];
        this.$nextTick(() => setTimeout(() => { this.isLoaded = true; }, 60));
      }
    } catch (error) {
      console.error('Error fetching article details:', error);
    } finally {
      this.isLoading = false;
    }
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,600;0,700;1,400&family=Lora:ital,wght@0,400;0,500;1,400&display=swap');

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
/* ── Page shell ─────────────────────────────────────── */
.reader-page {
  background: var(--paper);
  min-height: 100vh;
  font-family: 'Lora', Georgia, serif;
  color: var(--ink);
}

/* ── Reading progress ───────────────────────────────── */
.progress-bar {
  position: fixed;
  top: 0; left: 0;
  height: 3px;
  background: var(--accent);
  z-index: 200;
  transition: width 0.1s linear;
  border-radius: 0 2px 2px 0;
}

/* ── Sticky nav ─────────────────────────────────────── */
.reader-nav {
  position: fixed;
  top: 0; left: 0; right: 0;
  height: 52px;
  background: var(--nav-bg);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--border);
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 0 24px;
  z-index: 100;
  transform: translateY(-100%);
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.reader-nav.visible {
  transform: translateY(0);
}

.nav-back {
  display: flex;
  align-items: center;
  gap: 6px;
  font-family: 'Lora', serif;
  font-size: 0.85rem;
  color: var(--accent);
  text-decoration: none;
  white-space: nowrap;
  flex-shrink: 0;
}

.nav-back:hover {
  color: var(--ink);
}

.nav-title {
  font-size: 0.8rem;
  color: var(--ink-muted);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* ── Main container ─────────────────────────────────── */
.reader-container {
  max-width: 720px;
  margin: 0 auto;
  padding: 48px 24px 80px;
}

/* ── Skeleton loader ────────────────────────────────── */
.skeleton-wrapper {
  padding-top: 12px;
}

.sk {
  background: linear-gradient(90deg, var(--border) 25%, var(--paper-warm) 50%, var(--border) 75%);
  background-size: 200% 100%;
  animation: shimmer 1.6s infinite;
  border-radius: 4px;
  margin-bottom: 14px;
  height: 14px;
}

@keyframes shimmer {
  0%   { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

.sk-back    { width: 120px; height: 12px; margin-bottom: 32px; }
.sk-title-1 { width: 100%; height: 36px; border-radius: 2px; }
.sk-title-2 { width: 88%;  height: 36px; border-radius: 2px; }
.sk-title-3 { width: 60%;  height: 36px; border-radius: 2px; margin-bottom: 28px; }
.sk-meta    { width: 200px; height: 12px; margin: 0 auto 36px; }
.sk-image   { width: 100%; height: 360px; border-radius: 6px; margin-bottom: 36px; }
.sk-p       { width: 100%; }
.sk-short   { width: 72%; }
.sk-medium  { width: 85%; }

/* ── Back link ──────────────────────────────────────── */
.article-back {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 0.82rem;
  color: var(--accent);
  text-decoration: none;
  margin-bottom: 40px;
  letter-spacing: 0.03em;
  transition: gap 0.2s ease;
}

.article-back:hover {
  gap: 10px;
  color: var(--ink);
}

/* ── Article header ─────────────────────────────────── */
.article-header {
  margin-bottom: 36px;
  opacity: 0;
  transform: translateY(16px);
  transition: opacity 0.6s ease, transform 0.6s ease;
}

.loaded .article-header {
  opacity: 1;
  transform: translateY(0);
}

.article-title {
  font-family: 'Playfair Display', Georgia, serif;
  font-size: clamp(1.9rem, 5vw, 2.8rem);
  font-weight: 700;
  line-height: 1.22;
  color: var(--ink);
  margin: 0 0 28px;
  letter-spacing: -0.01em;
}

/* ── Ornament divider ───────────────────────────────── */
.article-ornament {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 20px;
}

.ornament-line {
  flex: 1;
  height: 1px;
  background: var(--border);
}

.ornament-diamond {
  font-size: 0.55rem;
  color: var(--accent);
  line-height: 1;
}

/* ── Metadata ───────────────────────────────────────── */
.article-meta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  font-size: 0.83rem;
  color: var(--ink-muted);
  font-family: 'Lora', serif;
  font-style: italic;
}

.meta-dot {
  width: 3px;
  height: 3px;
  border-radius: 50%;
  background: var(--ink-faint);
  flex-shrink: 0;
}

/* ── Hero image ─────────────────────────────────────── */
.article-hero {
  margin: 0 0 40px;
  opacity: 0;
  transform: translateY(12px);
  transition: opacity 0.7s ease 0.15s, transform 0.7s ease 0.15s;
}

.loaded .article-hero {
  opacity: 1;
  transform: translateY(0);
}

.hero-img {
  width: 100%;
  height: auto;
  border-radius: 4px;
  display: block;
  max-height: 480px;
  object-fit: cover;
}

/* ── Article body ───────────────────────────────────── */
.article-body {
  opacity: 0;
  transform: translateY(12px);
  transition: opacity 0.7s ease 0.25s, transform 0.7s ease 0.25s;
}

.loaded .article-body {
  opacity: 1;
  transform: translateY(0);
}

.article-body :deep(p) {
  font-family: 'Lora', Georgia, serif;
  font-size: 1.08rem;
  line-height: 1.85;
  color: var(--ink);
  margin: 0 0 1.4em;
}

.article-body :deep(h2),
.article-body :deep(h3) {
  font-family: 'Playfair Display', Georgia, serif;
  font-weight: 700;
  color: var(--ink);
  margin: 2em 0 0.6em;
  line-height: 1.3;
}

.article-body :deep(h2) { font-size: 1.5rem; }
.article-body :deep(h3) { font-size: 1.2rem; }

.article-body :deep(strong) {
  font-weight: 600;
  color: var(--ink);
}

.article-body :deep(em) {
  font-style: italic;
}

.article-body :deep(a) {
  color: var(--accent);
  text-decoration: underline;
  text-decoration-color: transparent;
  text-underline-offset: 3px;
  transition: text-decoration-color 0.2s;
}

.article-body :deep(a:hover) {
  text-decoration-color: var(--accent);
}

.article-body :deep(img) {
  max-width: 100% !important;
  height: auto !important;
  display: block;
  margin: 1.8em auto;
  border-radius: 4px;
}

.article-body :deep(blockquote) {
  border-left: 3px solid var(--accent);
  margin: 2em 0;
  padding: 0.5em 0 0.5em 1.5em;
  font-style: italic;
  color: var(--ink-muted);
}

.article-body :deep(ul),
.article-body :deep(ol) {
  padding-left: 1.5em;
  margin-bottom: 1.4em;
}

.article-body :deep(li) {
  font-family: 'Lora', serif;
  font-size: 1.08rem;
  line-height: 1.8;
  margin-bottom: 0.4em;
}

/* ── Footer ─────────────────────────────────────────── */
.article-footer {
  margin-top: 64px;
  text-align: center;
}

.footer-rule {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 28px;
  gap: 12px;
}

.footer-rule::before,
.footer-rule::after {
  content: '';
  flex: 1;
  max-width: 80px;
  height: 1px;
  background: var(--border);
}

.ornament-diamond.small {
  font-size: 0.45rem;
  color: var(--ink-faint);
}

.footer-back {
  font-family: 'Lora', serif;
  font-size: 0.9rem;
  color: var(--accent);
  text-decoration: none;
  letter-spacing: 0.02em;
  transition: letter-spacing 0.2s ease;
}

.footer-back:hover {
  letter-spacing: 0.05em;
  color: var(--ink);
}

/* ── Error state ─────────────────────────────────────── */
.error-state {
  text-align: center;
  padding: 80px 0;
}

.error-msg {
  font-family: 'Playfair Display', serif;
  font-size: 1.2rem;
  color: var(--ink-muted);
  margin-bottom: 20px;
}

.error-link {
  font-size: 0.9rem;
  color: var(--accent);
  text-decoration: none;
}

/* ── Mobile ─────────────────────────────────────────── */
@media (max-width: 600px) {
  .reader-container {
    padding: 32px 20px 60px;
  }

  .article-title {
    font-size: 1.75rem;
  }

  .article-body :deep(p),
  .article-body :deep(li) {
    font-size: 1rem;
    line-height: 1.78;
  }

  .sk-image {
    height: 220px;
  }
}
</style>

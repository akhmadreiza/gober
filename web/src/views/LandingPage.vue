<template>
  <div class="landing-page">
    <header class="header">
      <div class="header-container">
        <h1 class="header-title">GOBER - Go Berita</h1>
        <button class="menu-toggle" @click="toggleMenu" aria-label="Toggle Menu">
          ☰
        </button>
      </div>
      <nav class="menu" :class="{ open: isMenuOpen }">
        <button
          v-for="website in websites"
          :key="website.name"
          :class="['menu-item', { active: activeSource === website.name }]"
          @click="setActiveSource(website.name)"
        >
          {{ website.displayName }}
        </button>
      </nav>
    </header>


    <div v-if="isLoading" class="loading-message">Loading articles...</div>

    <div v-else>
      <div class="articles-container">
        <div
          class="article-card"
          v-for="(article, index) in currentWebsite.articles.slice(0, currentWebsite.visibleCount)"
          :key="article.source_url || index"
        >
          <div class="article-image-container">
            <img
              v-if="article.img_url"
              :src="article.img_url"
              alt="Article Image"
              class="article-image"
            />
          </div>
          <h2 class="article-title">
            <a
              :href="'/detail' + '?source=' + activeSource + '&detailUrl=' + encodeURIComponent(article.source_url)"
              class="read-link"
            >
            {{ article.title }}
            </a>
          </h2>
          <p class="article-date">{{ article.timestamp || '' }}</p>
          <p class="article-description">
            Source: <a
              :href="article.source_url"
              target="_blank"
              class="read-origin-link"
            >
              {{ article.description || 'Original Link' }}
            </a>
          </p>
        </div>
      </div>
      <button
        v-if="currentWebsite.visibleCount < currentWebsite.articles.length"
        @click="loadMore"
        class="load-more-button"
      >
        View All
      </button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "LandingPage",
  data() {
    return {
      websites: [
        { name: "detik", displayName: "Detik.com", articles: [], visibleCount: 8 },
        { name: "kompas", displayName: "Kompas.com", articles: [], visibleCount: 8 },
      ],
      activeSource: "detik", // Default active source
      isLoading: true,
      isMenuOpen: false,
    };
  },
  computed: {
    currentWebsite() {
      return this.websites.find((site) => site.name === this.activeSource) || {};
    },
  },
  methods: {
    async fetchArticles() {
      try {
        const requests = this.websites.map((site) =>
          axios
            .get(`/articles/popular?source=${site.name}`)
            .then((response) => {
              if (
                response.data.status === "Success" &&
                Array.isArray(response.data.articles)
              ) {
                site.articles = response.data.articles;
              } else {
                console.error(`Error fetching articles for ${site.name}`);
              }
            })
        );
        await Promise.all(requests);
      } catch (error) {
        console.error("Error fetching articles:", error);
      } finally {
        this.isLoading = false;
      }
    },
    toggleMenu() {
      this.isMenuOpen = !this.isMenuOpen; // Toggle menu visibility
    },
    setActiveSource(source) {
      this.activeSource = source;
    },
    loadMore() {
      const website = this.currentWebsite;
      if (website) {
        website.visibleCount = website.articles.length;
      }
    },
  },
  async created() {
    await this.fetchArticles();
  },
  toggleMenu() {
    this.isMenuOpen = !this.isMenuOpen;
  },
  setActiveSource(source) {
    this.activeSource = source;
    this.isMenuOpen = false; // Close menu on selection
  },
};
</script>

<style scoped>
/* Global Styles */

.landing-page {
  font-family: 'Arial', sans-serif;
  background-color: #fafafa; /* Light gray background */
  color: #333333; /* Dark gray text */
  text-align: center;
}

.page-title {
  font-size: 2.5rem;
  margin-bottom: 20px;
  color: #2c3e50; /* Dark blue for the title */
}

.articles-container {
  display: grid;
  grid-template-columns: repeat(1, 1fr); /* Default: 1 column for small screens */
  gap: 20px;
  padding: 0;
  max-width: 1200px; /* Maximum width to prevent stretching */
  margin: 0 auto; /* Center the grid container horizontally */
}

@media (min-width: 768px) {
  .articles-container {
    grid-template-columns: repeat(3, 1fr); /* 2 columns for medium screens */
  }
}

@media (min-width: 1024px) {
  .articles-container {
    grid-template-columns: repeat(4, 1fr); /* 3 columns for large screens */
  }
}

/* Article Card Styles */
.article-card {
  background-color: #ffffff; /* White background for cards */
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border-left: 5px solid #ecf0f1; /* Light gray border left */
  max-width: 350px; /* Prevent cards from stretching too much */
}

.article-title {
  font-size: 1rem;
  font-weight: bold;
  color: #2c3e50; /* Dark blue for titles */
  margin-bottom: 10px;
  text-align: left;
}

.article-date {
  font-size: 0.8rem;
  color: #555;
  margin: 10px 0px 10px 0px;
  text-align: left;
}

.article-description {
  font-size: 0.5rem;
  color: #555;
  margin: 10px 0px 10px 0px;
  text-align: left;
}

.article-links {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.read-origin-link {
  color: #3498db;
}

.read-more-link,
.read-link {
  color: #3498db;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.read-origin-link:hover,
.read-more-link:hover,
.read-link:hover {
  color: #2c3e50; /* Blue text on hover */
}

.no-url {
  font-size: 1rem;
  color: #e74c3c; /* Red color for error message */
  text-align: left;
}

.loading-message {
  font-size: 1.2rem;
  color: #888;
}

.article-image-container {
  margin-bottom: 15px;
  overflow: hidden;
  border-radius: 8px;
}

.article-image {
  width: 100%;
  height: auto;
  object-fit: cover;
  border-radius: 8px;
}

.website-title {
  font-size: 2rem;
  margin: 20px 0;
  color: #2c3e50; /* Dark blue for website titles */
}

.load-more-button {
  display: inline-block;
  margin: 20px auto;
  padding: 10px 20px;
  font-size: 1rem;
  color: white;
  background-color: #3498db;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.load-more-button:hover {
  background-color: #2980b9; /* Darker blue */
}

/* Header Styles */
.header {
  display: flex;
  flex-direction: column;
  background-color: #2c3e50;
  color: white;
  padding: 20px 20px;
  margin-bottom: 25px;
}

.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title {
  font-size: 1.8rem;
  font-weight: bold;
  margin: 0;
}

/* Menu Styles */
.menu {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  margin-top: 10px;
  overflow: hidden;
  transition: max-height 0.3s ease;
}

.menu.open {
  max-height: 200px; /* Adjust based on expected number of items */
}

.menu-item {
  background: none;
  color: white;
  font-size: 1rem;
  padding: 10px 15px;
  border: 1px solid transparent;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s, border 0.3s;
}

.menu-item:hover {
  background-color: #34495e;
}

.menu-item.active {
  background-color: #3498db;
  border-color: white;
}

/* Menu Toggle Button */
.menu-toggle {
  display: none; /* Hidden on larger screens */
  font-size: 1.5rem;
  color: white;
  background: none;
  border: none;
  cursor: pointer;
}

.menu-toggle:focus {
  outline: none;
}

.loading-message {
  font-size: 1.2rem;
  color: #888;
  margin: 20px 0;
  text-align: center;
}

@media (max-width: 480px) {
  .header-title {
    font-size: 1.5rem;
  }

  .menu-item {
    font-size: 0.9rem;
    padding: 10px;
  }
}

@media (max-width: 768px) {
  .page-title {
    font-size: 2rem;
  }

  .article-card {
    padding: 15px;
  }
  
  .header-container {
    flex-direction: row;
  }

  .menu {
    flex-direction: column;
    max-height: 0; /* Collapsed by default */
  }

  .menu-toggle {
    display: block; /* Show menu toggle on smaller screens */
  }

  .menu.open {
    max-height: 300px; /* Adjust for the expanded menu */
  }

  .menu-item {
    width: 100%; /* Full width for menu items */
    text-align: left;
    padding: 15px;
  }
}
</style>

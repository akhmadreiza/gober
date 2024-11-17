<template>
  <div class="landing-page">
    <h1 class="page-title">Popular Articles</h1>
    <div class="articles-container" v-if="articles.length">
      <div class="article-card" v-for="article in articles" :key="article.source_url">
        <div class="article-image-container">
          <!-- Add the image dynamically from img_url -->
          <img v-if="article.img_url" :src="article.img_url" alt="Article Image" class="article-image" />
        </div>
        <h2 class="article-title">{{ article.title }}</h2>
        <p class="article-description">{{ article.description || 'No description available.' }}</p>
        <div class="article-links">
          <!-- Link to detail page -->
          <a :href="'/detail?detailUrl=' + encodeURIComponent(article.source_url)" class="read-link">Read</a>
          <!-- Link to external source -->
          <a :href="article.source_url" target="_blank" class="read-more-link">Read from Source</a>
        </div>
      </div>
    </div>
    <p v-else class="loading-message">Loading articles...</p>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "LandingPage",
  data() {
    return {
      articles: [],
    };
  },
  async created() {
    try {
      const response = await axios.get("/articles/popular?source=detik");
      if (response.data.status === "Success" && Array.isArray(response.data.articles)) {
        this.articles = response.data.articles;
      } else {
        console.error("Error: Articles are not available or status is not Success");
      }
    } catch (error) {
      console.error("Error fetching popular articles:", error);
    }
  },
};
</script>

<style scoped>
/* Global Styles */
.landing-page {
  font-family: 'Arial', sans-serif;
  padding: 30px;
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
    grid-template-columns: repeat(2, 1fr); /* 2 columns for medium screens */
  }
}

@media (min-width: 1024px) {
  .articles-container {
    grid-template-columns: repeat(3, 1fr); /* 3 columns for large screens */
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
  margin: 0 auto; /* Center each card within the grid cell */
}

.article-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.article-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #2c3e50; /* Dark blue for titles */
  margin-bottom: 10px;
  text-align: left;
}

.article-description {
  font-size: 1rem;
  color: #555;
  margin-bottom: 15px;
  text-align: left;
}

.article-links {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.read-more-link,
.read-link {
  color: #3498db; /* Bright blue for links */
  text-decoration: none;
  font-weight: bold;
  border: 1px solid #3498db;
  padding: 5px 10px;
  border-radius: 5px;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.read-more-link:hover,
.read-link:hover {
  background-color: #3498db; /* Blue background on hover */
  color: white; /* White text on hover */
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

@media (max-width: 768px) {
  .page-title {
    font-size: 2rem;
  }

  .article-card {
    padding: 15px;
  }

  .article-title {
    font-size: 1.3rem;
  }
}
</style>

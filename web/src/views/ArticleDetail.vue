<template>
    <div class="article-detail">
      <div v-if="article" class="article-content">
        <h1 class="article-title">{{ article.title }}</h1>
        <p class="article-metadata">
          <strong>Author:</strong> {{ article.author || 'Unknown' }} <br />
          <strong>Published:</strong> {{ article.timestamp || 'N/A' }}
        </p>
        <img v-if="article.img_url" :src="article.img_url" alt="Article Image" class="article-image" />
        <div class="article-body" v-html="article.content"></div>
      </div>
      <p v-else class="loading-message">Loading article details...</p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: 'ArticleDetail',
    data() {
      return {
        article: null,
      };
    },
    async created() {
      const detailUrl = this.$route.query.detailUrl;
      const sourceWebsite = this.$route.query.source;
      if (!detailUrl) {
        console.error('Detail URL is missing');
        return;
      }
      if (!sourceWebsite) {
        console.error('Source is missing')
      }
      try {
        const response = await axios.get(`/article?source=${sourceWebsite}&detailUrl=${detailUrl}`);
        if (response.data.status === 'Success' && response.data.articles.length) {
          this.article = response.data.articles[0];
        } else {
          console.error('Error: Article detail is not available or status is not Success');
        }
      } catch (error) {
        console.error('Error fetching article details:', error);
      }
    },
  };
  </script>
  
  <style scoped>
  .article-detail {
    font-family: 'Arial', sans-serif;
    padding: 30px;
    max-width: 800px;
    margin: 0 auto;
    background-color: #f9f9f9; /* Light gray background */
    color: #333; /* Dark gray text */
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); /* Subtle shadow for ambience */
  }
  
  .article-title {
    font-size: 2rem;
    margin-bottom: 20px;
    color: #2c3e50; /* Dark blue for titles */
    text-align: left;
  }
  
  .article-metadata {
    font-size: 1rem;
    color: #7f8c8d; /* Subtle gray for metadata */
    margin-bottom: 20px;
    text-align: left;
  }
  
  .article-image {
    width: 100%;
    height: auto;
    border-radius: 8px;
    margin-bottom: 20px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1); /* Shadow for better visibility */
  }
  
  .article-body {
    font-size: 1.1rem;
    line-height: 1.8;
    color: #333;
    text-align: left;
    overflow-wrap: break-word;
    word-wrap: break-word;
    word-break: break-word;
  }
  
  .loading-message {
    font-size: 1.2rem;
    color: #888; /* Gray for loading text */
    text-align: center;
    margin-top: 20px;
  }
  </style>
  
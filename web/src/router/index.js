import { createRouter, createWebHistory } from 'vue-router';
import LandingPage from '../views/LandingPage.vue';
import ArticleDetail from '../views/ArticleDetail.vue';
import HowItWorks from '../views/HowItWorks.vue';
import ContentAttribution from '../views/ContentAttribution.vue';

const routes = [
  {
    path: '/',
    name: 'Gober - Berita Bebas Iklan - Home',
    component: LandingPage,
  },
  {
    path: '/detail',
    name: 'Gober - Berita Bebas Iklan - Detail',
    component: ArticleDetail,
  },
  {
    path: '/about/how-it-works',
    name: 'Gober - How It Works',
    component: HowItWorks,
  },
  {
    path: '/about/content-attribution',
    name: 'Gober - Content Attribution',
    component: ContentAttribution,
  },
];

const router = createRouter({
  history: createWebHistory(), // No need for base
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.name;
  next();
});

export default router;

console.log("hello from router/index.js")
console.log("be url: " + process.env.VUE_APP_GOBER_API_URL)
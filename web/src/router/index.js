import { createRouter, createWebHistory } from 'vue-router';
import LandingPage from '../views/LandingPage.vue';
import ArticleDetail from '../views/ArticleDetail.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: LandingPage,
  },
  {
    path: '/detail',
    name: 'ArticleDetail',
    component: ArticleDetail,
  },
];

const router = createRouter({
  history: createWebHistory(), // No need for base
  routes,
});

export default router;

console.log("hello from router/index.js")
import { createRouter, createWebHistory } from 'vue-router';
import LandingPage from '../views/LandingPage.vue';
import ArticleDetail from '../views/ArticleDetail.vue';

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
import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../views/LandingPage.vue';
import ArticleDetail from '../views/ArticleDetail.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: LandingPage
  },
  {
    path: '/detail',
    name: 'ArticleDetail', 
    component: ArticleDetail
  },
]

const router = createRouter({
  history: createWebHistory(),
  base: process.env.BASE_URL,
  routes,
})

export default router

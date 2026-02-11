import { createRouter, createWebHistory } from 'vue-router';
import app from './views/index.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: app,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;

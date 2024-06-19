import { createRouter, createWebHistory } from 'vue-router'
import home from '../views/home.vue'
import register from '../views/register.vue'
import login from '../views/login.vue'
import profile from '../views/profile.vue'
import library from '../views/library.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: home,
    },
    {
      path: '/register',
      component: register,
    },
    {
      path: '/login',
      component: login,
    },
    {
      path: '/profile',
      component: profile,
    },
    {
      path: '/library',
      component: library,
    }
  ]
})

export default router
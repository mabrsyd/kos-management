import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from './pages/Dashboard.vue'
import Users from './pages/Users.vue'
import Kamars from './pages/Kamars.vue'
import Penyewas from './pages/Penyewas.vue'
import Tagihans from './pages/Tagihans.vue'
import Transaksis from './pages/Transaksis.vue'
import Report from './pages/Report.vue'
import Login from './pages/Login.vue'
import Profile from './pages/Profile.vue'

const routes = [
  { 
    path: '/login', 
    name: 'login',
    component: Login,
    meta: { requiresAuth: false }
  },
  { 
    path: '/', 
    name: 'dashboard',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  { 
    path: '/users', 
    name: 'users',
    component: Users,
    meta: { requiresAuth: true }
  },
  { 
    path: '/kamars', 
    name: 'kamars',
    component: Kamars,
    meta: { requiresAuth: true }
  },
  { 
    path: '/penyewas', 
    name: 'penyewas',
    component: Penyewas,
    meta: { requiresAuth: true }
  },
  { 
    path: '/tagihans', 
    name: 'tagihans',
    component: Tagihans,
    meta: { requiresAuth: true }
  },
  { 
    path: '/transaksis', 
    name: 'transaksis',
    component: Transaksis,
    meta: { requiresAuth: true }
  },
  { 
    path: '/report', 
    name: 'report',
    component: Report,
    meta: { requiresAuth: true }
  },
  { 
    path: '/profile', 
    name: 'profile',
    component: Profile,
    meta: { requiresAuth: true }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('token')
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else if (to.path === '/login' && isAuthenticated) {
    next('/')
  } else {
    next()
  }
})

export default router 
<template>
  <header class="h-16 bg-white dark:bg-slate-800 shadow-sm border-b border-gray-200 dark:border-slate-700 flex items-center px-4 justify-between sticky top-0 z-30 transition-colors duration-200">
    <div class="flex items-center gap-3">
      <!-- Mobile menu toggle -->
      <button 
        @click="$emit('toggle-menu')" 
        class="p-2 rounded-lg text-gray-500 hover:bg-gray-100 dark:text-gray-400 dark:hover:bg-slate-700 lg:hidden transition-colors duration-200"
        aria-label="Toggle Menu"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
      
      <h1 class="text-lg font-medium text-gray-800 dark:text-gray-200">
        <span v-if="pageTitle">{{ pageTitle }}</span>
        <span v-else>Dashboard</span>
      </h1>
    </div>
    
    <div class="flex items-center gap-2 md:gap-3">
      <!-- Theme toggle button -->
      <button 
        @click="$emit('toggle-dark')" 
        class="p-2 rounded-lg text-gray-500 hover:bg-gray-100 dark:text-gray-400 dark:hover:bg-slate-700 transition-colors duration-200" 
        aria-label="Toggle Dark Mode"
      >
        <svg v-if="!isDark" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v1m0 16v1m8.66-13.66l-.71.71M4.05 19.95l-.71.71M21 12h-1M4 12H3m16.66 6.66l-.71-.71M4.05 4.05l-.71-.71M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
        </svg>
        <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M21 12.79A9 9 0 1111.21 3a7 7 0 109.79 9.79z"/>
        </svg>
      </button>
      
      <!-- Notifications button - example feature (hidden on smallest screens) -->
      <button class="hidden xs:block p-2 rounded-lg text-gray-500 hover:bg-gray-100 dark:text-gray-400 dark:hover:bg-slate-700 transition-colors duration-200 relative" aria-label="Notifications">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
        </svg>
        <span class="absolute top-1 right-1 h-2 w-2 rounded-full bg-red-500 ring-2 ring-white dark:ring-slate-800"></span>
      </button>
      
      <!-- User Dropdown -->
      <div class="relative" ref="profileDropdown">
        <button 
          @click="toggleDropdown" 
          class="flex items-center gap-2 hover:bg-gray-100 dark:hover:bg-slate-700 py-1 px-2 rounded-lg transition-colors duration-200"
        >
          <div class="flex-shrink-0 h-8 w-8 rounded-full bg-gradient-to-r from-blue-500 to-blue-600 flex items-center justify-center text-white font-medium shadow-sm">
            {{ userInitials }}
          </div>
          <div class="hidden md:block text-left">
            <div class="text-sm font-medium text-gray-800 dark:text-gray-200">{{ user.name }}</div>
            <div class="text-xs text-gray-500 dark:text-gray-400">{{ user.role }}</div>
          </div>
          <svg class="w-4 h-4 text-gray-400 ml-1 hidden sm:block" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
          </svg>
        </button>
        
        <!-- Dropdown Menu -->
        <div 
          v-if="isDropdownOpen" 
          class="absolute right-0 mt-2 w-56 bg-white dark:bg-slate-800 rounded-lg shadow-lg overflow-hidden z-50 border border-gray-200 dark:border-slate-700 animate-fade-in-down"
        >
          <div class="p-4 border-b border-gray-200 dark:border-slate-700">
            <div class="text-sm font-medium text-gray-800 dark:text-gray-200">{{ user.name }}</div>
            <div class="text-xs text-gray-500 dark:text-gray-400">{{ user.username }}</div>
          </div>
          <div class="py-1">
            <router-link 
              to="/profile" 
              class="flex items-center px-4 py-2 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-slate-700 transition-colors duration-150"
              @click="isDropdownOpen = false"
            >
              <svg class="mr-3 w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
              Profil
            </router-link>
            <a 
              href="#" 
              @click.prevent="handleLogout" 
              class="flex items-center px-4 py-2 text-sm text-red-600 dark:text-red-400 hover:bg-gray-100 dark:hover:bg-slate-700 transition-colors duration-150"
            >
              <svg class="mr-3 w-5 h-5 text-red-500 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
              </svg>
              Logout
            </a>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const isDark = computed(() => document.documentElement.classList.contains('dark'))
const isDropdownOpen = ref(false)
const profileDropdown = ref(null)

// Define emits
const emit = defineEmits(['toggle-dark', 'toggle-menu'])

// Compute current page title from route
const pageTitle = computed(() => {
  switch (route.path) {
    case '/': return 'Dashboard'
    case '/kamars': return 'Kamar'
    case '/penyewas': return 'Penyewa'
    case '/tagihans': return 'Tagihan'
    case '/transaksis': return 'Transaksi'
    case '/report': return 'Report'
    case '/users': return 'Users'
    case '/profile': return 'Profile'
    default: return ''
  }
})

// Get user data from localStorage
const user = ref({
  name: 'User',
  username: '',
  role: '',
})

// Get user initials for avatar
const userInitials = computed(() => {
  if (!user.value.name) return 'U'
  return user.value.name
    .split(' ')
    .map(name => name.charAt(0))
    .join('')
    .substring(0, 2)
    .toUpperCase()
})

// Toggle dropdown menu
const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

// Handle click outside to close dropdown
const handleClickOutside = (event) => {
  if (profileDropdown.value && !profileDropdown.value.contains(event.target)) {
    isDropdownOpen.value = false
  }
}

// Handle logout function
const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/login')
}

// Load user data on component mount
onMounted(() => {
  // Add click outside listener
  document.addEventListener('click', handleClickOutside)
  
  // Load user data from localStorage
  const userData = localStorage.getItem('user')
  if (userData) {
    try {
      const parsedUser = JSON.parse(userData)
      user.value = {
        name: parsedUser.name || 'User',
        username: parsedUser.username || '',
        role: parsedUser.role || '',
      }
    } catch (error) {
      console.error('Error parsing user data:', error)
    }
  }
})

// Clean up event listener on component unmount
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// Watch route changes to close dropdown
watch(route, () => {
  isDropdownOpen.value = false
})
</script>

<style scoped>
.animate-fade-in-down {
  animation: fadeInDown 0.2s ease-out forwards;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Add xs breakpoint for smallest screens */
@media (min-width: 480px) {
  .xs\:block {
    display: block;
  }
}
</style> 
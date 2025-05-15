<template>
  <div class="flex min-h-screen bg-gray-50 dark:bg-slate-900 text-gray-800 dark:text-gray-200 antialiased overflow-hidden">
    <!-- Mobile sidebar overlay -->
    <div 
      v-if="isMobileMenuOpen" 
      class="fixed inset-0 bg-gray-900/50 backdrop-blur-sm z-40 lg:hidden transition-opacity duration-300"
      @click="isMobileMenuOpen = false"
    ></div>

    <!-- Sidebar component with responsive behavior -->
    <div 
      :class="[
        'fixed top-0 left-0 h-screen z-40 transition-all duration-300',
        isMobileMenuOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0',
        'w-[260px]'
      ]"
    >
      <Sidebar @close-menu="isMobileMenuOpen = false" />
    </div>
    
    <!-- Main content area -->
    <div class="flex-1 flex flex-col min-h-screen lg:pl-[260px] transition-all duration-300 w-full">
      <Navbar @toggle-dark="toggleDarkMode" @toggle-menu="toggleMobileMenu" />
      <main class="flex-1 p-4 md:p-6 bg-gray-50 dark:bg-slate-900 transition-colors duration-300 overflow-auto">
        <slot />
      </main>
      <footer class="py-3 px-4 mt-auto text-center">
        <div class="text-xs text-gray-500 dark:text-gray-400">
          &copy; {{ new Date().getFullYear() }} Kos Management System. All rights reserved.
        </div>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import { useRoute } from 'vue-router'
import Sidebar from '../components/layout/Sidebar.vue'
import Navbar from '../components/layout/Navbar.vue'

const isMobileMenuOpen = ref(false)
const route = useRoute()

function toggleDarkMode() {
  document.documentElement.classList.toggle('dark')
  localStorage.setItem('theme', document.documentElement.classList.contains('dark') ? 'dark' : 'light')
}

function toggleMobileMenu() {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

// Close mobile menu when route changes
watch(route, () => {
  isMobileMenuOpen.value = false
})

// Close mobile menu when screen resizes to desktop
function handleResize() {
  if (window.innerWidth >= 1024) { // lg breakpoint
    isMobileMenuOpen.value = false
  }
}

// Handle escape key to close mobile menu
function handleEscKey(e) {
  if (e.key === 'Escape' && isMobileMenuOpen.value) {
    isMobileMenuOpen.value = false
  }
}

// Initialize theme based on user preference
onMounted(() => {
  // Check for saved theme or system preference
  const savedTheme = localStorage.getItem('theme')
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  
  if (savedTheme === 'dark' || (!savedTheme && prefersDark)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
  
  // Add event listeners
  window.addEventListener('resize', handleResize)
  document.addEventListener('keydown', handleEscKey)
})

onBeforeUnmount(() => {
  // Clean up event listeners
  window.removeEventListener('resize', handleResize)
  document.removeEventListener('keydown', handleEscKey)
})
</script>

<style>
/* Apply smooth scrolling to the entire page */
html {
  scroll-behavior: smooth;
}

/* Ensure a minimal width for the layout to avoid issues with very small screens */
.min-h-screen {
  min-width: 320px;
}

/* For Webkit browsers (Chrome, Safari) */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  @apply bg-gray-100 dark:bg-slate-800;
}

::-webkit-scrollbar-thumb {
  @apply bg-gray-300 dark:bg-slate-700 rounded-full;
}

::-webkit-scrollbar-thumb:hover {
  @apply bg-gray-400 dark:bg-slate-600;
}

/* Prevent scroll on body when mobile menu is open */
body.menu-open {
  overflow: hidden;
}

/* Ensure proper sizing of content areas */
@media screen and (min-width: 1024px) {
  main {
    max-width: calc(100vw - 260px);
  }
}
</style> 
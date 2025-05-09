<template>
  <div class="flex min-h-screen bg-blue-50 dark:bg-slate-900">
    <Sidebar />
    <div class="flex-1 flex flex-col min-h-screen pl-64">
      <Navbar @toggle-dark="toggleDarkMode" />
      <main class="flex-1 p-6 md:p-10 bg-blue-50 dark:bg-slate-900 transition-colors duration-300">
        <slot />
      </main>
      <footer class="text-xs text-gray-400 text-center py-4 mt-auto">&copy; 2024 Kos Management. All rights reserved.</footer>
    </div>
  </div>
</template>
<script setup>
import { onMounted } from 'vue'
import Sidebar from '../components/layout/Sidebar.vue'
import Navbar from '../components/layout/Navbar.vue'

function toggleDarkMode() {
  document.documentElement.classList.toggle('dark')
  localStorage.setItem('theme', document.documentElement.classList.contains('dark') ? 'dark' : 'light')
}

// Inisialisasi tema berdasarkan preferensi user yang tersimpan
onMounted(() => {
  // Set default light theme
  const savedTheme = localStorage.getItem('theme') || 'light'
  if (savedTheme === 'dark') {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
})
</script> 
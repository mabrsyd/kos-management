<template>
  <header class="h-16 bg-white dark:bg-slate-900 shadow flex items-center px-8 justify-between sticky top-0 z-30">
    <div class="flex items-center gap-3">
    </div>
    <div class="flex items-center gap-4">
      <button class="bg-blue-100 text-blue-700 px-3 py-1 rounded hover:bg-blue-200 transition">Logout</button>
      <button @click="$emit('toggle-dark')" class="p-2 rounded hover:bg-blue-50 dark:hover:bg-slate-800 transition">
        <svg v-if="!isDark" class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 3v1m0 16v1m8.66-13.66l-.71.71M4.05 19.95l-.71.71M21 12h-1M4 12H3m16.66 6.66l-.71-.71M4.05 4.05l-.71-.71M16 12a4 4 0 11-8 0 4 4 0 018 0z"/></svg>
        <svg v-else class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M21 12.79A9 9 0 1111.21 3a7 7 0 109.79 9.79z"/></svg>
      </button>
      <div class="w-9 h-9 rounded-full bg-blue-200 flex items-center justify-center font-bold text-blue-700">A</div>
    </div>
  </header>
</template>
<script setup>
import { ref, computed } from 'vue'
const isDark = computed(() => document.documentElement.classList.contains('dark'))
</script> 
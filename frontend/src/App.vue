<script setup>
import DefaultLayout from './layouts/DefaultLayout.vue'
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const notif = ref('')
const notifType = ref('success') // success, error, warning, info

function setNotif(msg, type = 'success') {
  notif.value = msg
  notifType.value = type
  setTimeout(() => notif.value = '', 3000)
}

// Check if the current page requires layout
const requiresLayout = computed(() => {
  return route.path !== '/login'
})
</script>

<template>
  <template v-if="requiresLayout">
    <DefaultLayout>
      <div 
        v-if="notif" 
        class="fixed top-20 right-4 px-4 py-3 rounded-lg shadow-lg z-50 flex items-center gap-3 transition-all duration-300 transform animate-slideIn"
        :class="{
          'bg-green-100 text-green-800 dark:bg-green-800/30 dark:text-green-200': notifType === 'success',
          'bg-red-100 text-red-800 dark:bg-red-800/30 dark:text-red-200': notifType === 'error',
          'bg-yellow-100 text-yellow-800 dark:bg-yellow-800/30 dark:text-yellow-200': notifType === 'warning',
          'bg-blue-100 text-blue-800 dark:bg-blue-800/30 dark:text-blue-200': notifType === 'info'
        }"
      >
        <svg v-if="notifType === 'success'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <svg v-if="notifType === 'error'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <svg v-if="notifType === 'warning'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
        <svg v-if="notifType === 'info'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        {{ notif }}
      </div>
      <router-view :setNotif="setNotif" />
    </DefaultLayout>
  </template>
  <template v-else>
    <div 
      v-if="notif" 
      class="fixed top-4 right-4 px-4 py-3 rounded-lg shadow-lg z-50 flex items-center gap-3 transition-all duration-300 transform animate-slideIn"
      :class="{
        'bg-green-100 text-green-800 dark:bg-green-800/30 dark:text-green-200': notifType === 'success',
        'bg-red-100 text-red-800 dark:bg-red-800/30 dark:text-red-200': notifType === 'error',
        'bg-yellow-100 text-yellow-800 dark:bg-yellow-800/30 dark:text-yellow-200': notifType === 'warning',
        'bg-blue-100 text-blue-800 dark:bg-blue-800/30 dark:text-blue-200': notifType === 'info'
      }"
    >
      <svg v-if="notifType === 'success'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <svg v-if="notifType === 'error'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <svg v-if="notifType === 'warning'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
      </svg>
      <svg v-if="notifType === 'info'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      {{ notif }}
    </div>
    <router-view :setNotif="setNotif" />
  </template>
</template>

<style>
/* Global styling */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');

body {
  font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* Animation for notifications */
.animate-slideIn {
  animation: slideIn 0.3s ease-out forwards;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}
</style>

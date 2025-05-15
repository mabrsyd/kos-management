<template>
  <aside class="h-full w-full bg-white dark:bg-slate-800 shadow-md border-r border-gray-200 dark:border-slate-700 flex flex-col z-40 transition-all duration-300 overflow-hidden">
    <!-- Logo and Brand with close button for mobile -->
    <div class="flex items-center justify-between h-16 border-b border-gray-200 dark:border-slate-700 px-4">
      <div class="flex items-center gap-3">
        <img src="/ipin.png" alt="Logo" class="w-8 h-8 rounded-lg" />
        <div>
          <span class="font-bold text-lg bg-gradient-to-r from-blue-400 to-green-500 bg-clip-text text-transparent">Kos Muhandis</span>
        </div>
      </div>
      
      <!-- Mobile close button -->
      <button 
        @click="$emit('close-menu')" 
        class="p-2 rounded-lg text-gray-500 hover:bg-gray-100 dark:hover:bg-slate-700 lg:hidden"
        aria-label="Close menu"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
    
    <!-- Navigation Menu -->
    <nav class="flex-1 py-4 px-3 overflow-y-auto scrollbar-thin">
      <div class="mb-5 px-3">
        <h3 class="text-xs uppercase font-medium text-gray-500 dark:text-gray-400 tracking-wider">Menu</h3>
      </div>
      <ul class="space-y-1">
        <li v-for="item in menu" :key="item.to">
          <router-link
            :to="item.to"
            class="flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-colors duration-200"
            :class="{ 
              'text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-slate-700': $route.path !== item.to,
              'bg-blue-50 text-green-600 dark:bg-blue-900/30 dark:text-green-400': $route.path === item.to 
            }"
          >
            <component :is="item.icon" class="w-5 h-5" :class="{ 'text-gray-500 dark:text-gray-400': $route.path !== item.to }" />
            {{ item.label }}
            
            <!-- Active indicator -->
            <div v-if="$route.path === item.to" class="ml-auto h-1.5 w-1.5 rounded-full bg-green-600 dark:bg-green-400"></div>
          </router-link>
        </li>
      </ul>
      
      <!-- Additional section - could be used for quick links, system info, etc. -->
      <div class="mt-6 px-3">
        <h3 class="text-xs uppercase font-medium text-gray-500 dark:text-gray-400 tracking-wider mb-3">Admin</h3>
        <ul class="space-y-1">
          <li>
            <router-link
              to="/users"
              class="flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-colors duration-200"
              :class="{ 
                'text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-slate-700': $route.path !== '/users',
                'bg-blue-50 text-blue-600 dark:bg-blue-900/30 dark:text-blue-400': $route.path === '/users' 
              }"
            >
              <svg class="w-5 h-5" :class="{ 'text-gray-500 dark:text-gray-400': $route.path !== '/users' }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
              Users
              <div v-if="$route.path === '/users'" class="ml-auto h-1.5 w-1.5 rounded-full bg-blue-600 dark:bg-blue-400"></div>
            </router-link>
          </li>
        </ul>
      </div>
    </nav>
    
    <!-- Footer section -->
    <div class="p-4 mt-auto border-t border-gray-200 dark:border-slate-700">
      <div class="flex items-center text-xs text-gray-500 dark:text-gray-400">
        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span>v1.0.0</span>
      </div>
    </div>
  </aside>
</template>

<script setup>
import { defineEmits } from 'vue'
import { HomeIcon, BuildingOffice2Icon, UsersIcon, DocumentTextIcon, BanknotesIcon, ChartBarIcon } from '@heroicons/vue/24/outline'

defineEmits(['close-menu'])

const menu = [
  { to: '/', label: 'Dashboard', icon: HomeIcon },
  { to: '/kamars', label: 'Kamar', icon: BuildingOffice2Icon },
  { to: '/penyewas', label: 'Penyewa', icon: UsersIcon },
  { to: '/tagihans', label: 'Tagihan', icon: DocumentTextIcon },
  { to: '/transaksis', label: 'Transaksi', icon: BanknotesIcon },
  { to: '/report', label: 'Report', icon: ChartBarIcon },
]
</script>

<style scoped>
.scrollbar-thin::-webkit-scrollbar {
  width: 4px;
}

.scrollbar-thin::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.scrollbar-thin::-webkit-scrollbar-thumb {
  @apply bg-gray-300 dark:bg-slate-700 rounded-full;
}
</style> 
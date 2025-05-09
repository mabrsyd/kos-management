<template>
  <aside class="fixed top-0 left-0 h-screen w-64 bg-white dark:bg-slate-900 shadow-lg flex flex-col">
    <div class="flex items-center gap-2 px-6 py-6">
      <HomeIcon class="w-7 h-7 text-blue-600" />
      <span class="font-bold text-xl text-blue-700 dark:text-blue-300">Kos Muhandis</span>
    </div>
    <nav class="flex-1">
      <ul class="space-y-1">
        <li v-for="item in menu" :key="item.to">
          <router-link
            :to="item.to"
            class="flex items-center gap-3 px-6 py-3 rounded-lg hover:bg-blue-50 dark:hover:bg-slate-800 transition font-medium"
            :class="{ 'bg-blue-100 text-blue-700 dark:bg-slate-800 dark:text-blue-300': $route.path === item.to }"
          >
            <component :is="item.icon" class="w-5 h-5" />
            {{ item.label }}
          </router-link>
        </li>
      </ul>
    </nav>
  </aside>
</template>
<script setup>
import { HomeIcon, BuildingOffice2Icon, UsersIcon, DocumentTextIcon, BanknotesIcon, ChartBarIcon } from '@heroicons/vue/24/outline'
const menu = [
  { to: '/', label: 'Home', icon: HomeIcon },
  { to: '/kamars', label: 'Kamar', icon: BuildingOffice2Icon },
  { to: '/penyewas', label: 'Penyewa', icon: UsersIcon },
  { to: '/tagihans', label: 'Tagihan', icon: DocumentTextIcon },
  { to: '/transaksis', label: 'Transaksi', icon: BanknotesIcon },
  { to: '/report', label: 'Report', icon: ChartBarIcon },
]
</script> 
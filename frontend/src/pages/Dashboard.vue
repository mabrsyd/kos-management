<template>
  <div class="space-y-10">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-10">
      <DashboardCard label="Total Kamar" :value="stat.kamar" :icon="icons.kamar" color="blue" />
      <DashboardCard label="Total Penyewa" :value="stat.penyewa" :icon="icons.penyewa" color="green" />
      <DashboardCard label="Tagihan Aktif" :value="stat.tagihan" :icon="icons.tagihan" color="yellow" />
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <AppChart
        chartType="bar"
        :options="chartOptions"
        :series="chartSeries"
        height="300"
      />
      <AppChart
        chartType="line"
        :options="chartOptions2"
        :series="chartSeries2"
        height="300"
      />
    </div>
    <div class="bg-white dark:bg-slate-900 rounded-xl shadow p-8 mt-10">
      <div class="text-2xl font-semibold mb-2 text-blue-700 dark:text-blue-300">Selamat datang di aplikasi manajemen kos!</div>
      <div class="text-gray-600 dark:text-gray-300">Gunakan menu di samping untuk mengelola data kamar, penyewa, tagihan, transaksi, dan melihat laporan.</div>
    </div>
  </div>
</template>

<script setup>
import DashboardCard from '../components/ui/DashboardCard.vue'
import AppChart from '../components/ui/AppChart.vue'
import { BuildingOffice2Icon, UsersIcon, DocumentTextIcon } from '@heroicons/vue/24/outline'

const stat = { kamar: 24, penyewa: 18, tagihan: 7 } // Dummy, ganti dengan fetch API
const icons = {
  kamar: BuildingOffice2Icon,
  penyewa: UsersIcon,
  tagihan: DocumentTextIcon,
}
const chartOptions = {
  chart: { toolbar: { show: false }, foreColor: '#64748b', fontFamily: 'Inter, sans-serif' },
  xaxis: { categories: ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun'] },
  theme: { mode: 'light' },
  colors: ['#2563eb'],
}
const chartSeries = [{ name: 'Pemasukan', data: [1200, 1400, 1100, 1600, 1800, 1500] }]
const chartOptions2 = {
  chart: { toolbar: { show: false }, foreColor: '#64748b', fontFamily: 'Inter, sans-serif' },
  xaxis: { categories: ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun'] },
  theme: { mode: 'light' },
  colors: ['#facc15'],
}
const chartSeries2 = [{ name: 'Tagihan', data: [5, 7, 6, 8, 7, 9] }]
</script> 
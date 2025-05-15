<template>
  <div class="space-y-6">
    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6">
      <DashboardCard label="Total Kamar" :value="stats.totalKamar" :icon="icons.kamar" color="blue" />
      <DashboardCard label="Total Penyewa" :value="stats.totalPenyewa" :icon="icons.penyewa" color="green" />
      <DashboardCard label="Tagihan Belum Lunas" :value="stats.tagihanBelumLunas" :icon="icons.tagihan" color="yellow" />
      <DashboardCard label="Tagihan Jatuh Tempo" :value="stats.tagihanJatuhTempo" :icon="icons.debt" color="red" />
    </div>
    
    <!-- Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 md:gap-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <h3 class="text-base md:text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Pendapatan & Pengeluaran</h3>
        <AppChart
          chartType="bar"
          :options="incomeChartOptions"
          :series="incomeChartSeries"
          height="300"
        />
      </div>
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <h3 class="text-base md:text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Status Kamar</h3>
        <AppChart
          chartType="pie"
          :options="roomChartOptions"
          :series="roomChartSeries"
          height="300"
        />
      </div>
    </div>
    
    <!-- Tables -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 md:gap-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-base md:text-lg font-semibold text-gray-700 dark:text-gray-300">Tagihan Terbaru</h3>
          <router-link to="/tagihans" class="text-sm text-blue-600 hover:text-blue-800 dark:text-blue-400 dark:hover:text-blue-300 flex items-center gap-1">
            <span>Lihat Semua</span>
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path></svg>
          </router-link>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
            <thead class="bg-gray-50 dark:bg-slate-700">
              <tr>
                <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Penyewa</th>
                <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Jumlah</th>
                <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Status</th>
              </tr>
            </thead>
            <tbody class="bg-white dark:bg-slate-800 divide-y divide-gray-200 dark:divide-gray-700">
              <tr v-for="tagihan in latestTagihans" :key="tagihan.id" class="hover:bg-gray-50 dark:hover:bg-slate-700 transition-colors duration-150">
                <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ tagihan.penyewa_nama }}</td>
                <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ formatRupiah(tagihan.jumlah_tagihan) }}</td>
                <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap">
                  <span 
                    :class="{
                      'px-2 py-1 rounded-full text-xs font-medium inline-flex items-center': true,
                      'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400': tagihan.status === 'lunas',
                      'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400': tagihan.status === 'belum_lunas',
                      'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400': tagihan.status === 'jatuh_tempo'
                    }"
                  >
                    <span class="h-1.5 w-1.5 rounded-full mr-1" :class="{
                      'bg-green-500 dark:bg-green-400': tagihan.status === 'lunas',
                      'bg-yellow-500 dark:bg-yellow-400': tagihan.status === 'belum_lunas',
                      'bg-red-500 dark:bg-red-400': tagihan.status === 'jatuh_tempo'
                    }"></span>
                    {{ formatStatus(tagihan.status) }}
                  </span>
                </td>
              </tr>
              <tr v-if="latestTagihans.length === 0">
                <td colspan="3" class="px-3 py-2 md:px-4 md:py-3 text-center text-sm text-gray-500 dark:text-gray-400">
                  Belum ada tagihan
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-base md:text-lg font-semibold text-gray-700 dark:text-gray-300">Transaksi Terbaru</h3>
          <router-link to="/transaksis" class="text-sm text-blue-600 hover:text-blue-800 dark:text-blue-400 dark:hover:text-blue-300 flex items-center gap-1">
            <span>Lihat Semua</span>
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path></svg>
          </router-link>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
            <thead class="bg-gray-50 dark:bg-slate-700">
              <tr>
                <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Tipe</th>
                <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Jumlah</th>
                <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Tanggal</th>
              </tr>
            </thead>
            <tbody class="bg-white dark:bg-slate-800 divide-y divide-gray-200 dark:divide-gray-700">
              <tr v-for="transaksi in latestTransaksis" :key="transaksi.id" class="hover:bg-gray-50 dark:hover:bg-slate-700 transition-colors duration-150">
                <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm">
                  <span 
                    :class="{
                      'px-2 py-1 rounded-full text-xs font-medium inline-flex items-center': true,
                      'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400': transaksi.tipe === 'pemasukan',
                      'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400': transaksi.tipe === 'pengeluaran'
                    }"
                  >
                    <span class="h-1.5 w-1.5 rounded-full mr-1" :class="{
                      'bg-green-500 dark:bg-green-400': transaksi.tipe === 'pemasukan',
                      'bg-red-500 dark:bg-red-400': transaksi.tipe === 'pengeluaran'
                    }"></span>
                    {{ transaksi.tipe }}
                  </span>
                </td>
                <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ formatRupiah(transaksi.jumlah) }}</td>
                <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ formatDate(transaksi.tanggal_transaksi) }}</td>
              </tr>
              <tr v-if="latestTransaksis.length === 0">
                <td colspan="3" class="px-3 py-2 md:px-4 md:py-3 text-center text-sm text-gray-500 dark:text-gray-400">
                  Belum ada transaksi
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
    
    <!-- Add a new section to show tenants with debt after the two tables -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-base md:text-lg font-semibold text-gray-700 dark:text-gray-300">Penyewa dengan Tunggakan</h3>
        <router-link to="/penyewas" class="text-sm text-blue-600 hover:text-blue-800 dark:text-blue-400 dark:hover:text-blue-300 flex items-center gap-1">
          <span>Lihat Semua</span>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path></svg>
        </router-link>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead class="bg-gray-50 dark:bg-slate-700">
            <tr>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Penyewa</th>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Kamar</th>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Total Tunggakan</th>
            </tr>
          </thead>
          <tbody class="bg-white dark:bg-slate-800 divide-y divide-gray-200 dark:divide-gray-700">
            <tr v-for="penyewa in tenantsWithDebt" :key="penyewa.id" class="hover:bg-gray-50 dark:hover:bg-slate-700 transition-colors duration-150">
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ penyewa.nama }}</td>
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">
                Kamar {{ penyewa.kamar_info }}
              </td>
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200 font-semibold text-red-600 dark:text-red-400">
                {{ formatRupiah(penyewa.total_debt) }}
              </td>
            </tr>
            <tr v-if="tenantsWithDebt.length === 0">
              <td colspan="3" class="px-3 py-2 md:px-4 md:py-3 text-center text-sm text-gray-500 dark:text-gray-400">
                Tidak ada penyewa dengan tunggakan
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    
    <!-- Welcome Banner -->
    <div class="bg-gradient-to-r from-blue-500 to-blue-700 dark:from-blue-700 dark:to-blue-900 rounded-xl shadow-sm p-6 mt-4 transition-colors duration-200">
      <div class="text-xl md:text-2xl font-semibold mb-2 text-white">Selamat datang di aplikasi manajemen kos!</div>
      <div class="text-blue-100">Gunakan menu di samping untuk mengelola data kamar, penyewa, tagihan, transaksi, dan melihat laporan.</div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import DashboardCard from '../components/ui/DashboardCard.vue'
import AppChart from '../components/ui/AppChart.vue'
import { BuildingOffice2Icon, UsersIcon, DocumentTextIcon, ExclamationTriangleIcon } from '@heroicons/vue/24/outline'
import api from '../utils/api'

// Data sources
const kamars = ref([])
const penyewas = ref([])
const tagihans = ref([])
const transaksis = ref([])
const tenantsWithDebt = ref([])
const loading = ref(false)

// Stats computed properties
const stats = computed(() => {
  return {
    totalKamar: kamars.value.length,
    totalPenyewa: penyewas.value.length,
    tagihanBelumLunas: tagihans.value.filter(t => t.status === 'belum_lunas').length,
    tagihanJatuhTempo: tagihans.value.filter(t => t.status === 'jatuh_tempo').length,
  }
})

// Icons for dashboard cards
const icons = {
  kamar: BuildingOffice2Icon,
  penyewa: UsersIcon,
  tagihan: DocumentTextIcon,
  debt: ExclamationTriangleIcon,
}

// Format currency
const formatRupiah = (value) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value)
}

// Format date
const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('id-ID', { 
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  }).format(date)
}

// Format status
const formatStatus = (status) => {
  switch(status) {
    case 'lunas': return 'Lunas'
    case 'belum_lunas': return 'Belum Lunas'
    case 'jatuh_tempo': return 'Jatuh Tempo'
    default: return status
  }
}

// Computed properties for latest items
const latestTagihans = computed(() => {
  return tagihans.value
    .filter(t => t.status === 'belum_lunas' || t.status === 'jatuh_tempo')
    .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
    .slice(0, 5)
    .map(tagihan => {
      const penyewa = penyewas.value.find(p => p.id === tagihan.id_penyewa) || {}
      return {
        ...tagihan,
        penyewa_nama: penyewa.nama || `ID: ${tagihan.id_penyewa}`
      }
    })
})

const latestTransaksis = computed(() => {
  return transaksis.value
    .sort((a, b) => new Date(b.tanggal_transaksi) - new Date(a.tanggal_transaksi))
    .slice(0, 5)
})

// Income chart options and data
const incomeChartOptions = ref({
  chart: {
    type: 'bar',
    stacked: false,
    toolbar: { show: false },
    foreColor: '#9ca3af' // gray-400
  },
  plotOptions: {
    bar: {
      horizontal: false,
      columnWidth: '55%',
      borderRadius: 5
    },
  },
  dataLabels: { enabled: false },
  stroke: {
    show: true,
    width: 2,
    colors: ['transparent']
  },
  xaxis: { 
    categories: [],
    title: {
      text: 'Bulan',
      style: {
        fontSize: '14px',
        fontWeight: 600
      }
    }
  },
  yaxis: {
    title: {
      text: 'Total (Rupiah)',
      style: {
        fontSize: '14px',
        fontWeight: 600
      }
    },
    labels: {
      formatter: function (value) {
        return formatRupiah(value).split(',')[0]
      }
    }
  },
  fill: { opacity: 1 },
  tooltip: {
    y: {
      formatter: function (val) {
        return formatRupiah(val)
      }
    }
  },
  legend: { position: 'top' },
  colors: ['#10b981', '#ef4444'] // green-500, red-500
})

const incomeChartSeries = ref([
  { name: 'Pemasukan', data: [] },
  { name: 'Pengeluaran', data: [] }
])

// Room chart options and data
const roomChartOptions = ref({
  chart: {
    type: 'pie',
    foreColor: '#9ca3af' // gray-400
  },
  labels: ['Terisi', 'Kosong'],
  colors: ['#3b82f6', '#94a3b8'], // blue-500, gray-400
  legend: { position: 'bottom' },
  responsive: [{
    breakpoint: 480,
    options: {
      chart: { width: 200 },
      legend: { position: 'bottom' }
    }
  }],
  tooltip: {
    y: {
      formatter: function (val) {
        return val + ' kamar'
      }
    }
  }
})

const roomChartSeries = ref([0, 0])

// Fetch all data
async function fetchData() {
  loading.value = true
  
  try {
    await Promise.all([
      fetchKamars(),
      fetchPenyewas(),
      fetchTagihans(),
      fetchTransaksis()
    ])
    
    await fetchTenantsWithDebt()
    
    updateChartData()
  } catch (error) {
    console.error('Error fetching dashboard data:', error)
  } finally {
    loading.value = false
  }
}

async function fetchKamars() {
  try {
    const res = await api.get('/kamars')
    kamars.value = res.data
  } catch (error) {
    console.error('Error fetching kamars:', error)
  }
}

async function fetchPenyewas() {
  try {
    const res = await api.get('/penyewas')
    penyewas.value = res.data
  } catch (error) {
    console.error('Error fetching penyewas:', error)
  }
}

async function fetchTagihans() {
  try {
    const res = await api.get('/tagihans')
    tagihans.value = res.data
  } catch (error) {
    console.error('Error fetching tagihans:', error)
  }
}

async function fetchTransaksis() {
  try {
    const res = await api.get('/transaksis')
    transaksis.value = res.data
  } catch (error) {
    console.error('Error fetching transaksis:', error)
  }
}

async function fetchTenantsWithDebt() {
  try {
    loading.value = true
    const res = await api.get('/penyewas/with-debt')
    
    // For each tenant with debt, get their total debt
    const tenantsData = await Promise.all(res.data.map(async (penyewa) => {
      const debtRes = await api.get(`/tagihans/debt/by-penyewa/${penyewa.id}`)
      const kamar = kamars.value.find(k => k.id === penyewa.id_kamar) || {}
      return {
        ...penyewa,
        total_debt: debtRes.data.total_debt,
        kamar_info: kamar.nomer || '-'
      }
    }))
    
    tenantsWithDebt.value = tenantsData.sort((a, b) => b.total_debt - a.total_debt)
  } catch (error) {
    console.error('Error fetching tenants with debt:', error)
  } finally {
    loading.value = false
  }
}

// Update chart data based on current data
function updateChartData() {
  updateIncomeChart()
  updateRoomChart()
}

// Update income/expense chart
function updateIncomeChart() {
  // Get last 6 months
  const months = 6
  const monthData = {}
  
  // Initialize months
  for (let i = 0; i < months; i++) {
    const date = new Date()
    date.setMonth(date.getMonth() - i)
    const monthKey = `${date.getFullYear()}-${date.getMonth() + 1}`
    monthData[monthKey] = { 
      name: getMonthName(date),
      pemasukan: 0, 
      pengeluaran: 0 
    }
  }
  
  // Fill in data
  transaksis.value.forEach(t => {
    const date = new Date(t.tanggal_transaksi)
    const monthKey = `${date.getFullYear()}-${date.getMonth() + 1}`
    
    if (monthData[monthKey]) {
      if (t.tipe === 'pemasukan') {
        monthData[monthKey].pemasukan += t.jumlah
      } else {
        monthData[monthKey].pengeluaran += t.jumlah
      }
    }
  })
  
  // Sort months in ascending order
  const sortedMonths = Object.keys(monthData)
    .sort((a, b) => {
      const [yearA, monthA] = a.split('-').map(Number)
      const [yearB, monthB] = b.split('-').map(Number)
      if (yearA !== yearB) return yearA - yearB
      return monthA - monthB
    })
    .map(key => monthData[key])
  
  // Update chart data
  incomeChartOptions.value.xaxis.categories = sortedMonths.map(m => m.name)
  incomeChartSeries.value[0].data = sortedMonths.map(m => m.pemasukan)
  incomeChartSeries.value[1].data = sortedMonths.map(m => m.pengeluaran)
}

// Update room status chart
function updateRoomChart() {
  const occupied = kamars.value.filter(k => k.status === 'terisi').length
  const vacant = kamars.value.length - occupied
  
  roomChartSeries.value = [occupied, vacant]
}

// Get month name
const getMonthName = (date) => {
  return new Intl.DateTimeFormat('id-ID', { month: 'short' }).format(date)
}

onMounted(() => {
  fetchData()
})
</script> 
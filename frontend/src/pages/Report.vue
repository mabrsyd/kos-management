<template>
  <div>
    <div class="flex flex-wrap justify-between items-center gap-4 mb-6">
      <div>
        <h2 class="text-xl sm:text-2xl font-bold text-gray-800 dark:text-gray-100">Laporan Keuangan</h2>
        <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">Analisis pendapatan dan pengeluaran</p>
      </div>
      
      <!-- Filter Period -->
      <div class="flex flex-wrap items-center gap-3">
        <div>
          <label for="periodFilter" class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Periode</label>
          <select 
            v-model="periodFilter" 
            id="periodFilter"
            class="border border-gray-300 dark:border-slate-600 rounded-lg px-3 py-2 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent transition-colors duration-200"
            @change="updateChartData"
          >
            <option value="3">3 Bulan Terakhir</option>
            <option value="6">6 Bulan Terakhir</option>
            <option value="12">1 Tahun</option>
          </select>
        </div>
        
        <button 
          @click="fetchData" 
          class="flex items-center gap-2 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg shadow-sm transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          Refresh Data
        </button>
      </div>
    </div>
    
    <!-- Summary Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6 mb-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-5 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-green-100 dark:bg-green-900/30 p-2 mr-3">
            <svg class="w-5 h-5 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xs font-medium text-gray-500 dark:text-gray-400 mb-1">Total Pendapatan</h3>
            <p class="text-lg font-bold text-green-600 dark:text-green-400">{{ formatRupiah(totalPemasukan) }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-5 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-red-100 dark:bg-red-900/30 p-2 mr-3">
            <svg class="w-5 h-5 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xs font-medium text-gray-500 dark:text-gray-400 mb-1">Total Pengeluaran</h3>
            <p class="text-lg font-bold text-red-600 dark:text-red-400">{{ formatRupiah(totalPengeluaran) }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-5 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-blue-100 dark:bg-blue-900/30 p-2 mr-3">
            <svg class="w-5 h-5 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3" />
            </svg>
          </div>
          <div>
            <h3 class="text-xs font-medium text-gray-500 dark:text-gray-400 mb-1">Saldo</h3>
            <p class="text-lg font-bold" :class="saldo >= 0 ? 'text-blue-600 dark:text-blue-400' : 'text-red-600 dark:text-red-400'">
              {{ formatRupiah(saldo) }}
            </p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-5 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-indigo-100 dark:bg-indigo-900/30 p-2 mr-3">
            <svg class="w-5 h-5 text-indigo-600 dark:text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
          </div>
          <div>
            <h3 class="text-xs font-medium text-gray-500 dark:text-gray-400 mb-1">Tingkat Hunian</h3>
            <p class="text-lg font-bold text-indigo-600 dark:text-indigo-400">{{ occupancyRate }}%</p>
          </div>
        </div>
      </div>
    </div>
    
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 md:gap-6 mb-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <h3 class="text-base md:text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Pendapatan & Pengeluaran Bulanan</h3>
        <AppChart 
          chart-type="bar"
          :options="incomeChartOptions" 
          :series="incomeChartSeries"
          height="300" 
        />
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <h3 class="text-base md:text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Status Kamar</h3>
        <AppChart 
          chart-type="pie"
          :options="roomChartOptions" 
          :series="roomChartSeries"
          height="300" 
        />
      </div>
    </div>
    
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 mb-6 transition-colors duration-200">
      <h3 class="text-base md:text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Ringkasan Transaksi per Kategori</h3>
      <AppChart 
        chart-type="bar"
        :options="categoryChartOptions" 
        :series="categoryChartSeries"
        height="300" 
      />
    </div>
    
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
      <h3 class="text-base md:text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Daftar Tagihan Tertunggak</h3>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead class="bg-gray-50 dark:bg-slate-700">
            <tr>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Penyewa</th>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Kamar</th>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Periode</th>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Jumlah</th>
              <th class="px-3 py-2 md:px-4 md:py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Status</th>
            </tr>
          </thead>
          <tbody class="bg-white dark:bg-slate-800 divide-y divide-gray-200 dark:divide-gray-700">
            <tr v-for="tagihan in unpaidTagihans" :key="tagihan.id" class="hover:bg-gray-50 dark:hover:bg-slate-700 transition-colors duration-150">
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ tagihan.penyewa_nama }}</td>
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ tagihan.kamar_info }}</td>
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">
                {{ formatDate(tagihan.awal_periode) }} - {{ formatDate(tagihan.akhir_periode) }}
              </td>
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ formatRupiah(tagihan.jumlah_tagihan) }}</td>
              <td class="px-3 py-2 md:px-4 md:py-3 whitespace-nowrap">
                <span 
                  :class="{
                    'px-2 py-1 rounded-full text-xs font-medium inline-flex items-center': true,
                    'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400': tagihan.status === 'belum lunas',
                    'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400': tagihan.status === 'nyicil'
                  }"
                >
                  <span class="h-1.5 w-1.5 rounded-full mr-1" :class="{
                    'bg-red-500 dark:bg-red-400': tagihan.status === 'belum lunas',
                    'bg-yellow-500 dark:bg-yellow-400': tagihan.status === 'nyicil'
                  }"></span>
                  {{ tagihan.status }}
                </span>
              </td>
            </tr>
            <tr v-if="unpaidTagihans.length === 0">
              <td colspan="5" class="px-3 py-2 md:px-4 md:py-3 text-center text-sm text-gray-500 dark:text-gray-400">
                Tidak ada tagihan tertunggak
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '../utils/api'
import AppChart from '../components/ui/AppChart.vue'

// Data sources
const transaksis = ref([])
const tagihans = ref([])
const penyewas = ref([])
const kamars = ref([])
const periodFilter = ref('3')
const loading = ref(false)

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

// Get month name
const getMonthName = (date) => {
  return new Intl.DateTimeFormat('id-ID', { month: 'short' }).format(date)
}

// Calculate unpaid tagihans
const unpaidTagihans = computed(() => {
  return tagihans.value
    .filter(t => t.status === 'belum lunas' || t.status === 'nyicil')
    .map(tagihan => {
      const penyewa = penyewas.value.find(p => p.id === tagihan.id_penyewa) || {}
      const kamar = kamars.value.find(k => k.id === penyewa.id_kamar) || {}
      
      return {
        ...tagihan,
        penyewa_nama: penyewa.nama || `ID: ${tagihan.id_penyewa}`,
        kamar_info: kamar.nomer ? `Kamar ${kamar.nomer}` : '-',
      }
    })
})

// Calculate filtered transaksis (for current period)
const filteredTransaksis = computed(() => {
  const months = parseInt(periodFilter.value)
  const startDate = new Date()
  startDate.setMonth(startDate.getMonth() - months)
  
  return transaksis.value.filter(t => {
    const transDate = new Date(t.tanggal_transaksi)
    return transDate >= startDate
  })
})

// Calculate total income
const totalPemasukan = computed(() => {
  return filteredTransaksis.value
    .filter(t => t.tipe === 'pemasukan')
    .reduce((sum, t) => sum + t.jumlah, 0)
})

// Calculate total expenses
const totalPengeluaran = computed(() => {
  return filteredTransaksis.value
    .filter(t => t.tipe === 'pengeluaran')
    .reduce((sum, t) => sum + t.jumlah, 0)
})

// Calculate balance
const saldo = computed(() => {
  return totalPemasukan.value - totalPengeluaran.value
})

// Calculate occupancy rate
const occupancyRate = computed(() => {
  if (kamars.value.length === 0) return 0
  
  const occupied = kamars.value.filter(k => k.status === 'terisi').length
  return Math.round((occupied / kamars.value.length) * 100)
})

// Charts data
const incomeChartOptions = ref({
  chart: {
    type: 'bar',
    stacked: false,
    toolbar: {
      show: false
    },
    foreColor: '#9ca3af' // gray-400
  },
  plotOptions: {
    bar: {
      horizontal: false,
      columnWidth: '55%',
      borderRadius: 5
    },
  },
  dataLabels: {
    enabled: false
  },
  stroke: {
    show: true,
    width: 2,
    colors: ['transparent']
  },
  xaxis: {
    categories: [],
  },
  yaxis: {
    title: {
      text: 'Rupiah'
    },
    labels: {
      formatter: function (value) {
        return formatRupiah(value).split(',')[0]
      }
    }
  },
  fill: {
    opacity: 1
  },
  tooltip: {
    y: {
      formatter: function (val) {
        return formatRupiah(val)
      }
    }
  },
  legend: {
    position: 'top'
  },
  colors: ['#10b981', '#ef4444'] // green-500, red-500
})

const incomeChartSeries = ref([
  {
    name: 'Pemasukan',
    data: []
  },
  {
    name: 'Pengeluaran',
    data: []
  }
])

const roomChartOptions = ref({
  chart: {
    type: 'pie',
    foreColor: '#9ca3af' // gray-400
  },
  labels: ['Terisi', 'Kosong'],
  colors: ['#3b82f6', '#94a3b8'], // blue-500, gray-400
  legend: {
    position: 'bottom'
  },
  responsive: [{
    breakpoint: 480,
    options: {
      chart: {
        width: 200
      },
      legend: {
        position: 'bottom'
      }
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

const categoryChartOptions = ref({
  chart: {
    type: 'bar',
    toolbar: {
      show: false
    },
    foreColor: '#9ca3af' // gray-400
  },
  plotOptions: {
    bar: {
      horizontal: true,
      borderRadius: 5
    },
  },
  dataLabels: {
    enabled: false
  },
  xaxis: {
    categories: [],
    labels: {
      formatter: function (value) {
        return formatRupiah(value).split(',')[0]
      }
    }
  },
  colors: ['#3b82f6'] // blue-500
})

const categoryChartSeries = ref([{
  name: 'Jumlah',
  data: []
}])

// Fetch all data
async function fetchData() {
  loading.value = true
  
  try {
    await Promise.all([
      fetchTransaksis(),
      fetchTagihans(),
      fetchPenyewas(),
      fetchKamars()
    ])
    
    updateChartData()
  } catch (error) {
    console.error('Error fetching report data:', error)
  } finally {
    loading.value = false
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

async function fetchTagihans() {
  try {
    const res = await api.get('/tagihans')
    tagihans.value = res.data
  } catch (error) {
    console.error('Error fetching tagihans:', error)
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

async function fetchKamars() {
  try {
    const res = await api.get('/kamars')
    kamars.value = res.data
  } catch (error) {
    console.error('Error fetching kamars:', error)
  }
}

// Update chart data based on current data
function updateChartData() {
  updateIncomeChart()
  updateRoomChart()
  updateCategoryChart()
}

// Update income/expense chart
function updateIncomeChart() {
  // Get last X months
  const months = parseInt(periodFilter.value)
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
  filteredTransaksis.value.forEach(t => {
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

// Update category chart
function updateCategoryChart() {
  // Group transactions by category and calculate totals
  const categories = {}
  
  filteredTransaksis.value.forEach(t => {
    if (!categories[t.kategori]) {
      categories[t.kategori] = 0
    }
    categories[t.kategori] += t.jumlah
  })
  
  // Sort categories by amount (descending)
  const sortedCategories = Object.entries(categories)
    .sort((a, b) => b[1] - a[1])
    .slice(0, 10) // Top 10 categories
  
  categoryChartOptions.value.xaxis.categories = sortedCategories.map(([name]) => name)
  categoryChartSeries.value[0].data = sortedCategories.map(([, amount]) => amount)
}

onMounted(() => {
  fetchData()
})
</script> 
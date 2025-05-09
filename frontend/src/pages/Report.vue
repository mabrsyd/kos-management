<template>
  <div class="p-2 max-w-8xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h2 class="text-2xl font-bold text-blue-700 dark:text-blue-300">Laporan Keuangan</h2>
      
      <!-- Filter Period -->
      <div class="flex items-center gap-4">
        <div>
          <label for="periodFilter" class="block text-sm font-medium text-gray-600 dark:text-gray-300 mb-1">Periode</label>
          <select 
            v-model="periodFilter" 
            id="periodFilter"
            class="border border-gray-300 rounded-md px-3 py-2 focus:ring-blue-500 focus:border-blue-500 dark:bg-slate-700 dark:border-slate-600"
            @change="updateChartData"
          >
            <option value="3">3 Bulan Terakhir</option>
            <option value="6">6 Bulan Terakhir</option>
            <option value="12">1 Tahun</option>
          </select>
        </div>
        
        <button 
          @click="fetchData" 
          class="self-end bg-blue-600 text-white px-4 py-2 rounded-lg font-medium hover:bg-blue-700 transition"
        >
          Refresh Data
        </button>
      </div>
    </div>
    
    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
        <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">Total Pendapatan</h3>
        <p class="text-2xl font-bold text-green-600 dark:text-green-400">{{ formatRupiah(totalPemasukan) }}</p>
      </div>
      
      <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
        <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">Total Pengeluaran</h3>
        <p class="text-2xl font-bold text-red-600 dark:text-red-400">{{ formatRupiah(totalPengeluaran) }}</p>
      </div>
      
      <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
        <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">Saldo</h3>
        <p class="text-2xl font-bold" :class="saldo >= 0 ? 'text-blue-600 dark:text-blue-400' : 'text-red-600 dark:text-red-400'">
          {{ formatRupiah(saldo) }}
        </p>
      </div>
      
      <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
        <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">Tingkat Hunian</h3>
        <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ occupancyRate }}%</p>
      </div>
    </div>
    
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
      <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
        <h3 class="text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Pendapatan & Pengeluaran Bulanan</h3>
        <AppChart 
          chart-type="bar"
          :options="incomeChartOptions" 
          :series="incomeChartSeries"
          height="300" 
        />
      </div>
      
      <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
        <h3 class="text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Status Kamar</h3>
        <AppChart 
          chart-type="pie"
          :options="roomChartOptions" 
          :series="roomChartSeries"
          height="300" 
        />
      </div>
    </div>
    
    <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow mb-8">
      <h3 class="text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Ringkasan Transaksi per Kategori</h3>
      <AppChart 
        chart-type="bar"
        :options="categoryChartOptions" 
        :series="categoryChartSeries"
        height="300" 
      />
    </div>
    
    <div class="bg-white dark:bg-slate-800 p-6 rounded-xl shadow">
      <h3 class="text-lg font-semibold mb-4 text-gray-700 dark:text-gray-300">Daftar Tagihan Tertunggak</h3>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
          <thead>
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Penyewa</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Kamar</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Periode</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Jumlah</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Status</th>
            </tr>
          </thead>
          <tbody class="bg-white dark:bg-slate-800 divide-y divide-gray-200 dark:divide-gray-700">
            <tr v-for="tagihan in unpaidTagihans" :key="tagihan.id" class="hover:bg-gray-50 dark:hover:bg-slate-700">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ tagihan.penyewa_nama }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ tagihan.kamar_info }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">
                {{ formatDate(tagihan.awal_periode) }} - {{ formatDate(tagihan.akhir_periode) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-200">{{ formatRupiah(tagihan.jumlah_tagihan) }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span 
                  :class="{
                    'px-2 py-1 rounded-full text-xs font-medium': true,
                    'bg-red-100 text-red-800': tagihan.status === 'belum lunas',
                    'bg-yellow-100 text-yellow-800': tagihan.status === 'nyicil'
                  }"
                >
                  {{ tagihan.status }}
                </span>
              </td>
            </tr>
            <tr v-if="unpaidTagihans.length === 0">
              <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500 dark:text-gray-400">
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
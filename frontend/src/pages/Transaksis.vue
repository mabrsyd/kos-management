<template>
  <div>
    <div class="flex flex-wrap justify-between items-center gap-4 mb-6">
      <div>
        <h2 class="text-xl sm:text-2xl font-bold text-gray-800 dark:text-gray-100">Manajemen Transaksi</h2>
        <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">Kelola transaksi keuangan</p>
      </div>
      
      <!-- Add Button -->
      <button 
        @click="showForm = true" 
        class="bg-blue-600 hover:bg-blue-700 text-white flex items-center gap-2 px-4 py-2 rounded-lg shadow-sm transition-colors duration-200"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        <span>Tambah Transaksi</span>
      </button>
    </div>
    
    <!-- Summary Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6 mb-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-green-100 dark:bg-green-900/30 p-3 mr-4">
            <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Total Pemasukan</h3>
            <p class="text-xl font-bold text-green-600 dark:text-green-400">{{ formatRupiah(totalPemasukan) }}</p>
          </div>
        </div>
      </div>
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-red-100 dark:bg-red-900/30 p-3 mr-4">
            <svg class="w-6 h-6 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Total Pengeluaran</h3>
            <p class="text-xl font-bold text-red-600 dark:text-red-400">{{ formatRupiah(totalPengeluaran) }}</p>
          </div>
        </div>
      </div>
      <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 transition-colors duration-200">
        <div class="flex items-center">
          <div class="rounded-full bg-blue-100 dark:bg-blue-900/30 p-3 mr-4">
            <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3" />
            </svg>
          </div>
          <div>
            <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">Saldo</h3>
            <p class="text-xl font-bold" :class="saldo >= 0 ? 'text-blue-600 dark:text-blue-400' : 'text-red-600 dark:text-red-400'">
              {{ formatRupiah(saldo) }}
            </p>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Filter Controls -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 p-4 md:p-6 mb-6 transition-colors duration-200">
      <h3 class="text-base font-medium text-gray-700 dark:text-gray-300 mb-4">Filter Transaksi</h3>
      <div class="flex flex-wrap gap-4 items-center">
        <div>
          <label for="filterTipe" class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Tipe</label>
          <select 
            v-model="filterTipe" 
            id="filterTipe" 
            class="border border-gray-300 dark:border-slate-600 rounded-lg px-3 py-2 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent transition-colors duration-200"
          >
            <option value="">Semua</option>
            <option value="pemasukan">Pemasukan</option>
            <option value="pengeluaran">Pengeluaran</option>
          </select>
        </div>
        
        <div>
          <label for="startDate" class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Dari Tanggal</label>
          <input 
            type="date" 
            v-model="startDate" 
            id="startDate"
            class="border border-gray-300 dark:border-slate-600 rounded-lg px-3 py-2 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent transition-colors duration-200"
          />
        </div>
        
        <div>
          <label for="endDate" class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Sampai Tanggal</label>
          <input 
            type="date" 
            v-model="endDate" 
            id="endDate"
            class="border border-gray-300 dark:border-slate-600 rounded-lg px-3 py-2 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent transition-colors duration-200"
          />
        </div>
        
        <button 
          @click="resetFilter" 
          class="self-end bg-gray-200 dark:bg-slate-700 text-gray-700 dark:text-gray-300 px-4 py-2 rounded-lg font-medium hover:bg-gray-300 dark:hover:bg-slate-600 transition-colors duration-200 flex items-center gap-2"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          Reset
        </button>
      </div>
    </div>

    <!-- Table of Transaksis -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 overflow-hidden">
      <div class="overflow-x-auto">
        <AppTable :columns="columns" :data="filteredTransaksis">
          <template #action="{ row }">
            <div class="flex justify-end gap-2">
              <button @click="editTransaksi(row)" class="inline-flex items-center justify-center p-2 rounded-full bg-blue-50 hover:bg-blue-100 dark:bg-blue-900/20 dark:hover:bg-blue-800/30 text-blue-600 dark:text-blue-400 transition-colors" title="Edit">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>
              <button @click="hapusTransaksi(row.id)" class="inline-flex items-center justify-center p-2 rounded-full bg-red-50 hover:bg-red-100 dark:bg-red-900/20 dark:hover:bg-red-800/30 text-red-600 dark:text-red-400 transition-colors" title="Hapus">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </template>
          
          <!-- Custom template for type column -->
          <template #tipe="{ row }">
            <span 
              :class="{
                'px-2 py-1 rounded-full text-xs font-medium inline-flex items-center': true,
                'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400': row.tipe === 'pemasukan',
                'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400': row.tipe === 'pengeluaran'
              }"
            >
              <span class="h-1.5 w-1.5 rounded-full mr-1" :class="{
                'bg-green-500 dark:bg-green-400': row.tipe === 'pemasukan',
                'bg-red-500 dark:bg-red-400': row.tipe === 'pengeluaran'
              }"></span>
              {{ row.tipe }}
            </span>
          </template>
          
          <!-- Format currency for jumlah -->
          <template #jumlah="{ row }">
            {{ formatRupiah(row.jumlah) }}
          </template>
        </AppTable>
      </div>
    </div>

    <!-- Modal Form -->
    <div v-if="showForm" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto border border-gray-200 dark:border-slate-700">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-gray-800 dark:text-gray-100">{{ editingId ? 'Edit' : 'Tambah' }} Transaksi</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 p-1 rounded-lg hover:bg-gray-100 dark:hover:bg-slate-700 transition-colors duration-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <!-- Form Tambah/Edit Transaksi -->
        <form @submit.prevent="submitTransaksi" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <select v-model="form.tipe" required id="tipe" class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200">
                <option value="pemasukan">Pemasukan</option>
                <option value="pengeluaran">Pengeluaran</option>
              </select>
              <label for="tipe" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Tipe</label>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input v-model="form.kategori" required id="kategori" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" />
              <label for="kategori" class="absolute left-0 top-2 text-gray-400 dark:text-gray-500 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 dark:peer-focus:text-blue-400 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500 dark:peer-valid:text-blue-400">Kategori</label>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input v-model.number="form.jumlah" type="number" min="0" required id="jumlah" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" />
              <label for="jumlah" class="absolute left-0 top-2 text-gray-400 dark:text-gray-500 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 dark:peer-focus:text-blue-400 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500 dark:peer-valid:text-blue-400">Jumlah</label>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input v-model="form.tanggal_transaksi" type="date" required id="tanggal_transaksi" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" />
              <label for="tanggal_transaksi" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Tanggal Transaksi</label>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg md:col-span-2">
            <div class="relative">
              <input v-model="form.catatan" id="catatan" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" />
              <label for="catatan" class="absolute left-0 top-2 text-gray-400 dark:text-gray-500 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 dark:peer-focus:text-blue-400 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500 dark:peer-valid:text-blue-400">Catatan</label>
            </div>
          </div>
          
          <div class="col-span-2 mt-4 flex gap-2 justify-end">
            <button type="button" @click="closeForm" class="bg-gray-200 dark:bg-slate-700 text-gray-700 dark:text-gray-300 px-4 py-2 rounded-lg font-medium hover:bg-gray-300 dark:hover:bg-slate-600 transition-colors duration-200">Batal</button>
            <button type="submit" :disabled="loadingBtn" class="bg-blue-600 text-white px-6 py-2 rounded-lg font-medium shadow hover:bg-blue-700 transition-colors duration-200 disabled:opacity-60 flex items-center gap-2">
              <svg v-if="loadingBtn" class="animate-spin w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path></svg>
              {{ editingId ? 'Update' : 'Tambah' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '../utils/api'
import { useAttrs } from 'vue'
import AppTable from '../components/ui/AppTable.vue'

const columns = [
  { key: 'tipe', label: 'Tipe' },
  { key: 'kategori', label: 'Kategori' },
  { key: 'jumlah', label: 'Jumlah' },
  { key: 'tanggal_transaksi', label: 'Tanggal' },
  { key: 'catatan', label: 'Catatan' },
]
const transaksis = ref([])
const loadingBtn = ref(false)
const showForm = ref(false)
const form = ref({
  tipe: 'pemasukan',
  kategori: '',
  jumlah: 0,
  tanggal_transaksi: '',
  catatan: '',
  user_id: null
})
const editingId = ref(null)
const attrs = useAttrs()
const setNotif = attrs.setNotif || (() => {})
const errorMessage = ref('')

// Filters
const filterTipe = ref('')
const startDate = ref('')
const endDate = ref('')

// Format currency
const formatRupiah = (value) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value)
}

// Filter transaksis based on filters
const filteredTransaksis = computed(() => {
  return transaksis.value.filter(t => {
    // Filter by type
    if (filterTipe.value && t.tipe !== filterTipe.value) {
      return false
    }
    
    // Filter by date range
    if (startDate.value && new Date(t.tanggal_transaksi) < new Date(startDate.value)) {
      return false
    }
    
    if (endDate.value && new Date(t.tanggal_transaksi) > new Date(endDate.value)) {
      return false
    }
    
    return true
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

// Reset filters
function resetFilter() {
  filterTipe.value = ''
  startDate.value = ''
  endDate.value = ''
}

// Dapatkan user_id dari local storage
const getUserId = () => {
  try {
    const userData = localStorage.getItem('user')
    if (userData) {
      const user = JSON.parse(userData)
      return user.id
    }
    return null
  } catch (error) {
    console.error('Error getting user ID:', error)
    return null
  }
}

async function fetchTransaksis() {
  try {
    const res = await api.get('/transaksis')
    transaksis.value = res.data
  } catch (error) {
    errorMessage.value = 'Gagal memuat data transaksi'
    console.error('Error fetching transaksis:', error)
  }
}

function closeForm() {
  showForm.value = false
  resetForm()
}

function resetForm() {
  editingId.value = null
  form.value = { 
    tipe: 'pemasukan', 
    kategori: '', 
    jumlah: 0, 
    tanggal_transaksi: '', 
    catatan: '', 
    user_id: null
  }
}

async function submitTransaksi() {
  loadingBtn.value = true
  
  // Set user_id dari local storage
  form.value.user_id = getUserId()
  
  try {    
    if (editingId.value) {
      await api.put(`/transaksis/${editingId.value}`, form.value)
      setNotif('Transaksi berhasil diupdate!')
    } else {
      await api.post('/transaksis', form.value)
      setNotif('Transaksi berhasil ditambah!')
    }
    closeForm()
    await fetchTransaksis()
  } catch (error) {
    console.error('Error saving transaksi:', error)
    if (error.response) {
      setNotif(`Gagal menyimpan transaksi: ${error.response.data.error || 'Terjadi kesalahan'}`)
    } else {
      setNotif('Gagal menyimpan transaksi: Tidak dapat terhubung ke server')
    }
  } finally {
    loadingBtn.value = false
  }
}

function editTransaksi(transaksi) {
  form.value = { ...transaksi }
  editingId.value = transaksi.id
  showForm.value = true
}

function batalEdit() {
  closeForm()
}

async function hapusTransaksi(id) {
  if (confirm('Yakin ingin menghapus transaksi ini?')) {
    try {
      await api.delete(`/transaksis/${id}`)
      setNotif('Transaksi berhasil dihapus!')
      await fetchTransaksis()
    } catch (error) {
      console.error('Error deleting transaksi:', error)
      if (error.response) {
        setNotif(`Gagal menghapus transaksi: ${error.response.data.error || 'Terjadi kesalahan'}`)
      } else {
        setNotif('Gagal menghapus transaksi: Tidak dapat terhubung ke server')
      }
    }
  }
}

onMounted(() => {
  fetchTransaksis()
})
</script> 
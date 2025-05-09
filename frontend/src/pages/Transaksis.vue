<template>
  <div class="p-2 max-w-8xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h2 class="text-2xl font-bold text-blue-700 dark:text-blue-300">Manajemen Transaksi</h2>
      
      <!-- Add Button -->
      <button @click="showForm = true" class="bg-blue-600 text-white p-2 rounded-full shadow hover:bg-blue-700 transition">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
      </button>
    </div>
    
    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow">
        <h3 class="text-lg font-semibold text-gray-600 dark:text-gray-300 mb-2">Total Pemasukan</h3>
        <p class="text-2xl font-bold text-green-600 dark:text-green-400">{{ formatRupiah(totalPemasukan) }}</p>
      </div>
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow">
        <h3 class="text-lg font-semibold text-gray-600 dark:text-gray-300 mb-2">Total Pengeluaran</h3>
        <p class="text-2xl font-bold text-red-600 dark:text-red-400">{{ formatRupiah(totalPengeluaran) }}</p>
      </div>
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow">
        <h3 class="text-lg font-semibold text-gray-600 dark:text-gray-300 mb-2">Saldo</h3>
        <p class="text-2xl font-bold" :class="saldo >= 0 ? 'text-blue-600 dark:text-blue-400' : 'text-red-600 dark:text-red-400'">
          {{ formatRupiah(saldo) }}
        </p>
      </div>
    </div>
    
    <!-- Filter Controls -->
    <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow mb-6 flex flex-wrap gap-4 items-center">
      <div>
        <label for="filterTipe" class="block text-sm font-medium text-gray-600 dark:text-gray-300 mb-1">Filter Tipe</label>
        <select 
          v-model="filterTipe" 
          id="filterTipe" 
          class="border border-gray-300 rounded-md px-3 py-2 focus:ring-blue-500 focus:border-blue-500 dark:bg-slate-700 dark:border-slate-600"
        >
          <option value="">Semua</option>
          <option value="pemasukan">Pemasukan</option>
          <option value="pengeluaran">Pengeluaran</option>
        </select>
      </div>
      
      <div>
        <label for="startDate" class="block text-sm font-medium text-gray-600 dark:text-gray-300 mb-1">Dari Tanggal</label>
        <input 
          type="date" 
          v-model="startDate" 
          id="startDate"
          class="border border-gray-300 rounded-md px-3 py-2 focus:ring-blue-500 focus:border-blue-500 dark:bg-slate-700 dark:border-slate-600"
        />
      </div>
      
      <div>
        <label for="endDate" class="block text-sm font-medium text-gray-600 dark:text-gray-300 mb-1">Sampai Tanggal</label>
        <input 
          type="date" 
          v-model="endDate" 
          id="endDate"
          class="border border-gray-300 rounded-md px-3 py-2 focus:ring-blue-500 focus:border-blue-500 dark:bg-slate-700 dark:border-slate-600"
        />
      </div>
      
      <button 
        @click="resetFilter" 
        class="self-end bg-gray-200 text-gray-700 px-4 py-2 rounded-lg font-medium hover:bg-gray-300 transition"
      >
        Reset Filter
      </button>
    </div>

    <!-- Table of Transaksis -->
    <AppTable :columns="columns" :data="filteredTransaksis">
      <template #action="{ row }">
        <button @click="editTransaksi(row)" class="text-blue-600 hover:bg-blue-100 p-2 rounded transition" title="Edit">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l12-12a2 2 0 00-2.828-2.828L3 17z"/></svg>
        </button>
        <button @click="hapusTransaksi(row.id)" class="text-red-500 hover:bg-red-100 p-2 rounded transition" title="Hapus">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </template>
      
      <!-- Custom template for type column -->
      <template #tipe="{ row }">
        <span 
          :class="{
            'px-2 py-1 rounded-full text-xs font-medium': true,
            'bg-green-100 text-green-800': row.tipe === 'pemasukan',
            'bg-red-100 text-red-800': row.tipe === 'pengeluaran'
          }"
        >
          {{ row.tipe }}
        </span>
      </template>
      
      <!-- Format currency for jumlah -->
      <template #jumlah="{ row }">
        {{ formatRupiah(row.jumlah) }}
      </template>
    </AppTable>

    <!-- Modal Form -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold text-blue-700 dark:text-blue-300">{{ editingId ? 'Edit' : 'Tambah' }} Transaksi</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <!-- Form Tambah/Edit Transaksi -->
        <form @submit.prevent="submitTransaksi" class="flex flex-wrap gap-6 items-end">
          <div class="relative flex-1 min-w-[180px]">
            <select v-model="form.tipe" required id="tipe" class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent">
              <option value="pemasukan">Pemasukan</option>
              <option value="pengeluaran">Pengeluaran</option>
            </select>
            <label for="tipe" class="absolute left-0 top-2 text-gray-400 text-sm">Tipe</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <input v-model="form.kategori" required id="kategori" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="kategori" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Kategori</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <input v-model.number="form.jumlah" type="number" min="0" required id="jumlah" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="jumlah" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Jumlah</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <input v-model="form.tanggal_transaksi" type="date" required id="tanggal_transaksi" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="tanggal_transaksi" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Tanggal Transaksi</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <input v-model="form.catatan" id="catatan" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="catatan" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Catatan</label>
          </div>
          <div class="flex mt-6 w-full gap-2 justify-end">
            <button type="button" @click="closeForm" class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg font-semibold hover:bg-gray-300 transition">Batal</button>
            <button type="submit" :disabled="loadingBtn" class="bg-blue-600 text-white px-6 py-2 rounded-lg font-semibold shadow hover:bg-blue-700 transition disabled:opacity-60 flex items-center gap-2">
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
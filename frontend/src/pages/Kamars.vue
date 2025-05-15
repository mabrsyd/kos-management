<template>
  <div class="p-4 md:p-6 max-w-8xl mx-auto">
    <div class="flex flex-col md:flex-row md:justify-between md:items-center mb-8 gap-4">
      <h2 class="text-2xl md:text-3xl font-bold text-blue-700 dark:text-blue-300 tracking-tight">Manajemen Kamar</h2>
      
      <!-- Add Button -->
      <button 
        @click="showForm = true" 
        class="flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg shadow-sm hover:bg-blue-700 hover:shadow-md transition-all duration-200 self-start md:self-auto"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        <span>Tambah Kamar</span>
      </button>
    </div>

    <!-- Filter Controls -->
    <div class="bg-white dark:bg-slate-800 rounded-xl p-4 mb-6 shadow-sm border border-gray-100 dark:border-slate-700">
      <h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-3">Filter & Urutan</h3>
      <div class="flex flex-wrap gap-4 items-center">
        <div class="flex items-center space-x-2">
          <label for="filter-tipe" class="text-sm font-medium text-gray-700 dark:text-gray-300">Tipe Kamar:</label>
          <select 
            id="filter-tipe" 
            v-model="filterTipe" 
            class="rounded-lg border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-700 dark:border-slate-600 dark:text-white text-sm py-2 px-3 transition-all"
          >
            <option value="">Semua Tipe</option>
            <option value="murah">Murah</option>
            <option value="menengah">Menengah</option>
            <option value="mahal">Mahal</option>
          </select>
        </div>
        
        <div class="flex items-center space-x-2">
          <label for="sort-order" class="text-sm font-medium text-gray-700 dark:text-gray-300">Urutan:</label>
          <div class="relative">
            <select 
              id="sort-order" 
              v-model="sortOrder" 
              class="pl-9 rounded-lg border border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-700 dark:border-slate-600 dark:text-white text-sm py-2 pr-3 transition-all appearance-none"
            >
              <option value="asc">No. Kamar (Naik)</option>
              <option value="desc">No. Kamar (Turun)</option>
            </select>
            <div class="absolute left-3 top-1/2 transform -translate-y-1/2 pointer-events-none text-gray-500 dark:text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path v-if="sortOrder === 'asc'" stroke-linecap="round" stroke-linejoin="round" d="M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12" />
                <path v-else stroke-linecap="round" stroke-linejoin="round" d="M3 4h13M3 8h9m-9 4h9m5-4v12m0 0l-4-4m4 4l4-4" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Stats Summary -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-blue-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-600 dark:text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Total Kamar</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ kamars.length }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-green-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-green-600 dark:text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Kamar Kosong</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ kamars.filter(k => k.status === 'kosong').length }}
            </p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-red-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-red-600 dark:text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Kamar Terisi</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ kamars.filter(k => k.status === 'terisi').length }}
            </p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-purple-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-purple-600 dark:text-purple-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Pendapatan</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ formatRupiah(kamars.filter(k => k.status === 'terisi').reduce((sum, kamar) => sum + (kamar.harga || 0), 0)) }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Table of Kamars -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm overflow-hidden border border-gray-100 dark:border-slate-700">
      <AppTable :columns="columns" :data="filteredAndSortedKamars">
        <template #nomer="{ row }">
          <span class="font-medium">{{ row.nomer }}</span>
        </template>
        
        <template #tipe="{ row }">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium capitalize" 
                :class="{
                  'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300': row.tipe === 'murah',
                  'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300': row.tipe === 'menengah',
                  'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-300': row.tipe === 'mahal'
                }">
            {{ row.tipe }}
          </span>
        </template>
        
        <template #harga="{ row }">
          <span class="font-medium">{{ formatRupiah(row.harga) }}</span>
        </template>
        
        <template #status="{ row }">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium capitalize" 
                :class="{
                  'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300': row.status === 'kosong',
                  'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300': row.status === 'terisi'
                }">
            {{ row.status }}
          </span>
        </template>
        
        <template #action="{ row }">
          <div class="flex justify-end gap-2">
            <button @click="editKamar(row)" class="inline-flex items-center justify-center p-2 rounded-full bg-blue-50 hover:bg-blue-100 dark:bg-blue-900/20 dark:hover:bg-blue-800/30 text-blue-600 dark:text-blue-400 transition-colors" title="Edit">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
            </button>
            <button @click="hapusKamar(row.id)" class="inline-flex items-center justify-center p-2 rounded-full bg-red-50 hover:bg-red-100 dark:bg-red-900/20 dark:hover:bg-red-800/30 text-red-600 dark:text-red-400 transition-colors" title="Hapus">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </template>
      </AppTable>
      
      <!-- Empty state -->
      <div v-if="filteredAndSortedKamars.length === 0" class="py-12 flex flex-col items-center justify-center text-center px-4">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white">Tidak ada kamar ditemukan</h3>
        <p class="mt-1 text-gray-500 dark:text-gray-400 max-w-md">
          Tidak ada kamar yang cocok dengan filter yang dipilih. Coba ubah filter atau tambahkan kamar baru.
        </p>
        <button 
          @click="showForm = true" 
          class="mt-4 flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg shadow-sm hover:bg-blue-700 hover:shadow-md transition-all duration-200"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          <span>Tambah Kamar</span>
        </button>
      </div>
    </div>

    <!-- Modal Form with Auto-Save -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-70 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div class="bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto border border-gray-200 dark:border-slate-700">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-blue-700 dark:text-blue-300">{{ editingId ? 'Edit' : 'Tambah' }} Kamar</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 bg-gray-100 dark:bg-slate-700 p-2 rounded-lg transition-all duration-200">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <div class="flex items-center text-sm text-gray-500 mb-6 bg-blue-50 dark:bg-slate-700 p-3 rounded-lg">
          <svg v-if="autoSaving" class="animate-spin h-4 w-4 text-blue-500 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else-if="lastSaved" class="h-4 w-4 text-green-500 mr-2" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
          </svg>
          <span>{{ autoSaving ? 'Menyimpan...' : lastSaved ? `Tersimpan pada ${lastSaved}` : 'Otomatis tersimpan saat diubah' }}</span>
        </div>
        
        <!-- Form fields -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="nomer" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Nomor Kamar</label>
              <input 
                v-model="form.nomer" 
                required 
                id="nomer" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all"
                @input="debouncedAutoSave"
              />
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="tipe" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Tipe Kamar</label>
              <select 
                v-model="form.tipe" 
                required 
                id="tipe" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all"
                @change="debouncedAutoSave"
              >
                <option value="murah">Murah</option>
                <option value="menengah">Menengah</option>
                <option value="mahal">Mahal</option>
              </select>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="harga" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Harga Kamar</label>
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <span class="text-gray-500 dark:text-gray-400">Rp</span>
                </div>
                <input 
                  v-model.number="form.harga" 
                  type="number" 
                  min="0" 
                  required 
                  id="harga" 
                  class="pl-10 block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all"
                  @input="debouncedAutoSave"
                />
              </div>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="status" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Status Kamar</label>
              <select 
                v-model="form.status" 
                required 
                id="status" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all"
                @change="debouncedAutoSave"
              >
                <option value="kosong">Kosong</option>
                <option value="terisi">Terisi</option>
              </select>
            </div>
          </div>
        </div>
        
        <div class="mt-8 flex justify-end">
          <button type="button" @click="closeForm" class="bg-gray-200 text-gray-700 px-5 py-2 rounded-lg font-medium hover:bg-gray-300 transition-all duration-200 dark:bg-slate-700 dark:text-gray-300 dark:hover:bg-slate-600">Tutup</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import api from '../utils/api'
import { useAttrs } from 'vue'
import AppTable from '../components/ui/AppTable.vue'

const columns = [
  { key: 'nomer', label: 'Nomor' },
  { key: 'tipe', label: 'Tipe' },
  { key: 'harga', label: 'Harga' },
  { key: 'status', label: 'Status' }
]
const kamars = ref([])
const loadingBtn = ref(false)
const showForm = ref(false)
const form = ref({ nomer: '', tipe: 'murah', harga: 0, status: 'kosong', user_id: null })
const editingId = ref(null)
const attrs = useAttrs()
const setNotif = attrs.setNotif || (() => {})
const errorMessage = ref('')
const autoSaving = ref(false)
const lastSaved = ref(null)
const autoSaveTimer = ref(null)

// Filtering and sorting
const filterTipe = ref('')
const sortOrder = ref('asc')

// Format currency
const formatRupiah = (value) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value)
}

// Computed property for filtered and sorted kamars
const filteredAndSortedKamars = computed(() => {
  // Filter by tipe if filterTipe is set
  let filtered = filterTipe.value 
    ? kamars.value.filter(kamar => kamar.tipe === filterTipe.value)
    : [...kamars.value]
  
  // Sort by nomer (as number)
  filtered.sort((a, b) => {
    const aNum = Number(a.nomer)
    const bNum = Number(b.nomer)
    
    // Handle non-numeric values or NaN
    if (isNaN(aNum) && isNaN(bNum)) return 0
    if (isNaN(aNum)) return 1
    if (isNaN(bNum)) return -1
    
    // Normal numeric comparison
    return sortOrder.value === 'asc' ? aNum - bNum : bNum - aNum
  })
  
  return filtered
})

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

// Format time for last saved
const formatTime = () => {
  const now = new Date()
  return `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}:${now.getSeconds().toString().padStart(2, '0')}`
}

async function fetchKamars() {
  try {
    const res = await api.get('/kamars')
    kamars.value = res.data
  } catch (error) {
    errorMessage.value = 'Gagal memuat data kamar'
    console.error('Error fetching kamars:', error)
  }
}

function debouncedAutoSave() {
  if (autoSaveTimer.value) {
    clearTimeout(autoSaveTimer.value)
  }
  
  autoSaveTimer.value = setTimeout(() => {
    autoSave()
  }, 1000) // 1 second debounce
}

async function autoSave() {
  // Don't save if form is empty
  if (!form.value.nomer.trim()) return
  
  autoSaving.value = true
  form.value.user_id = getUserId()
  
  try {
    if (editingId.value) {
      await api.put(`/kamars/${editingId.value}`, form.value)
    } else {
      const response = await api.post('/kamars', form.value)
      // Update editing ID so next save is an update, not a new creation
      editingId.value = response.data.id || response.data
      setNotif('Kamar berhasil ditambah!')
    }
    
    lastSaved.value = formatTime()
    await fetchKamars()
  } catch (error) {
    console.error('Error auto-saving kamar:', error)
    if (error.response) {
      setNotif(`Gagal menyimpan kamar: ${error.response.data.error || 'Terjadi kesalahan'}`)
    } else {
      setNotif('Gagal menyimpan kamar: Tidak dapat terhubung ke server')
    }
  } finally {
    autoSaving.value = false
  }
}

function closeForm() {
  showForm.value = false
  resetForm()
}

function resetForm() {
  editingId.value = null
  form.value = { nomer: '', tipe: 'murah', harga: 0, status: 'kosong', user_id: null }
  lastSaved.value = null
}

function editKamar(kamar) {
  form.value = { ...kamar }
  editingId.value = kamar.id
  showForm.value = true
  lastSaved.value = null
}

function batalEdit() {
  closeForm()
}

async function hapusKamar(id) {
  if (confirm('Yakin ingin menghapus kamar ini?')) {
    try {
      await api.delete(`/kamars/${id}`)
      setNotif('Kamar berhasil dihapus!')
      await fetchKamars()
    } catch (error) {
      console.error('Error deleting kamar:', error)
      if (error.response) {
        setNotif(`Gagal menghapus kamar: ${error.response.data.error || 'Terjadi kesalahan'}`)
      } else {
        setNotif('Gagal menghapus kamar: Tidak dapat terhubung ke server')
      }
    }
  }
}

onMounted(() => {
  fetchKamars()
})
</script>

<style scoped>
.dark select {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%239ca3af' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e");
  background-position: right 0.5rem center;
  background-repeat: no-repeat;
  background-size: 1.5em 1.5em;
  padding-right: 2.5rem;
}

/* For webkit browsers like Chrome/Safari */
input[type=number]::-webkit-inner-spin-button, 
input[type=number]::-webkit-outer-spin-button { 
  -webkit-appearance: none; 
  margin: 0; 
}

/* For Firefox */
input[type=number] {
  -moz-appearance: textfield;
}
</style> 
<template>
  <div class="p-4 md:p-6 max-w-8xl mx-auto">
    <div class="flex flex-col md:flex-row md:justify-between md:items-center mb-8 gap-4">
      <h2 class="text-2xl md:text-3xl font-bold text-blue-700 dark:text-blue-300 tracking-tight">Manajemen Penyewa</h2>
      
      <!-- Add Button -->
      <button @click="showForm = true" class="flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg shadow-sm hover:bg-blue-700 hover:shadow-md transition-all duration-200 self-start md:self-auto">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        <span>Tambah Penyewa</span>
      </button>
    </div>

    <!-- Stats Summary -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-blue-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-600 dark:text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Total Penyewa</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ penyewas.length }}</p>
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
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Penyewa Aktif</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ penyewas.filter(p => p.status === 'aktif').length }}
            </p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-yellow-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-yellow-600 dark:text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Masa Kontrak Habis</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ penyewasMasaHabis.length }}
            </p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center">
          <div class="p-3 rounded-lg bg-purple-50 dark:bg-slate-700">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-purple-600 dark:text-purple-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Status Pelajar</p>
            <p class="text-lg font-bold text-gray-900 dark:text-white">
              <span class="inline-block mr-2">Mahasiswa: {{ penyewas.filter(p => p.status_pelajar === 'mahasiswa').length }}</span>
              <span class="inline-block">Kerja: {{ penyewas.filter(p => p.status_pelajar === 'kerja').length }}</span>
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Table of Penyewas -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm overflow-hidden border border-gray-100 dark:border-slate-700">
      <AppTable :columns="columns" :data="penyewas">
        <template #nama="{ row }">
          <span class="font-medium">{{ row.nama }}</span>
        </template>
        
        <template #status_pelajar="{ row }">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium capitalize" 
                :class="{
                  'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300': row.status_pelajar === 'mahasiswa',
                  'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-300': row.status_pelajar === 'kerja'
                }">
            {{ row.status_pelajar }}
          </span>
        </template>
        
        <template #tanggal_mulai="{ row }">
          <span :class="{'text-gray-900 dark:text-white': true}">
            {{ formatDate(row.tanggal_mulai) }}
          </span>
        </template>
        
        <template #tanggal_selesai="{ row }">
          <span :class="{
            'text-red-600 dark:text-red-400': isNearlyExpired(row.tanggal_selesai),
            'text-gray-900 dark:text-white': !isNearlyExpired(row.tanggal_selesai)
          }">
            {{ formatDate(row.tanggal_selesai) }}
          </span>
        </template>
        
        <template #status="{ row }">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium capitalize" 
                :class="{
                  'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300': row.status === 'aktif',
                  'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300': row.status === 'keluar'
                }">
            {{ row.status }}
          </span>
        </template>
        
        <template #id_kamar="{ row }">
          <div class="flex items-center">
            <svg class="h-4 w-4 text-blue-500 mr-1" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
            </svg>
            <span class="text-sm font-medium">{{ getKamarInfo(row.id_kamar) }}</span>
          </div>
        </template>
        
        <template #action="{ row }">
          <div class="flex justify-end gap-2">
            <button @click="editPenyewa(row)" class="inline-flex items-center justify-center p-2 rounded-full bg-blue-50 hover:bg-blue-100 dark:bg-blue-900/20 dark:hover:bg-blue-800/30 text-blue-600 dark:text-blue-400 transition-colors" title="Edit">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
            </button>
            <button @click="hapusPenyewa(row.id)" class="inline-flex items-center justify-center p-2 rounded-full bg-red-50 hover:bg-red-100 dark:bg-red-900/20 dark:hover:bg-red-800/30 text-red-600 dark:text-red-400 transition-colors" title="Hapus">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </template>
      </AppTable>
      
      <!-- Empty state -->
      <div v-if="penyewas.length === 0" class="py-12 flex flex-col items-center justify-center text-center px-4">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white">Belum ada penyewa</h3>
        <p class="mt-1 text-gray-500 dark:text-gray-400 max-w-md">
          Anda belum memiliki data penyewa. Tambahkan penyewa baru untuk mulai mengelola kos Anda.
        </p>
        <button 
          @click="showForm = true" 
          class="mt-4 flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg shadow-sm hover:bg-blue-700 hover:shadow-md transition-all duration-200"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          <span>Tambah Penyewa</span>
        </button>
      </div>
    </div>

    <!-- Modal Form -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-70 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div class="bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto border border-gray-200 dark:border-slate-700">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-blue-700 dark:text-blue-300">{{ editingId ? 'Edit' : 'Tambah' }} Penyewa</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 bg-gray-100 dark:bg-slate-700 p-2 rounded-lg transition-all duration-200">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <!-- Form Tambah/Edit Penyewa -->
        <form @submit.prevent="submitPenyewa" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="nama" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Nama Penyewa</label>
              <input 
                v-model="form.nama" 
                required 
                id="nama" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all" 
              />
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="status_pelajar" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Status Pelajar</label>
              <select 
                v-model="form.status_pelajar" 
                required 
                id="status_pelajar" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all"
              >
                <option value="mahasiswa">Mahasiswa</option>
                <option value="kerja">Kerja</option>
              </select>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="tanggal_mulai" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Tanggal Mulai</label>
              <input 
                v-model="form.tanggal_mulai" 
                type="date" 
                required 
                id="tanggal_mulai" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all" 
              />
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="tanggal_selesai" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Tanggal Selesai</label>
              <input 
                v-model="form.tanggal_selesai" 
                type="date" 
                required 
                id="tanggal_selesai" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all" 
              />
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="status" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Status Penyewa</label>
              <select 
                v-model="form.status" 
                required 
                id="status" 
                class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all"
              >
                <option value="aktif">Aktif</option>
                <option value="keluar">Keluar</option>
              </select>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700 p-5 rounded-xl">
            <div class="relative">
              <label for="id_kamar" class="block text-xs font-medium text-blue-600 dark:text-blue-400 mb-1">Pilih Kamar</label>
              <div class="relative">
                <select 
                  v-model="form.id_kamar" 
                  required 
                  id="id_kamar" 
                  class="block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:ring-opacity-20 dark:bg-slate-800 dark:border-slate-600 dark:text-white transition-all max-h-40 overflow-y-auto"
                  size="5"
                >
                  <option v-for="kamar in availableKamars" :key="kamar.id" :value="kamar.id" class="p-2 hover:bg-blue-50 dark:hover:bg-slate-700 transition-all">
                    Kamar {{ kamar.nomer }} - {{ formatRupiah(kamar.harga) }}
                  </option>
                </select>
                <p class="mt-2 text-xs flex items-center text-gray-500 dark:text-gray-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  Scroll untuk melihat semua kamar
                </p>
              </div>
            </div>
          </div>
          
          <div class="col-span-2 mt-6 flex gap-3 justify-end">
            <button type="button" @click="closeForm" class="bg-gray-200 text-gray-700 px-5 py-2 rounded-lg font-medium hover:bg-gray-300 transition-all duration-200 dark:bg-slate-700 dark:text-gray-300 dark:hover:bg-slate-600">Batal</button>
            <button type="submit" :disabled="loadingBtn" class="bg-blue-600 text-white px-6 py-2 rounded-lg font-medium shadow-sm hover:bg-blue-700 hover:shadow-md transition-all duration-200 disabled:opacity-60 flex items-center gap-2">
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
  { key: 'nama', label: 'Nama' },
  { key: 'status_pelajar', label: 'Status Pelajar' },
  { key: 'tanggal_mulai', label: 'Tanggal Mulai' },
  { key: 'tanggal_selesai', label: 'Tanggal Selesai' },
  { key: 'status', label: 'Status' },
  { key: 'id_kamar', label: 'Kamar' }
]
const penyewas = ref([])
const loadingBtn = ref(false)
const showForm = ref(false)
const kamars = ref([])
const form = ref({
  nama: '',
  status_pelajar: 'mahasiswa',
  tanggal_mulai: '',
  tanggal_selesai: '',
  status: 'aktif',
  id_kamar: '',
  user_id: null
})
const editingId = ref(null)
const attrs = useAttrs()
const setNotif = attrs.setNotif || (() => {})
const errorMessage = ref('')
const previousKamarId = ref(null)

// Computed property to get available kamars (only kosong)
const availableKamars = computed(() => {
  return kamars.value.filter(k => k.status === 'kosong')
})

// Computed property for penyewas whose contracts are nearing expiry (within 30 days)
const penyewasMasaHabis = computed(() => {
  const today = new Date()
  const thirtyDaysLater = new Date()
  thirtyDaysLater.setDate(today.getDate() + 30)
  
  return penyewas.value.filter(p => {
    if (p.status !== 'aktif') return false
    const tanggalSelesai = new Date(p.tanggal_selesai)
    return tanggalSelesai <= thirtyDaysLater && tanggalSelesai >= today
  })
})

// Check if a date is nearly expired (within 30 days)
const isNearlyExpired = (dateStr) => {
  const today = new Date()
  const thirtyDaysLater = new Date()
  thirtyDaysLater.setDate(today.getDate() + 30)
  
  const date = new Date(dateStr)
  return date <= thirtyDaysLater && date >= today
}

// Format date for display
const formatDate = (dateStr) => {
  const options = { year: 'numeric', month: 'short', day: 'numeric' }
  return new Date(dateStr).toLocaleDateString('id-ID', options)
}

// Format currency
const formatRupiah = (value) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value)
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

// Get kamar info for display in table
const getKamarInfo = (kamarId) => {
  const kamar = kamars.value.find(k => k.id === kamarId)
  return kamar ? `Kamar ${kamar.nomer} - ${formatRupiah(kamar.harga)}` : kamarId
}

async function fetchPenyewas() {
  try {
    const res = await api.get('/penyewas')
    penyewas.value = res.data
  } catch (error) {
    errorMessage.value = 'Gagal memuat data penyewa'
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

async function updateKamarStatus(kamarId, status) {
  try {
    const kamar = kamars.value.find(k => k.id === kamarId)
    if (kamar) {
      await api.put(`/kamars/${kamarId}`, { ...kamar, status })
      await fetchKamars() // Refresh the kamars list
    }
  } catch (error) {
    console.error('Error updating kamar status:', error)
  }
}

function closeForm() {
  showForm.value = false
  resetForm()
}

function resetForm() {
  editingId.value = null
  previousKamarId.value = null
  form.value = { 
    nama: '', 
    status_pelajar: 'mahasiswa', 
    tanggal_mulai: '', 
    tanggal_selesai: '', 
    status: 'aktif', 
    id_kamar: '',
    user_id: null
  }
}

async function submitPenyewa() {
  loadingBtn.value = true
  
  // Set user_id dari local storage
  form.value.user_id = getUserId()
  
  try {
    if (editingId.value) {
      // If room changed, update status of both old and new rooms
      if (previousKamarId.value && previousKamarId.value !== form.value.id_kamar) {
        // Free up the old room
        await updateKamarStatus(previousKamarId.value, 'kosong')
        // Mark new room as occupied
        await updateKamarStatus(form.value.id_kamar, 'terisi')
      } 
      // If penyewa status changed to 'keluar', free up the room
      else if (form.value.status === 'keluar') {
        await updateKamarStatus(form.value.id_kamar, 'kosong')
      }
      
      await api.put(`/penyewas/${editingId.value}`, form.value)
      setNotif('Penyewa berhasil diupdate!')
    } else {
      // For new penyewa, mark the room as occupied
      await updateKamarStatus(form.value.id_kamar, 'terisi')
      
      await api.post('/penyewas', form.value)
      setNotif('Penyewa berhasil ditambah!')
    }
    closeForm()
    await fetchPenyewas()
  } catch (error) {
    console.error('Error saving penyewa:', error)
    if (error.response) {
      setNotif(`Gagal menyimpan penyewa: ${error.response.data.error || 'Terjadi kesalahan'}`)
    } else {
      setNotif('Gagal menyimpan penyewa: Tidak dapat terhubung ke server')
    }
  } finally {
    loadingBtn.value = false
  }
}

function editPenyewa(penyewa) {
  form.value = { ...penyewa }
  editingId.value = penyewa.id
  previousKamarId.value = penyewa.id_kamar
  showForm.value = true
}

function batalEdit() {
  closeForm()
}

async function hapusPenyewa(id) {
  if (confirm('Yakin ingin menghapus penyewa ini?')) {
    try {
      // Get penyewa info first to free up the room
      const penyewa = penyewas.value.find(p => p.id === id)
      if (penyewa && penyewa.status === 'aktif') {
        await updateKamarStatus(penyewa.id_kamar, 'kosong')
      }
      
      await api.delete(`/penyewas/${id}`)
      setNotif('Penyewa berhasil dihapus!')
      await fetchPenyewas()
    } catch (error) {
      console.error('Error deleting penyewa:', error)
      if (error.response) {
        setNotif(`Gagal menghapus penyewa: ${error.response.data.error || 'Terjadi kesalahan'}`)
      } else {
        setNotif('Gagal menghapus penyewa: Tidak dapat terhubung ke server')
      }
    }
  }
}

onMounted(() => {
  fetchPenyewas()
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

select[size] {
  background-image: none;
}

/* Custom styling for select options */
select[size] option {
  padding: 8px;
  margin: 2px 0;
  border-radius: 4px;
}

/* Remove default styling for date inputs */
input[type="date"] {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
}

/* Custom styling for date inputs */
input[type="date"]::-webkit-calendar-picker-indicator {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='%236B7280'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z'/%3E%3C/svg%3E");
  cursor: pointer;
}

.dark input[type="date"]::-webkit-calendar-picker-indicator {
  filter: invert(0.8);
}
</style>

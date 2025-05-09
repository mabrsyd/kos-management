<template>
  <div class="p-2 max-w-8xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h2 class="text-2xl font-bold text-blue-700 dark:text-blue-300">Manajemen Kamar</h2>
      
      <!-- Add Button -->
      <button @click="showForm = true" class="bg-blue-600 text-white p-2 rounded-full shadow hover:bg-blue-700 transition">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
      </button>
    </div>

    <!-- Table of Kamars -->
    <AppTable :columns="columns" :data="kamars">
      <template #action="{ row }">
        <button @click="editKamar(row)" class="text-blue-600 hover:bg-blue-100 p-2 rounded transition" title="Edit">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l12-12a2 2 0 00-2.828-2.828L3 17z"/></svg>
        </button>
        <button @click="hapusKamar(row.id)" class="text-red-500 hover:bg-red-100 p-2 rounded transition" title="Hapus">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </template>
    </AppTable>

    <!-- Modal Form with Auto-Save -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-blue-700 dark:text-blue-300">{{ editingId ? 'Edit' : 'Tambah' }} Kamar</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <div class="text-sm text-gray-500 mb-4">
          <div class="flex items-center">
            <svg v-if="autoSaving" class="animate-spin h-4 w-4 text-blue-500 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else-if="lastSaved" class="h-4 w-4 text-green-500 mr-2" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
            </svg>
            <span>{{ autoSaving ? 'Menyimpan...' : lastSaved ? `Tersimpan pada ${lastSaved}` : 'Otomatis tersimpan saat diubah' }}</span>
          </div>
        </div>
        
        <!-- Form fields -->
        <div class="flex flex-wrap gap-6 items-end">
          <div class="relative flex-1 min-w-[180px]">
            <input 
              v-model="form.nomer" 
              required 
              id="nomer" 
              class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" 
              @input="debouncedAutoSave"
            />
            <label for="nomer" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Nomor</label>
          </div>
          
          <div class="relative flex-1 min-w-[180px]">
            <select 
              v-model="form.tipe" 
              required 
              id="tipe" 
              class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent"
              @change="debouncedAutoSave"
            >
              <option value="murah">Murah</option>
              <option value="menengah">Menengah</option>
            </select>
            <label for="tipe" class="absolute left-0 top-2 text-gray-400 text-sm">Tipe</label>
          </div>
          
          <div class="relative flex-1 min-w-[180px]">
            <input 
              v-model.number="form.harga" 
              type="number" 
              min="0" 
              required 
              id="harga" 
              class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition"
              @input="debouncedAutoSave"
            />
            <label for="harga" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Harga</label>
          </div>
          
          <div class="relative flex-1 min-w-[180px]">
            <select 
              v-model="form.status" 
              required 
              id="status" 
              class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent"
              @change="debouncedAutoSave"
            >
              <option value="kosong">Kosong</option>
              <option value="terisi">Terisi</option>
            </select>
            <label for="status" class="absolute left-0 top-2 text-gray-400 text-sm">Status</label>
          </div>
        </div>
        
        <div class="mt-8 flex justify-end">
          <button type="button" @click="closeForm" class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg font-semibold hover:bg-gray-300 transition">Tutup</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import api from '../utils/api'
import { useAttrs } from 'vue'
import AppTable from '../components/ui/AppTable.vue'

const columns = [
  { key: 'nomer', label: 'Nomor' },
  { key: 'tipe', label: 'Tipe' },
  { key: 'harga', label: 'Harga' },
  { key: 'status', label: 'Status' },
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
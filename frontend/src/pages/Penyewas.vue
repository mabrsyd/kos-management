<template>
  <div class="p-2 max-w-8xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h2 class="text-2xl font-bold text-blue-700 dark:text-blue-300">Manajemen Penyewa</h2>
      
      <!-- Add Button -->
      <button @click="showForm = true" class="bg-blue-600 text-white p-2 rounded-full shadow hover:bg-blue-700 transition">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
      </button>
    </div>

    <!-- Table of Penyewas -->
    <AppTable :columns="columns" :data="penyewas">
      <template #action="{ row }">
        <button @click="editPenyewa(row)" class="text-blue-600 hover:bg-blue-100 p-2 rounded transition" title="Edit">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l12-12a2 2 0 00-2.828-2.828L3 17z"/></svg>
        </button>
        <button @click="hapusPenyewa(row.id)" class="text-red-500 hover:bg-red-100 p-2 rounded transition" title="Hapus">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </template>
      
      <template #id_kamar="{ row }">
        {{ getKamarInfo(row.id_kamar) }}
      </template>
    </AppTable>

    <!-- Modal Form -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold text-blue-700 dark:text-blue-300">{{ editingId ? 'Edit' : 'Tambah' }} Penyewa</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <!-- Form Tambah/Edit Penyewa -->
        <form @submit.prevent="submitPenyewa" class="flex flex-wrap gap-6 items-end">
          <div class="relative flex-1 min-w-[180px]">
            <input v-model="form.nama" required id="nama" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="nama" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Nama</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <select v-model="form.status_pelajar" required id="status_pelajar" class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent">
              <option value="mahasiswa">Mahasiswa</option>
              <option value="kerja">Kerja</option>
            </select>
            <label for="status_pelajar" class="absolute left-0 top-2 text-gray-400 text-sm">Status Pelajar</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <input v-model="form.tanggal_mulai" type="date" required id="tanggal_mulai" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="tanggal_mulai" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Tanggal Mulai</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <input v-model="form.tanggal_selesai" type="date" required id="tanggal_selesai" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="tanggal_selesai" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Tanggal Selesai</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <select v-model="form.status" required id="status" class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent">
              <option value="aktif">Aktif</option>
              <option value="keluar">Keluar</option>
            </select>
            <label for="status" class="absolute left-0 top-2 text-gray-400 text-sm">Status</label>
          </div>
          <div class="relative flex-1 min-w-[180px]">
            <select 
              v-model="form.id_kamar" 
              required 
              id="id_kamar" 
              class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent"
            >
              <option v-for="kamar in availableKamars" :key="kamar.id" :value="kamar.id">
                Kamar {{ kamar.nomer }} - {{ formatRupiah(kamar.harga) }}
              </option>
            </select>
            <label for="id_kamar" class="absolute left-0 top-2 text-gray-400 text-sm">Kamar</label>
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
  { key: 'nama', label: 'Nama' },
  { key: 'status_pelajar', label: 'Status Pelajar' },
  { key: 'tanggal_mulai', label: 'Tanggal Mulai' },
  { key: 'tanggal_selesai', label: 'Tanggal Selesai' },
  { key: 'status', label: 'Status' },
  { key: 'id_kamar', label: 'Kamar' },
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
  // Include currently assigned room (for edit mode) plus all available rooms
  return kamars.value.filter(k => k.status === 'kosong' || k.id === form.value.id_kamar)
})

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

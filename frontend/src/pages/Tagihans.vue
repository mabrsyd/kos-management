<template>
  <div class="p-2 max-w-8xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h2 class="text-2xl font-bold text-blue-700 dark:text-blue-300">Manajemen Tagihan</h2>
      
      <!-- Add Button -->
      <button @click="showForm = true" class="bg-blue-600 text-white p-2 rounded-full shadow hover:bg-blue-700 transition">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
      </button>
    </div>

    <!-- Table of Tagihans -->
    <AppTable :columns="columns" :data="tagihanDisplay">
      <template #action="{ row }">
        <button @click="editTagihan(row)" class="text-blue-600 hover:bg-blue-100 p-2 rounded transition" title="Edit Status">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l12-12a2 2 0 00-2.828-2.828L3 17z"/></svg>
        </button>
        <button @click="hapusTagihan(row.id)" class="text-red-500 hover:bg-red-100 p-2 rounded transition" title="Hapus">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </template>

      <!-- Custom template for status column -->
      <template #status="{ row }">
        <span 
          :class="{
            'px-2 py-1 rounded-full text-xs font-medium': true,
            'bg-green-100 text-green-800': row.status === 'lunas',
            'bg-red-100 text-red-800': row.status === 'belum lunas',
            'bg-yellow-100 text-yellow-800': row.status === 'nyicil'
          }"
        >
          {{ row.status }}
        </span>
      </template>
    </AppTable>

    <!-- Modal Form for Edit Status -->
    <div v-if="showForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-bold text-blue-700 dark:text-blue-300">{{ editingId ? 'Edit Status Tagihan' : 'Tambah Tagihan' }}</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <!-- Form Tambah/Edit Tagihan -->
        <form @submit.prevent="submitTagihan" class="flex flex-wrap gap-6 items-end">
          <!-- Penyewa dropdown - show only when adding new tagihan -->
          <div v-if="!editingId" class="relative flex-1 min-w-[180px]">
            <select 
              v-model="form.id_penyewa" 
              required 
              id="id_penyewa" 
              class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent"
              @change="onPenyewaChange"
            >
              <option value="" disabled>Pilih Penyewa</option>
              <option v-for="penyewa in activePenyewas" :key="penyewa.id" :value="penyewa.id">
                {{ penyewa.nama }} - Kamar {{ getKamarInfo(penyewa.id_kamar) }}
              </option>
            </select>
            <label for="id_penyewa" class="absolute left-0 top-2 text-gray-400 text-sm">Penyewa</label>
          </div>
          
          <!-- For editing, display penyewa name but don't allow changing -->
          <div v-if="editingId" class="relative flex-1 min-w-[180px]">
            <p class="py-2 border-b-2 border-blue-200">{{ getPenyewaName(form.id_penyewa) }}</p>
            <label class="absolute left-0 -top-5 text-xs text-blue-500">Penyewa</label>
          </div>
          
          <div v-if="!editingId" class="relative flex-1 min-w-[180px]">
            <input v-model="form.awal_periode" type="date" required id="awal_periode" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="awal_periode" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Awal Periode</label>
          </div>
          
          <div v-if="!editingId" class="relative flex-1 min-w-[180px]">
            <input v-model="form.akhir_periode" type="date" required id="akhir_periode" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition" />
            <label for="akhir_periode" class="absolute left-0 top-2 text-gray-400 text-sm transition-all peer-focus:-top-5 peer-focus:text-xs peer-focus:text-blue-500 peer-valid:-top-5 peer-valid:text-xs peer-valid:text-blue-500">Akhir Periode</label>
          </div>
          
          <div v-if="!editingId" class="relative flex-1 min-w-[180px]">
            <input :value="displayHarga" disabled id="jumlah_tagihan" class="peer border-b-2 border-gray-200 outline-none w-full py-2 bg-transparent transition text-gray-500" />
            <label for="jumlah_tagihan" class="absolute left-0 -top-5 text-xs text-blue-500">Jumlah Tagihan</label>
          </div>
          
          <div class="relative flex-1 min-w-[180px]">
            <select v-model="form.status" required id="status" class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent">
              <option value="lunas">Lunas</option>
              <option value="belum lunas">Belum Lunas</option>
              <option value="nyicil">Nyicil</option>
            </select>
            <label for="status" class="absolute left-0 top-2 text-gray-400 text-sm">Status</label>
          </div>
          
          <div class="flex mt-6 w-full gap-2 justify-end">
            <button type="button" @click="closeForm" class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg font-semibold hover:bg-gray-300 transition">Batal</button>
            <button type="submit" :disabled="loadingBtn" class="bg-blue-600 text-white px-6 py-2 rounded-lg font-semibold shadow hover:bg-blue-700 transition disabled:opacity-60 flex items-center gap-2">
              <svg v-if="loadingBtn" class="animate-spin w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path></svg>
              {{ editingId ? 'Update Status' : 'Tambah' }}
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
  { key: 'penyewa_nama', label: 'Penyewa' },
  { key: 'kamar_info', label: 'Kamar' }, 
  { key: 'awal_periode', label: 'Awal Periode' },
  { key: 'akhir_periode', label: 'Akhir Periode' },
  { key: 'jumlah_tagihan', label: 'Jumlah Tagihan' },
  { key: 'status', label: 'Status' },
]
const tagihans = ref([])
const penyewas = ref([])
const kamars = ref([])
const loadingBtn = ref(false)
const showForm = ref(false)
const form = ref({
  id_penyewa: '',
  awal_periode: '',
  akhir_periode: '',
  jumlah_tagihan: 0,
  status: 'belum lunas',
  user_id: null
})
const editingId = ref(null)
const attrs = useAttrs()
const setNotif = attrs.setNotif || (() => {})
const errorMessage = ref('')

// Format currency
const formatRupiah = (value) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value)
}

// Get active penyewas (for dropdown)
const activePenyewas = computed(() => {
  return penyewas.value.filter(p => p.status === 'aktif')
})

// Display harga from kamar when penyewa is selected
const displayHarga = computed(() => {
  if (!form.value.id_penyewa) return formatRupiah(0)
  
  const penyewa = penyewas.value.find(p => p.id === form.value.id_penyewa)
  if (!penyewa) return formatRupiah(0)
  
  const kamar = kamars.value.find(k => k.id === penyewa.id_kamar)
  if (!kamar) return formatRupiah(0)
  
  // Set the form's jumlah_tagihan value
  form.value.jumlah_tagihan = kamar.harga
  
  return formatRupiah(kamar.harga)
})

// Create a combined data for display
const tagihanDisplay = computed(() => {
  return tagihans.value.map(tagihan => {
    const penyewa = penyewas.value.find(p => p.id === tagihan.id_penyewa) || {}
    const kamar = kamars.value.find(k => k.id === penyewa.id_kamar) || {}
    
    return {
      ...tagihan,
      penyewa_nama: penyewa.nama || `ID: ${tagihan.id_penyewa}`,
      kamar_info: kamar.nomer ? `Kamar ${kamar.nomer} - ${formatRupiah(kamar.harga)}` : '-',
      jumlah_tagihan: formatRupiah(tagihan.jumlah_tagihan)
    }
  })
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

// Get penyewa name
const getPenyewaName = (penyewaId) => {
  const penyewa = penyewas.value.find(p => p.id === penyewaId)
  return penyewa ? penyewa.nama : penyewaId
}

// Get kamar info 
const getKamarInfo = (kamarId) => {
  const kamar = kamars.value.find(k => k.id === kamarId)
  return kamar ? kamar.nomer : kamarId
}

// Handle penyewa selection change
function onPenyewaChange() {
  if (!form.value.id_penyewa) return
  
  const penyewa = penyewas.value.find(p => p.id === form.value.id_penyewa)
  if (!penyewa) return
  
  const kamar = kamars.value.find(k => k.id === penyewa.id_kamar)
  if (!kamar) return
  
  // Set the form's jumlah_tagihan value
  form.value.jumlah_tagihan = kamar.harga
}

async function fetchTagihans() {
  try {
    const res = await api.get('/tagihans')
    tagihans.value = res.data
  } catch (error) {
    errorMessage.value = 'Gagal memuat data tagihan'
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

function closeForm() {
  showForm.value = false
  resetForm()
}

function resetForm() {
  editingId.value = null
  form.value = { 
    id_penyewa: '', 
    awal_periode: '', 
    akhir_periode: '', 
    jumlah_tagihan: 0, 
    status: 'belum lunas',
    user_id: null
  }
}

async function submitTagihan() {
  loadingBtn.value = true
  
  // Set user_id dari local storage
  form.value.user_id = getUserId()
  
  try {
    if (editingId.value) {
      // When editing, only update the status
      const currentTagihan = tagihans.value.find(t => t.id === editingId.value)
      if (currentTagihan) {
        // Only update the status, preserve other fields
        await api.put(`/tagihans/${editingId.value}`, { 
          ...currentTagihan,
          status: form.value.status
        })
        setNotif('Status tagihan berhasil diupdate!')
      }
    } else {
      await api.post('/tagihans', form.value)
      setNotif('Tagihan berhasil ditambah!')
    }
    closeForm()
    await fetchTagihans()
  } catch (error) {
    console.error('Error saving tagihan:', error)
    if (error.response) {
      setNotif(`Gagal menyimpan tagihan: ${error.response.data.error || 'Terjadi kesalahan'}`)
    } else {
      setNotif('Gagal menyimpan tagihan: Tidak dapat terhubung ke server')
    }
  } finally {
    loadingBtn.value = false
  }
}

function editTagihan(tagihan) {
  // When editing, we only allow updating the status
  form.value = { 
    id_penyewa: tagihan.id_penyewa,
    status: tagihan.status,
    // Preserve other fields but they won't be editable
    awal_periode: tagihan.awal_periode,
    akhir_periode: tagihan.akhir_periode,
    jumlah_tagihan: tagihan.jumlah_tagihan
  }
  editingId.value = tagihan.id
  showForm.value = true
}

function batalEdit() {
  closeForm()
}

async function hapusTagihan(id) {
  if (confirm('Yakin ingin menghapus tagihan ini?')) {
    try {
      await api.delete(`/tagihans/${id}`)
      setNotif('Tagihan berhasil dihapus!')
      await fetchTagihans()
    } catch (error) {
      console.error('Error deleting tagihan:', error)
      if (error.response) {
        setNotif(`Gagal menghapus tagihan: ${error.response.data.error || 'Terjadi kesalahan'}`)
      } else {
        setNotif('Gagal menghapus tagihan: Tidak dapat terhubung ke server')
      }
    }
  }
}

onMounted(() => {
  fetchTagihans()
  fetchPenyewas()
  fetchKamars()
})
</script>

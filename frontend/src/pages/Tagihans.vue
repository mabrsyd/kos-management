<template>
  <div>
    <div class="flex flex-wrap justify-between items-center gap-4 mb-6">
      <div>
        <h2 class="text-xl sm:text-2xl font-bold text-gray-800 dark:text-gray-100">Manajemen Tagihan</h2>
        <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">Kelola tagihan bulanan penyewa kos</p>
      </div>
    </div>

    <!-- Statistik Tagihan -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center gap-3">
          <div class="bg-green-100 dark:bg-green-900/30 p-5 rounded-lg">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-10 text-green-600 dark:text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <div>
            <p class="text-sm cent text-gray-500 dark:text-gray-400">Lunas</p>
            <p class="text-xl font-bold text-gray-800 dark:text-gray-100">{{ stats.lunas }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center gap-3">
          <div class="bg-yellow-100 dark:bg-yellow-900/30 p-3 rounded-lg">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-yellow-600 dark:text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <p class="text-sm text-gray-500 dark:text-gray-400">Belum Lunas</p>
            <p class="text-xl font-bold text-gray-800 dark:text-gray-100">{{ stats.belumLunas }}</p>
          </div>
        </div>
      </div>
      
      <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700">
        <div class="flex items-center gap-3">
          <div class="bg-blue-100 dark:bg-blue-900/30 p-3 rounded-lg">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-600 dark:text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2z" />
            </svg>
          </div>
          <div>
            <p class="text-sm text-gray-500 dark:text-gray-400">Nyicil</p>
            <p class="text-xl font-bold text-gray-800 dark:text-gray-100">{{ stats.nyicil }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Statistik Pendapatan -->
    <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700 mb-6">
      <h3 class="font-semibold text-gray-800 dark:text-gray-200 mb-3">Pendapatan {{ filterMonth > 0 ? getMonthName(filterMonth) : 'Semua Bulan' }} {{ filterMonth > 0 ? form.tahun : '' }}</h3>
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <div class="bg-green-50 dark:bg-green-900/10 p-4 rounded-lg">
          <p class="text-sm text-gray-500 dark:text-gray-400">Pendapatan Lunas</p>
          <p class="text-lg font-bold text-green-600 dark:text-green-400">{{ formatRupiah(revenueStats.lunas) }}</p>
        </div>
        <div class="bg-blue-50 dark:bg-blue-900/10 p-4 rounded-lg">
          <p class="text-sm text-gray-500 dark:text-gray-400">Pendapatan Cicilan</p>
          <p class="text-lg font-bold text-blue-600 dark:text-blue-400">{{ formatRupiah(revenueStats.nyicil) }}</p>
        </div>
        <div class="bg-gray-50 dark:bg-gray-700/30 p-4 rounded-lg">
          <p class="text-sm text-gray-500 dark:text-gray-400">Total Pendapatan</p>
          <p class="text-lg font-bold text-gray-800 dark:text-gray-200">{{ formatRupiah(revenueStats.total) }}</p>
        </div>
      </div>
    </div>

    <!-- Filter dan Search Bar -->
    <div class="bg-white dark:bg-slate-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-slate-700 mb-6">
      <div class="flex flex-col md:flex-row gap-4">
        <!-- Search Box -->
        <div class="flex-1">
          <div class="relative">
            <input 
              v-model="searchQuery" 
              type="text" 
              placeholder="Cari nama penyewa atau kamar..."
              class="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 dark:border-slate-600 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            <div class="absolute left-3 top-2.5 text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Filter Controls -->
        <div class="flex flex-wrap gap-3">
          <select 
            v-model="filterMonth" 
            class="border border-gray-300 dark:border-slate-600 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 px-3 py-2 rounded-lg"
          >
            <option value="0">Semua Bulan</option>
            <option value="1">Januari</option>
            <option value="2">Februari</option>
            <option value="3">Maret</option>
            <option value="4">April</option>
            <option value="5">Mei</option>
            <option value="6">Juni</option>
            <option value="7">Juli</option>
            <option value="8">Agustus</option>
            <option value="9">September</option>
            <option value="10">Oktober</option>
            <option value="11">November</option>
            <option value="12">Desember</option>
          </select>
          
          <select 
            v-model="filterStatus" 
            class="border border-gray-300 dark:border-slate-600 bg-white dark:bg-slate-700 text-gray-700 dark:text-gray-200 px-3 py-2 rounded-lg"
          >
            <option value="">Semua Status</option>
            <option value="lunas">Lunas</option>
            <option value="belum_lunas">Belum Lunas</option>
            <option value="jatuh_tempo">Jatuh Tempo</option>
            <option value="nyicil">Nyicil</option>
          </select>
          
          <button 
            @click="showForm = true" 
            class="bg-blue-600 hover:bg-blue-700 text-white flex items-center gap-2 px-4 py-2 rounded-lg shadow-sm transition-colors duration-200"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            <span>Tambah Tagihan</span>
          </button>
          
          <button 
            @click="updateOverdueBills" 
            class="bg-yellow-600 hover:bg-yellow-700 text-white flex items-center gap-2 px-4 py-2 rounded-lg shadow-sm transition-colors duration-200"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Update Jatuh Tempo</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Table of Tagihans with responsive design -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-gray-100 dark:border-slate-700 overflow-hidden">
      <div class="overflow-x-auto">
        <AppTable :columns="columns" :data="filteredTagihans">
          <template #penyewa_nama="{ row }">
            <div class="font-medium text-blue-600 dark:text-blue-400">{{ row.penyewa_nama }}</div>
          </template>
          
          <template #action="{ row }">
            <div class="flex justify-end gap-2">
              <button @click="markAsPaid(row)" v-if="row.status !== 'lunas'" class="inline-flex items-center justify-center p-2 rounded-full bg-green-50 hover:bg-green-100 dark:bg-green-900/20 dark:hover:bg-green-800/30 text-green-600 dark:text-green-400 transition-colors" title="Tandai Lunas">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </button>
              <button @click="editTagihan(row)" class="inline-flex items-center justify-center p-2 rounded-full bg-blue-50 hover:bg-blue-100 dark:bg-blue-900/20 dark:hover:bg-blue-800/30 text-blue-600 dark:text-blue-400 transition-colors" title="Edit Status">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>
              <button @click="hapusTagihan(row.id)" class="inline-flex items-center justify-center p-2 rounded-full bg-red-50 hover:bg-red-100 dark:bg-red-900/20 dark:hover:bg-red-800/30 text-red-600 dark:text-red-400 transition-colors" title="Hapus">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </template>

          <!-- Custom template for status column -->
          <template #status="{ row }">
            <span 
              :class="{
                'px-2 py-1 rounded-full text-xs font-medium inline-flex items-center': true,
                'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400': row.status === 'lunas',
                'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400': row.status === 'belum_lunas',
                'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400': row.status === 'jatuh_tempo',
                'bg-blue-100 text-blue-800 dark:bg-blue-900/30 dark:text-blue-400': row.status === 'nyicil'
              }"
            >
              <span class="h-1.5 w-1.5 rounded-full mr-1" :class="{
                'bg-green-500 dark:bg-green-400': row.status === 'lunas',
                'bg-yellow-500 dark:bg-yellow-400': row.status === 'belum_lunas',
                'bg-red-500 dark:bg-red-400': row.status === 'jatuh_tempo',
                'bg-blue-500 dark:bg-blue-400': row.status === 'nyicil'
              }"></span>
              {{ formatStatus(row.status) }}
            </span>
          </template>

          <!-- Custom template for month/year column -->
          <template #period="{ row }">
            {{ getMonthName(row.bulan) }} {{ row.tahun }}
          </template>

          <!-- Custom template for due date -->
          <template #due_date="{ row }">
            {{ formatDate(row.tanggal_jatuh_tempo) }}
            <span v-if="isOverdue(row.tanggal_jatuh_tempo) && row.status !== 'lunas'" class="ml-1 text-xs text-red-500">(Lewat Jatuh Tempo)</span>
          </template>

          <!-- Tambahkan payment info column untuk menampilkan status pembayaran -->
          <template #payment_info="{ row }">
            <div>
              <div v-if="row.status === 'nyicil'" class="text-sm">
                <span class="text-gray-500 dark:text-gray-400">Dibayar: </span>
                <span class="font-medium">{{ formatRupiah(row.jumlah_dibayar || 0) }}</span>
                <span class="text-gray-400 dark:text-gray-500"> / {{ row.jumlah_tagihan_raw }}</span>
                <div class="text-sm text-blue-500 mt-1">
                  <span>Sisa: {{ formatRupiah(row.sisa_pembayaran) }}</span>
                </div>
                <!-- Debug info - hidden in production -->
                <div class="text-xs text-gray-400 mt-1" v-if="false">
                  Debug: jumlah_dibayar={{ row.jumlah_dibayar }}, 
                  jumlah_tagihan={{ row.jumlah_tagihan_raw }},
                  sisa={{ row.sisa_pembayaran }}
                </div>
              </div>
              <div v-else-if="row.status === 'lunas'" class="text-sm">
                <span class="text-green-500">Lunas</span>
              </div>
              <div v-else class="text-sm text-gray-400">-</div>
            </div>
          </template>
        </AppTable>
      </div>
    </div>

    <!-- Modal Form for Edit Status with improved design -->
    <div v-if="showForm" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50">
      <div class="bg-white dark:bg-slate-800 rounded-xl p-6 shadow-xl max-w-3xl w-full mx-4 max-h-[90vh] overflow-y-auto border border-gray-200 dark:border-slate-700">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-gray-800 dark:text-gray-100">{{ editingId ? 'Edit Status Tagihan' : 'Tambah Tagihan' }}</h3>
          <button @click="closeForm" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 p-1 rounded-lg hover:bg-gray-100 dark:hover:bg-slate-700 transition-colors duration-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <!-- Form Tambah/Edit Tagihan with improved design -->
        <form @submit.prevent="submitTagihan" class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Penyewa dropdown - show only when adding new tagihan -->
          <div v-if="!editingId" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <select 
                v-model="form.id_penyewa" 
                required 
                id="id_penyewa" 
                class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200"
                @change="onPenyewaChange"
              >
                <option value="" disabled>Pilih Penyewa</option>
                <option v-for="penyewa in activePenyewas" :key="penyewa.id" :value="penyewa.id">
                  {{ penyewa.nama }} - Kamar {{ getKamarInfo(penyewa.id_kamar) }}
                </option>
              </select>
              <label for="id_penyewa" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Penyewa</label>
            </div>
          </div>
          
          <!-- For editing, display penyewa name but don't allow changing -->
          <div v-if="editingId" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <p class="py-2 border-b-2 border-blue-200 dark:border-blue-500/30">{{ getPenyewaName(form.id_penyewa) }}</p>
              <label class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Penyewa</label>
            </div>
          </div>
          
          <div v-if="!editingId" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <select 
                v-model="form.bulan" 
                required 
                id="bulan" 
                class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200"
              >
                <option value="" disabled>Pilih Bulan</option>
                <option value="1">Januari</option>
                <option value="2">Februari</option>
                <option value="3">Maret</option>
                <option value="4">April</option>
                <option value="5">Mei</option>
                <option value="6">Juni</option>
                <option value="7">Juli</option>
                <option value="8">Agustus</option>
                <option value="9">September</option>
                <option value="10">Oktober</option>
                <option value="11">November</option>
                <option value="12">Desember</option>
              </select>
              <label for="bulan" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Bulan</label>
            </div>
          </div>
          
          <div v-if="!editingId" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input 
                v-model="form.tahun" 
                type="number" 
                required 
                min="2020" 
                max="2100" 
                id="tahun" 
                class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" 
              />
              <label for="tahun" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Tahun</label>
            </div>
          </div>
          
          <div v-if="!editingId" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input v-model="form.tanggal_jatuh_tempo" type="date" required id="tanggal_jatuh_tempo" class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" />
              <label for="tanggal_jatuh_tempo" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Tanggal Jatuh Tempo</label>
            </div>
          </div>
          
          <div v-if="!editingId" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input :value="displayHarga" disabled id="jumlah_tagihan" class="peer border-b-2 border-gray-200 dark:border-gray-600 outline-none w-full py-2 bg-transparent transition-colors duration-200 text-gray-500 dark:text-gray-400" />
              <label for="jumlah_tagihan" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Jumlah Tagihan</label>
            </div>
          </div>
          
          <div class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <select v-model="form.status" required id="status" class="border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200" @change="updateFormFields">
                <option value="lunas">Lunas</option>
                <option value="belum_lunas">Belum Lunas</option>
                <option value="jatuh_tempo">Jatuh Tempo</option>
                <option value="nyicil">Nyicil</option>
              </select>
              <label for="status" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Status</label>
            </div>
          </div>
          
          <!-- Input jumlah cicilan jika status = nyicil -->
          <div v-if="form.status === 'nyicil'" class="bg-gray-50 dark:bg-slate-700/50 p-4 rounded-lg">
            <div class="relative">
              <input 
                v-model="form.jumlah_dibayar" 
                type="number" 
                required 
                min="0" 
                id="jumlah_dibayar" 
                class="peer border-b-2 border-blue-200 focus:border-blue-500 outline-none w-full py-2 bg-transparent transition-colors duration-200"
                placeholder="Masukkan jumlah yang sudah dibayar" 
                @input="validatePaymentAmount"
              />
              <label for="jumlah_dibayar" class="absolute left-0 -top-5 text-xs text-blue-500 dark:text-blue-400">Jumlah Dibayar</label>
              <p class="text-xs text-gray-500 mt-1">
                <span class="font-medium">Total tagihan: {{ formatRupiah(form.jumlah_tagihan || 0) }}</span>
                <span v-if="form.jumlah_dibayar" class="block mt-1">
                  <span class="font-medium text-blue-500">Sisa pembayaran: {{ formatRupiah(getSisaPembayaran()) }}</span>
                </span>
              </p>
              <!-- Debugging info (hidden in production) -->
              <div v-if="false" class="mt-2 p-2 bg-gray-100 text-xs">
                <div>Debug info:</div>
                <div>jumlah_tagihan: {{ form.jumlah_tagihan }}</div>
                <div>jumlah_dibayar: {{ form.jumlah_dibayar }}</div>
                <div>sisa: {{ getSisaPembayaran() }}</div>
              </div>
            </div>
          </div>
          
          <div class="col-span-2 mt-4 flex gap-2 justify-end">
            <button type="button" @click="closeForm" class="bg-gray-200 dark:bg-slate-700 text-gray-700 dark:text-gray-300 px-4 py-2 rounded-lg font-medium hover:bg-gray-300 dark:hover:bg-slate-600 transition-colors duration-200">Batal</button>
            <button type="submit" :disabled="loadingBtn" class="bg-blue-600 text-white px-6 py-2 rounded-lg font-medium shadow hover:bg-blue-700 transition-colors duration-200 disabled:opacity-60 flex items-center gap-2">
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
import { ref, onMounted, computed, watch } from 'vue'
import api from '../utils/api'
import { useAttrs } from 'vue'
import AppTable from '../components/ui/AppTable.vue'

const columns = [
  { key: 'penyewa_nama', label: 'Penyewa' },
  { key: 'kamar_info', label: 'Kamar' }, 
  { key: 'period', label: 'Periode' },
  { key: 'due_date', label: 'Jatuh Tempo' },
  { key: 'jumlah_tagihan', label: 'Jumlah Tagihan' },
  { key: 'payment_info', label: 'Info Pembayaran' },
  { key: 'status', label: 'Status' }
]

const tagihans = ref([])
const penyewas = ref([])
const kamars = ref([])
const loadingBtn = ref(false)
const showForm = ref(false)
const filterMonth = ref(0)
const filterStatus = ref("")
const searchQuery = ref("")
const form = ref({
  id_penyewa: '',
  bulan: '',
  tahun: new Date().getFullYear(),
  tanggal_jatuh_tempo: new Date().toISOString().substr(0, 10),
  jumlah_tagihan: 0,
  jumlah_dibayar: 0,
  status: 'belum_lunas'
})
const editingId = ref(null)
const attrs = useAttrs()
const setNotif = attrs.setNotif || (() => {})
const errorMessage = ref('')

// Format currency
const formatRupiah = (value) => {
  // Handle invalid values
  if (value === undefined || value === null || isNaN(value)) {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(0)
  }
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value)
}

// Check if date is overdue
const isOverdue = (date) => {
  if (!date) return false
  const dueDate = new Date(date)
  const today = new Date()
  return dueDate < today
}

// Format date
const formatDate = (dateString) => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('id-ID', { 
    day: 'numeric', 
    month: 'long', 
    year: 'numeric' 
  }).format(date)
}

// Format status
const formatStatus = (status) => {
  switch(status) {
    case 'lunas': return 'Lunas'
    case 'belum_lunas': return 'Belum Lunas'
    case 'jatuh_tempo': return 'Jatuh Tempo'
    case 'nyicil': return 'Nyicil'
    default: return status
  }
}

// Get month name
const getMonthName = (month) => {
  const months = [
    'Januari', 'Februari', 'Maret', 'April', 'Mei', 'Juni',
    'Juli', 'Agustus', 'September', 'Oktober', 'November', 'Desember'
  ]
  return months[month - 1] || month
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

// Stats untuk menampilkan jumlah masing-masing status
const stats = computed(() => {
  return {
    lunas: tagihans.value.filter(t => t.status === 'lunas').length,
    belumLunas: tagihans.value.filter(t => t.status === 'belum_lunas').length,
    jatuhTempo: tagihans.value.filter(t => t.status === 'jatuh_tempo').length,
    nyicil: tagihans.value.filter(t => t.status === 'nyicil').length,
    total: tagihans.value.length
  }
})

// Revenue stats - pendapatan berdasarkan filter
const revenueStats = computed(() => {
  // Filter tagihan berdasarkan bulan jika filter bulan aktif
  let filteredData = tagihans.value
  if (filterMonth.value > 0) {
    filteredData = filteredData.filter(t => t.bulan == filterMonth.value)
  }
  
  // Hitung pendapatan dari yang lunas dan nyicil
  const lunasAmount = filteredData
    .filter(t => t.status === 'lunas')
    .reduce((sum, t) => sum + Number(t.jumlah_tagihan), 0)
    
  const nyicilAmount = filteredData
    .filter(t => t.status === 'nyicil')
    .reduce((sum, t) => {
      // Make sure jumlah_dibayar has a valid value
      const dibayar = Number(t.jumlah_dibayar) || 0
      return sum + dibayar
    }, 0)
    
  // Log for debugging
  console.log('Revenue stats calculation:')
  console.log('- Lunas amount:', lunasAmount)
  console.log('- Nyicil amount:', nyicilAmount)
  console.log('- Total:', lunasAmount + nyicilAmount)
    
  return {
    lunas: lunasAmount,
    nyicil: nyicilAmount,
    total: lunasAmount + nyicilAmount
  }
})

// Filtered tagihans based on search, month and status
const filteredTagihans = computed(() => {
  let filtered = tagihans.value
  
  // Apply search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(tagihan => {
      const penyewa = penyewas.value.find(p => p.id === tagihan.id_penyewa) || {}
      const kamar = kamars.value.find(k => k.id === penyewa.id_kamar) || {}
      
      // Search by penyewa name or kamar info
      return (penyewa.nama && penyewa.nama.toLowerCase().includes(query)) ||
             (kamar.nomer && kamar.nomer.toLowerCase().includes(query))
    })
  }
  
  // Apply month filter if selected
  if (filterMonth.value > 0) {
    filtered = filtered.filter(t => t.bulan == filterMonth.value)
  }
  
  // Apply status filter if selected
  if (filterStatus.value) {
    filtered = filtered.filter(t => t.status === filterStatus.value)
  }
  
  // Map to display format
  return filtered.map(tagihan => {
    const penyewa = penyewas.value.find(p => p.id === tagihan.id_penyewa) || {}
    const kamar = kamars.value.find(k => k.id === penyewa.id_kamar) || {}
    
    // Make sure jumlah_dibayar is handled properly, default to 0 if undefined or null
    const jumlahDibayar = Number(tagihan.jumlah_dibayar) || 0
    
    // Calculate remaining payment
    const jumlahTagihan = Number(tagihan.jumlah_tagihan) || 0
    const sisaPembayaran = Math.max(0, jumlahTagihan - jumlahDibayar)
    
    // Debug log individual transformation
    if (tagihan.status === 'nyicil') {
      console.log(`Processing nyicil: ID=${tagihan.id}, jumlahTagihan=${jumlahTagihan}, jumlahDibayar=${jumlahDibayar}, sisa=${sisaPembayaran}`)
    }
    
    return {
      ...tagihan,
      penyewa_nama: penyewa.nama || `ID: ${tagihan.id_penyewa}`,
      kamar_info: kamar.nomer ? `Kamar ${kamar.nomer} - ${formatRupiah(kamar.harga)}` : '-',
      jumlah_tagihan: formatRupiah(jumlahTagihan),
      jumlah_tagihan_raw: formatRupiah(jumlahTagihan), // Format rupiah untuk display
      jumlah_dibayar: jumlahDibayar,
      sisa_pembayaran: sisaPembayaran
    }
  })
})

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

// Validate and handle payment amount
function validatePaymentAmount() {
  console.log('Validating payment amount:', form.value.jumlah_dibayar)
  
  // If input is empty, don't change it (allow empty string)
  if (form.value.jumlah_dibayar === '' || form.value.jumlah_dibayar === null) {
    console.log('Empty input, skipping validation')
    return;
  }
  
  // Ensure jumlah_dibayar is a number
  form.value.jumlah_dibayar = Number(form.value.jumlah_dibayar)
  
  // Validate the amount
  if (isNaN(form.value.jumlah_dibayar)) {
    console.log('NaN detected, resetting to 0')
    form.value.jumlah_dibayar = 0
  }

  // Ensure the amount is positive
  if (form.value.jumlah_dibayar < 0) {
    console.log('Negative value detected, resetting to 0')
    form.value.jumlah_dibayar = 0
  }
  
  // Log final validated value
  console.log('Payment amount after validation:', form.value.jumlah_dibayar)
  
  // Calculate and log the remaining amount
  const sisa = getSisaPembayaran()
  console.log('Remaining payment after validation:', sisa)
}

// Calculate remaining payment safely
function getSisaPembayaran() {
  console.log('Calculating sisa pembayaran with values:', {
    jumlah_tagihan: form.value.jumlah_tagihan,
    jumlah_dibayar: form.value.jumlah_dibayar
  })
  
  const tagihan = Number(form.value.jumlah_tagihan) || 0
  
  // Handle empty input case
  if (form.value.jumlah_dibayar === '' || form.value.jumlah_dibayar === null) {
    console.log('Empty jumlah_dibayar, returning full tagihan:', tagihan)
    return tagihan
  }
  
  const dibayar = Number(form.value.jumlah_dibayar) || 0
  const sisa = Math.max(0, tagihan - dibayar)
  
  console.log(`Calculated: ${tagihan} - ${dibayar} = ${sisa}`)
  return sisa
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
    console.log('Fetched tagihans (raw):', res.data) // Debug log
    
    if (!Array.isArray(res.data)) {
      console.error('Expected array of tagihans but got:', typeof res.data)
      return
    }
    
    // Ensure jumlah_dibayar and jumlah_tagihan are properly parsed for each tagihan
    tagihans.value = res.data.map(tagihan => {
      // Make sure we have valid numbers for monetary values
      const processedTagihan = {
        ...tagihan,
        jumlah_tagihan: Number(tagihan.jumlah_tagihan) || 0,
        jumlah_dibayar: Number(tagihan.jumlah_dibayar) || 0
      }
      
      // Log nyicil tagihans for debugging
      if (tagihan.status === 'nyicil') {
        console.log('Nyicil tagihan processed:', processedTagihan)
      }
      
      return processedTagihan
    })
    
    console.log('Processed all tagihans') // Debug log
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

async function updateOverdueBills() {
  try {
    await api.put('/tagihans/update-overdue')
    setNotif('Tagihan jatuh tempo berhasil diperbarui!')
    await fetchTagihans()
  } catch (error) {
    console.error('Error updating overdue bills:', error)
    setNotif('Gagal memperbarui tagihan jatuh tempo')
  }
}

async function markAsPaid(tagihan) {
  if (confirm(`Tandai tagihan ${getMonthName(tagihan.bulan)} ${tagihan.tahun} untuk ${tagihan.penyewa_nama} sebagai lunas?`)) {
    try {
      await api.put(`/tagihans/${tagihan.id}/mark-paid`)
      setNotif('Tagihan berhasil ditandai lunas!')
      await fetchTagihans()
    } catch (error) {
      console.error('Error marking bill as paid:', error)
      setNotif('Gagal menandai tagihan sebagai lunas')
    }
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
    bulan: '',
    tahun: new Date().getFullYear(),
    tanggal_jatuh_tempo: new Date().toISOString().substr(0, 10),
    jumlah_tagihan: 0, 
    jumlah_dibayar: 0,
    status: 'belum_lunas'
  }
}

async function submitTagihan() {
  loadingBtn.value = true
  
  try {
    // Create a copy of the form data to modify
    const tagihanData = { ...form.value }
    
    // Ensure numeric fields are proper numbers, not strings or NaN
    tagihanData.id_penyewa = Number(tagihanData.id_penyewa) || 0
    tagihanData.bulan = Number(tagihanData.bulan) || 0
    tagihanData.tahun = Number(tagihanData.tahun) || 0
    tagihanData.jumlah_tagihan = Number(tagihanData.jumlah_tagihan) || 0
    
    // Validate important fields
    if (!tagihanData.id_penyewa || !tagihanData.bulan || !tagihanData.tahun || !tagihanData.jumlah_tagihan) {
      setNotif('Data tidak valid. Pastikan semua field telah diisi dengan benar')
      loadingBtn.value = false
      return
    }
    
    // Handle jumlah_dibayar based on status
    if (tagihanData.status === 'nyicil') {
      // Convert to number - make sure to handle empty string
      const dibayar = tagihanData.jumlah_dibayar === '' || tagihanData.jumlah_dibayar === null 
        ? 0 
        : Number(tagihanData.jumlah_dibayar)
      
      // Set the converted value
      tagihanData.jumlah_dibayar = dibayar
      
      // Validation - make sure jumlah_dibayar is valid
      if (dibayar <= 0) {
        setNotif('Jumlah cicilan harus lebih dari 0')
        loadingBtn.value = false
        return
      }
      
      if (dibayar >= tagihanData.jumlah_tagihan) {
        setNotif('Jumlah cicilan harus kurang dari total tagihan. Untuk pembayaran penuh, gunakan status Lunas.')
        loadingBtn.value = false
        return
      }
    } else if (tagihanData.status === 'lunas') {
      tagihanData.jumlah_dibayar = tagihanData.jumlah_tagihan
    } else {
      tagihanData.jumlah_dibayar = 0
    }

    // Convert date string to proper format if it exists
    if (tagihanData.tanggal_jatuh_tempo) {
      // Add time component to ensure proper datetime format
      const dateObj = new Date(tagihanData.tanggal_jatuh_tempo)
      // Make sure it's a valid date
      if (!isNaN(dateObj.getTime())) {
        // Keep the date in ISO format
        tagihanData.tanggal_jatuh_tempo = dateObj.toISOString()
      }
    }
    
    // Debug logging
    console.log('Sending data to API:', JSON.stringify(tagihanData))
    
    if (editingId.value) {
      // When editing, update the status and cicilan info
      await api.put(`/tagihans/${editingId.value}`, tagihanData)
      setNotif('Status tagihan berhasil diupdate!')
    } else {
      // For new tagihan
      await api.post('/tagihans', tagihanData)
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

// Get a single tagihan by ID for editing
async function fetchTagihanById(id) {
  try {
    const res = await api.get(`/tagihans/${id}`)
    console.log('Fetched single tagihan:', res.data)
    return res.data
  } catch (error) {
    console.error(`Error fetching tagihan with ID ${id}:`, error)
    return null
  }
}

async function editTagihan(tagihan) {
  // Log the raw data we're editing for debugging
  console.log('Editing tagihan (raw):', tagihan)
  
  try {
    // Fetch the most up-to-date tagihan data from the server
    const freshTagihan = await fetchTagihanById(tagihan.id)
    
    if (freshTagihan) {
      // Use the fresh data
      tagihan = freshTagihan
      console.log('Using fresh tagihan data:', tagihan)
    } else {
      console.log('Using existing tagihan data (fetch failed)')
    }
    
    // When editing, ensure we have clean numeric values
    const jumlahTagihan = Number(tagihan.jumlah_tagihan) || 0
    const jumlahDibayar = Number(tagihan.jumlah_dibayar) || 0
    
    console.log(`Processing values - jumlahTagihan: ${jumlahTagihan}, jumlahDibayar: ${jumlahDibayar}`)
    
    // When editing, set all the fields
    form.value = { 
      id_penyewa: tagihan.id_penyewa,
      bulan: tagihan.bulan,
      tahun: tagihan.tahun,
      status: tagihan.status,
      tanggal_jatuh_tempo: tagihan.tanggal_jatuh_tempo ? new Date(tagihan.tanggal_jatuh_tempo).toISOString().substr(0, 10) : '',
      jumlah_tagihan: jumlahTagihan,
      jumlah_dibayar: jumlahDibayar
    }
    
    // Log the form data for debugging
    console.log('Form data after setting:', form.value)
    
    // Force update of the sisa pembayaran calculation
    if (tagihan.status === 'nyicil') {
      console.log('Sisa pembayaran:', getSisaPembayaran())
    }
    
    editingId.value = tagihan.id
    showForm.value = true
  } catch (error) {
    console.error('Error in editTagihan:', error)
    setNotif('Terjadi kesalahan saat mengedit tagihan')
  }
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

// Update form fields based on status changes
function updateFormFields() {
  if (form.value.status === 'lunas') {
    form.value.jumlah_dibayar = form.value.jumlah_tagihan
  } else if (form.value.status === 'belum_lunas' || form.value.status === 'jatuh_tempo') {
    form.value.jumlah_dibayar = 0
  }
}

// Watch for changes to filterMonth to refresh data
watch(filterMonth, () => {
  // Refresh data if needed
})

onMounted(() => {
  fetchTagihans()
  fetchPenyewas()
  fetchKamars()
})
</script>

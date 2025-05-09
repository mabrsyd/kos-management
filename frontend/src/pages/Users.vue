<script setup>
import { ref, onMounted } from 'vue'
import api from '../utils/api'

const users = ref([])
const loading = ref(true)
const isModalOpen = ref(false)
const isEditMode = ref(false)
const currentUser = ref({
  username: '',
  name: '',
  password: '',
  role: 'user'
})
const errorMessage = ref('')

// Methods
const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await api.get('/users')
    users.value = res.data
  } catch (error) {
    console.error('Error fetching users:', error)
  } finally {
    loading.value = false
  }
}

const openAddModal = () => {
  isEditMode.value = false
  currentUser.value = {
    username: '',
    name: '',
    password: '',
    role: 'user'
  }
  isModalOpen.value = true
  errorMessage.value = ''
}

const openEditModal = (user) => {
  isEditMode.value = true
  currentUser.value = {
    id: user.id,
    username: user.username,
    name: user.name,
    role: user.role,
    password: ''
  }
  isModalOpen.value = true
  errorMessage.value = ''
}

const saveUser = async () => {
  try {
    if (isEditMode.value) {
      await api.put(`/users/${currentUser.value.id}`, currentUser.value)
    } else {
      await api.post('/users', currentUser.value)
    }
    isModalOpen.value = false
    fetchUsers()
  } catch (error) {
    errorMessage.value = error.response?.data?.error || 'Terjadi kesalahan saat menyimpan'
  }
}

const deleteUser = async (id) => {
  if (confirm('Apakah Anda yakin ingin menghapus user ini?')) {
    try {
      await api.delete(`/users/${id}`)
      fetchUsers()
    } catch (error) {
      alert('Gagal menghapus user')
    }
  }
}

onMounted(fetchUsers)
</script>

<template>
  <div class="p-4">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-semibold text-gray-800 dark:text-white">Manajemen Pengguna</h2>
      <button 
        @click="openAddModal" 
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded flex items-center"
      >
        <span class="mr-1">+</span> Tambah User
      </button>
    </div>
    
    <div v-if="loading" class="flex justify-center p-8">
      <div class="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-500"></div>
    </div>
    
    <div v-else class="bg-white dark:bg-gray-800 rounded-lg shadow overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
        <thead class="bg-gray-50 dark:bg-gray-700">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Username</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Nama</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Role</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
          <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50 dark:hover:bg-gray-700">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">{{ user.username }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">{{ user.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">{{ user.role }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button @click="openEditModal(user)" class="text-indigo-600 hover:text-indigo-900 dark:text-indigo-400 dark:hover:text-indigo-300 mr-3">Edit</button>
              <button @click="deleteUser(user.id)" class="text-red-600 hover:text-red-900 dark:text-red-400 dark:hover:text-red-300">Hapus</button>
            </td>
          </tr>
          <tr v-if="users.length === 0">
            <td colspan="4" class="px-6 py-4 text-center text-sm text-gray-500 dark:text-gray-400">Tidak ada data pengguna</td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- Modal Form -->
    <div v-if="isModalOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6 w-full max-w-md">
        <h3 class="text-lg font-semibold mb-4 text-gray-800 dark:text-white">{{ isEditMode ? 'Edit User' : 'Tambah User' }}</h3>
        
        <div v-if="errorMessage" class="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
          {{ errorMessage }}
        </div>
        
        <form @submit.prevent="saveUser" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Username</label>
            <input 
              v-model="currentUser.username" 
              type="text" 
              :disabled="isEditMode"
              required 
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Nama</label>
            <input 
              v-model="currentUser.name" 
              type="text" 
              required 
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Password {{ isEditMode ? '(Kosongkan jika tidak ingin mengubah)' : '' }}</label>
            <input 
              v-model="currentUser.password" 
              type="password" 
              :required="!isEditMode"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Role</label>
            <select 
              v-model="currentUser.role" 
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            >
              <option value="user">User</option>
              <option value="admin">Admin</option>
            </select>
          </div>
          
          <div class="flex justify-end space-x-3 mt-6">
            <button 
              type="button" 
              @click="isModalOpen = false" 
              class="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 dark:bg-gray-700 dark:text-white dark:border-gray-600 dark:hover:bg-gray-600"
            >
              Batal
            </button>
            <button 
              type="submit" 
              class="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Simpan
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
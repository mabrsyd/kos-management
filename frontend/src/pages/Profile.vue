<template>
  <div class="space-y-8">
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold text-gray-800 dark:text-white">Profil</h1>
    </div>
    
    <!-- Profile Info -->
    <div class="bg-white dark:bg-slate-800 rounded-xl shadow-md p-6">
      <div v-if="loading" class="flex justify-center py-10">
        <div class="animate-spin h-8 w-8 border-3 border-blue-500 rounded-full border-t-transparent"></div>
      </div>
      
      <div v-else class="space-y-6">
        <!-- Success message -->
        <div v-if="successMessage" class="mb-4 bg-green-100 border-l-4 border-green-500 text-green-700 p-4 rounded-md flex items-start">
          <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          <span>{{ successMessage }}</span>
        </div>
        
        <!-- Error message -->
        <div v-if="errorMessage" class="mb-4 bg-red-100 border-l-4 border-red-500 text-red-700 p-4 rounded-md flex items-start">
          <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
          </svg>
          <span>{{ errorMessage }}</span>
        </div>
        
        <!-- Profile header -->
        <div class="flex items-center space-x-4">
          <div class="h-24 w-24 bg-blue-600 rounded-full flex items-center justify-center text-white text-3xl font-bold shadow-lg">
            {{ userInitials }}
          </div>
          <div>
            <h2 class="text-xl font-bold text-gray-800 dark:text-white">{{ userData.name }}</h2>
            <p class="text-sm text-gray-500 dark:text-gray-400">{{ userData.username }}</p>
            <span class="inline-block mt-1 px-3 py-1 text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200 rounded-full">
              {{ userData.role === 'admin' ? 'Administrator' : 'Pengguna' }}
            </span>
          </div>
        </div>
        
        <hr class="my-6 border-gray-200 dark:border-gray-700">
        
        <!-- Profile Form -->
        <form @submit.prevent="updateProfile" class="space-y-6 max-w-lg">
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Nama</label>
            <input
              id="name"
              v-model="userData.name"
              type="text"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-700 dark:border-slate-600 dark:text-white"
              required
            />
          </div>
          
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Username</label>
            <input
              id="username"
              v-model="userData.username"
              type="text"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-700 dark:border-slate-600 dark:text-white dark:disabled:bg-slate-800 dark:disabled:border-slate-700"
              disabled
            />
            <p class="mt-1 text-xs text-gray-500 dark:text-gray-400">Username tidak dapat diubah</p>
          </div>
          
          <div>
            <label for="new-password" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Password Baru (opsional)</label>
            <input
              id="new-password"
              v-model="userData.newPassword"
              type="password"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-700 dark:border-slate-600 dark:text-white"
              placeholder="Biarkan kosong jika tidak ingin mengubah password"
            />
          </div>
          
          <div>
            <label for="confirm-password" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Konfirmasi Password Baru</label>
            <input
              id="confirm-password"
              v-model="userData.confirmPassword"
              type="password"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-700 dark:border-slate-600 dark:text-white"
              :disabled="!userData.newPassword"
            />
          </div>
          
          <div>
            <label for="current-password" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Password Saat Ini</label>
            <input
              id="current-password"
              v-model="userData.currentPassword"
              type="password"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-700 dark:border-slate-600 dark:text-white"
              required
              placeholder="Masukkan password saat ini untuk konfirmasi perubahan"
            />
          </div>
          
          <div class="flex justify-end">
            <button
              type="submit"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors duration-300 flex items-center"
              :disabled="isLoading"
            >
              <svg v-if="isLoading" class="animate-spin -ml-1 mr-2 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>{{ isLoading ? 'Menyimpan...' : 'Simpan Perubahan' }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import api from '../utils/api'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(true)
const isLoading = ref(false)
const successMessage = ref('')
const errorMessage = ref('')

const userData = ref({
  id: null,
  name: '',
  username: '',
  role: '',
  newPassword: '',
  confirmPassword: '',
  currentPassword: ''
})

// Get user initials for avatar
const userInitials = computed(() => {
  if (!userData.value.name) return 'U'
  return userData.value.name
    .split(' ')
    .map(name => name.charAt(0))
    .join('')
    .substring(0, 2)
    .toUpperCase()
})

// Load user data
onMounted(async () => {
  loading.value = true
  
  try {
    // Get user data from localStorage
    const userJson = localStorage.getItem('user')
    if (!userJson) {
      router.push('/login')
      return
    }
    
    const user = JSON.parse(userJson)
    userData.value.id = user.id
    userData.value.name = user.name
    userData.value.username = user.username
    userData.value.role = user.role
    
    // Fetch latest user data from server for verification
    const response = await api.get(`/users/${user.id}`)
    const fetchedUser = response.data
    
    // Update with fresh data from server
    userData.value.name = fetchedUser.name
    userData.value.username = fetchedUser.username
    userData.value.role = fetchedUser.role
    
  } catch (error) {
    console.error('Error loading user data:', error)
    errorMessage.value = 'Gagal memuat data profil'
  } finally {
    loading.value = false
  }
})

// Update profile
const updateProfile = async () => {
  // Reset messages
  successMessage.value = ''
  errorMessage.value = ''
  
  // Validate password if user is trying to change it
  if (userData.value.newPassword) {
    if (userData.value.newPassword.length < 6) {
      errorMessage.value = 'Password harus minimal 6 karakter'
      return
    }
    
    if (userData.value.newPassword !== userData.value.confirmPassword) {
      errorMessage.value = 'Password baru dan konfirmasi password tidak sama'
      return
    }
  }
  
  // Prepare update payload
  const updateData = {
    name: userData.value.name,
    password: userData.value.currentPassword
  }
  
  // Add new password if provided
  if (userData.value.newPassword) {
    updateData.newPassword = userData.value.newPassword
  }
  
  isLoading.value = true
  
  try {
    // Send update request to the profile endpoint
    const response = await api.post('/profile', updateData)
    
    // Update local storage with new data
    const user = {
      id: userData.value.id,
      username: userData.value.username,
      name: userData.value.name,
      role: userData.value.role
    }
    
    localStorage.setItem('user', JSON.stringify(user))
    
    // Clear password fields
    userData.value.newPassword = ''
    userData.value.confirmPassword = ''
    userData.value.currentPassword = ''
    
    // Show success message
    successMessage.value = 'Profil berhasil diperbarui'
  } catch (error) {
    console.error('Error updating profile:', error)
    
    if (error.response?.status === 401) {
      errorMessage.value = 'Password saat ini tidak valid'
    } else if (error.response?.data?.error) {
      errorMessage.value = error.response.data.error
    } else {
      errorMessage.value = 'Gagal memperbarui profil'
    }
  } finally {
    isLoading.value = false
  }
}
</script> 
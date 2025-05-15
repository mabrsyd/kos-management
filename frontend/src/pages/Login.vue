<template>
  <div class="flex min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
    <div class="m-auto w-full max-w-md p-8 bg-white rounded-xl shadow-xl transition-all duration-300">
      <!-- Logo & Header -->
      <div class="text-center mb-8">
        <div class="flex justify-center mb-3">
          <div class="h-20 w-20 bg-blue-600 rounded-full flex items-center justify-center text-white text-3xl font-bold shadow-lg">
            KM
          </div>
        </div>
        <h1 class="text-3xl font-bold bg-gradient-to-r from-blue-600 to-indigo-600 text-transparent bg-clip-text mb-2">KOS Management</h1>
        <p class="text-gray-600">{{ isRegisterMode ? 'Daftar akun baru' : 'Masuk ke akun Anda' }}</p>
      </div>
      
      <!-- Alert Messages -->
      <div v-if="errorMessage" class="mb-4 bg-red-100 border-l-4 border-red-500 text-red-700 p-4 rounded-md flex items-start">
        <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
        </svg>
        <span>{{ errorMessage }}</span>
      </div>
      
      <div v-if="successMessage" class="mb-4 bg-green-100 border-l-4 border-green-500 text-green-700 p-4 rounded-md flex items-start">
        <svg class="w-5 h-5 mr-2 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
        </svg>
        <span>{{ successMessage }}</span>
      </div>
      
      <!-- Form Login -->
      <form v-if="!isRegisterMode" @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="username" 
              id="username" 
              name="username"
              type="text" 
              autocomplete="username"
              required 
              class="pl-10 pr-4 py-2 block w-full rounded-md border-gray-300 shadow-sm focus:ring-blue-500 focus:border-blue-500"
              placeholder="Masukkan username"
            />
          </div>
        </div>
        
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="password" 
              id="password" 
              name="password"
              type="password" 
              autocomplete="current-password"
              required 
              class="pl-10 pr-4 py-2 block w-full rounded-md border-gray-300 shadow-sm focus:ring-blue-500 focus:border-blue-500"
              placeholder="Masukkan password"
            />
          </div>
        </div>
        
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <input id="remember_me" name="remember_me" type="checkbox" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
            <label for="remember_me" class="ml-2 block text-sm text-gray-700">
              Ingat saya
            </label>
          </div>
          
          <div class="text-sm">
            <a href="#" class="font-medium text-blue-600 hover:text-blue-500">
              Lupa password?
            </a>
          </div>
        </div>
        
        <div>
          <button 
            type="submit" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-300"
            :disabled="isLoading"
          >
            <span v-if="isLoading" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Memproses...
            </span>
            <span v-else>Masuk</span>
          </button>
        </div>
        
        <div class="text-center text-sm mt-4">
          <p class="text-gray-600">
            Belum memiliki akun? 
            <button 
              type="button" 
              @click="toggleMode" 
              class="text-blue-600 hover:text-blue-800 font-medium transition-colors duration-300">
              Daftar sekarang
            </button>
          </p>
        </div>
      </form>
      
      <!-- Form Register -->
      <form v-else @submit.prevent="handleRegister" class="space-y-6">
        <div>
          <label for="reg-username" class="block text-sm font-medium text-gray-700">Username</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="registerData.username" 
              id="reg-username" 
              name="reg-username"
              type="text" 
              autocomplete="username"
              required 
              class="pl-10 pr-4 py-2 block w-full rounded-md border-gray-300 shadow-sm focus:ring-blue-500 focus:border-blue-500"
              placeholder="Pilih username"
            />
          </div>
        </div>
        
        <div>
          <label for="reg-name" class="block text-sm font-medium text-gray-700">Nama Lengkap</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 005 10a6 6 0 0012 0c0-.35-.039-.69-.092-1.021A5.001 5.001 0 0010 11z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="registerData.name" 
              id="reg-name" 
              name="reg-name"
              type="text" 
              required 
              class="pl-10 pr-4 py-2 block w-full rounded-md border-gray-300 shadow-sm focus:ring-blue-500 focus:border-blue-500"
              placeholder="Masukkan nama lengkap"
            />
          </div>
        </div>
        
        <div>
          <label for="reg-password" class="block text-sm font-medium text-gray-700">Password</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="registerData.password" 
              id="reg-password" 
              name="reg-password"
              type="password" 
              autocomplete="new-password"
              required 
              class="pl-10 pr-4 py-2 block w-full rounded-md border-gray-300 shadow-sm focus:ring-blue-500 focus:border-blue-500"
              placeholder="Minimal 6 karakter"
            />
          </div>
        </div>
        
        <div>
          <label for="reg-confirm-password" class="block text-sm font-medium text-gray-700">Konfirmasi Password</label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="registerData.confirmPassword" 
              id="reg-confirm-password" 
              name="reg-confirm-password"
              type="password" 
              autocomplete="new-password"
              required 
              class="pl-10 pr-4 py-2 block w-full rounded-md border-gray-300 shadow-sm focus:ring-blue-500 focus:border-blue-500"
              placeholder="Ulangi password"
            />
          </div>
        </div>
        
        <div>
          <button 
            type="submit" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-300"
            :disabled="isLoading"
          >
            <span v-if="isLoading" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Memproses...
            </span>
            <span v-else>Daftar</span>
          </button>
        </div>
        
        <div class="text-center text-sm mt-4">
          <p class="text-gray-600">
            Sudah memiliki akun? 
            <button 
              type="button" 
              @click="toggleMode" 
              class="text-blue-600 hover:text-blue-800 font-medium transition-colors duration-300">
              Login
            </button>
          </p>
        </div>
      </form>

      <div v-if="checkingServer" class="mt-6 text-center text-gray-600">
        <p>Memeriksa koneksi server...</p>
        <div class="flex justify-center mt-2">
          <div class="animate-spin h-6 w-6 border-3 border-blue-500 rounded-full border-t-transparent"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import api from '../utils/api'

export default {
  name: 'LoginPage',
  data() {
    return {
      username: '',
      password: '',
      errorMessage: '',
      successMessage: '',
      isLoading: false,
      checkingServer: true,
      serverPorts: [8080, 8081, 3000],
      isRegisterMode: false,
      registerData: {
        username: '',
        name: '',
        password: '',
        confirmPassword: '',
        role: 'user'
      },
      activeServerPort: 8080,
      rememberMe: false
    }
  },
  async mounted() {
    // Check if returning user with saved token
    const token = localStorage.getItem('token')
    const user = localStorage.getItem('user')
    
    if (token && user) {
      try {
        const userData = JSON.parse(user)
        console.log('User found in localStorage:', userData.username)
      } catch (error) {
        console.error('Error parsing user data', error)
        localStorage.removeItem('token')
        localStorage.removeItem('user')
      }
    }
    
    // Check server connection on mount
    this.checkingServer = true
    await this.checkServerConnection()
    this.checkingServer = false
  },
  methods: {
    toggleMode() {
      this.isRegisterMode = !this.isRegisterMode
      this.errorMessage = ''
      this.successMessage = ''
    },
    async checkServerConnection() {
      for (const port of this.serverPorts) {
        try {
          await axios.get(`http://localhost:${port}/api/ping`, { 
            timeout: 2000,
            headers: { 'Content-Type': 'application/json' } 
          })
          console.log(`Server ditemukan pada port ${port}`)
          this.activeServerPort = port
          return true
        } catch (error) {
          console.log(`Port ${port} tidak merespon`)
        }
      }
      this.errorMessage = 'Tidak dapat terhubung dengan server. Pastikan server backend berjalan.'
      return false
    },
    async handleLogin() {
      this.isLoading = true
      this.errorMessage = ''
      this.successMessage = ''
      
      // Check server connection first
      const isServerConnected = await this.checkServerConnection()
      if (!isServerConnected) {
        this.isLoading = false
        return
      }
      
      try {
        const response = await axios.post(`http://localhost:${this.activeServerPort}/api/login`, {
          username: this.username,
          password: this.password
        }, {
          headers: {
            'Content-Type': 'application/json'
          },
          withCredentials: false,
          timeout: 5000
        })
        
        // Store token and user data
        localStorage.setItem('token', response.data.token)
        localStorage.setItem('user', JSON.stringify(response.data.user))
        
        // Redirect to dashboard
        this.$router.push('/')
      } catch (error) {
        if (error.code === 'ECONNABORTED') {
          this.errorMessage = 'Permintaan timeout. Server mungkin lambat merespon.'
        } else if (!error.response) {
          this.errorMessage = 'Tidak dapat terhubung dengan server. Pastikan server berjalan.'
        } else if (error.response) {
          this.errorMessage = error.response.data.error || 'Login gagal'
        } else {
          this.errorMessage = 'Terjadi kesalahan saat login'
        }
      } finally {
        this.isLoading = false
      }
    },
    async handleRegister() {
      // Validasi input
      if (this.registerData.password !== this.registerData.confirmPassword) {
        this.errorMessage = 'Password dan konfirmasi password tidak sama'
        return
      }
      
      if (this.registerData.password.length < 6) {
        this.errorMessage = 'Password harus minimal 6 karakter'
        return
      }
      
      this.isLoading = true
      this.errorMessage = ''
      this.successMessage = ''
      
      // Check server connection first
      const isServerConnected = await this.checkServerConnection()
      if (!isServerConnected) {
        this.isLoading = false
        return
      }
      
      try {
        await axios.post(`http://localhost:${this.activeServerPort}/api/register`, {
          username: this.registerData.username,
          name: this.registerData.name,
          password: this.registerData.password,
          role: this.registerData.role
        }, {
          headers: {
            'Content-Type': 'application/json'
          },
          withCredentials: false,
          timeout: 5000
        })
        
        // Reset form and show success message
        this.successMessage = 'Registrasi berhasil! Silahkan login.'
        this.isRegisterMode = false
        
        // Pre-fill login form with registered username
        this.username = this.registerData.username
        this.password = ''
        
        // Reset register form
        this.registerData = {
          username: '',
          name: '',
          password: '',
          confirmPassword: '',
          role: 'user'
        }
      } catch (error) {
        if (error.code === 'ECONNABORTED') {
          this.errorMessage = 'Permintaan timeout. Server mungkin lambat merespon.'
        } else if (!error.response) {
          this.errorMessage = 'Tidak dapat terhubung dengan server. Pastikan server berjalan.'
        } else if (error.response) {
          this.errorMessage = error.response.data.error || 'Registrasi gagal'
        } else {
          this.errorMessage = 'Terjadi kesalahan saat registrasi'
        }
      } finally {
        this.isLoading = false
      }
    }
  }
}
</script>

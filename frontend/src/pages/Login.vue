<template>
  <div class="flex min-h-screen bg-gray-100">
    <div class="m-auto w-full max-w-md p-8 bg-white rounded-lg shadow-md">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-800">KOS Management</h1>
        <p class="text-gray-600">{{ isRegisterMode ? 'Daftar akun baru' : 'Masuk ke akun Anda' }}</p>
      </div>
      
      <div v-if="errorMessage" class="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
        {{ errorMessage }}
      </div>
      
      <div v-if="successMessage" class="mb-4 bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded">
        {{ successMessage }}
      </div>
      
      <!-- Form Login -->
      <form v-if="!isRegisterMode" @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
          <input 
            v-model="username" 
            id="username" 
            name="username"
            type="text" 
            autocomplete="username"
            required 
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
          <input 
            v-model="password" 
            id="password" 
            name="password"
            type="password" 
            autocomplete="current-password"
            required 
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <button 
            type="submit" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            :disabled="isLoading"
          >
            <span v-if="isLoading" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Memproses...
            </span>
            <span v-else>Login</span>
          </button>
        </div>
        
        <div class="text-center text-sm mt-4">
          <p class="text-gray-600">
            Belum memiliki akun? 
            <button 
              type="button" 
              @click="toggleMode" 
              class="text-blue-600 hover:text-blue-800 font-medium">
              Daftar sekarang
            </button>
          </p>
        </div>
      </form>
      
      <!-- Form Register -->
      <form v-else @submit.prevent="handleRegister" class="space-y-6">
        <div>
          <label for="reg-username" class="block text-sm font-medium text-gray-700">Username</label>
          <input 
            v-model="registerData.username" 
            id="reg-username" 
            name="reg-username"
            type="text" 
            autocomplete="username"
            required 
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <label for="reg-name" class="block text-sm font-medium text-gray-700">Nama Lengkap</label>
          <input 
            v-model="registerData.name" 
            id="reg-name" 
            name="reg-name"
            type="text" 
            required 
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <label for="reg-password" class="block text-sm font-medium text-gray-700">Password</label>
          <input 
            v-model="registerData.password" 
            id="reg-password" 
            name="reg-password"
            type="password" 
            autocomplete="new-password"
            required 
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <label for="reg-confirm-password" class="block text-sm font-medium text-gray-700">Konfirmasi Password</label>
          <input 
            v-model="registerData.confirmPassword" 
            id="reg-confirm-password" 
            name="reg-confirm-password"
            type="password" 
            autocomplete="new-password"
            required 
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <button 
            type="submit" 
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
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
              class="text-blue-600 hover:text-blue-800 font-medium">
              Login
            </button>
          </p>
        </div>
      </form>

      <div v-if="checkingServer" class="mt-4 text-center text-gray-600">
        <p>Memeriksa koneksi server...</p>
        <div class="flex justify-center mt-2">
          <div class="animate-spin h-5 w-5 border-2 border-blue-500 rounded-full border-t-transparent"></div>
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
      activeServerPort: 8080
    }
  },
  async mounted() {
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

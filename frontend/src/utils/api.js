import axios from 'axios';

// Coba beberapa port secara berurutan
const tryPorts = async (ports) => {
  for (const port of ports) {
    const url = `http://localhost:${port}/api/ping`;
    try {
      await axios.get(url, { timeout: 1000 });
      return port;
    } catch (error) {
      console.log(`Port ${port} tidak merespon, mencoba port lain...`);
    }
  }
  // Jika semua port gagal, gunakan default
  console.log('Tidak dapat menemukan server yang aktif, menggunakan port default 8080');
  return 8080;
};

// API URL yang akan digunakan
let currentPort = 8080;

// Function untuk mendapatkan baseURL
const getBaseURL = () => `http://localhost:${currentPort}/api`;

// Buat instance axios dengan nilai awal
const api = axios.create({
  baseURL: getBaseURL(),
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: false,
  timeout: 5000
});

// Cek server yang aktif saat aplikasi dimulai
(async () => {
  try {
    currentPort = await tryPorts([8080, 8081, 3000]);
    // Update baseURL setelah menemukan port yang aktif
    api.defaults.baseURL = getBaseURL();
    console.log(`Terhubung ke server pada port ${currentPort}`);
  } catch (error) {
    console.error('Gagal terhubung ke server:', error);
  }
})();

// Add a request interceptor to attach the JWT token
api.interceptors.request.use(
  (config) => {
    // Update baseURL untuk setiap request
    config.baseURL = getBaseURL();
    
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Add a response interceptor to handle 401 errors
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.code === 'ECONNABORTED') {
      console.error('Koneksi timeout. Server mungkin tidak berjalan.');
    } else if (!error.response) {
      console.error('Tidak dapat terhubung ke server. Pastikan server berjalan.');
    } else if (error.response.status === 401) {
      // If unauthorized, redirect to login
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api; 
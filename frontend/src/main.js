import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueApexCharts from 'vue3-apexcharts'

// Import CSS files in correct order
import './assets/tailwind.css'
import './style.css'

const app = createApp(App)
app.use(router)
app.component('apexchart', VueApexCharts)
app.mount('#app')

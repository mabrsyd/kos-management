<script setup>
import DefaultLayout from './layouts/DefaultLayout.vue'
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const notif = ref('')

function setNotif(msg) {
  notif.value = msg
  setTimeout(() => notif.value = '', 2000)
}

// Cek apakah halaman saat ini memerlukan layout
const requiresLayout = computed(() => {
  return route.path !== '/login'
})
</script>

<template>
  <template v-if="requiresLayout">
    <DefaultLayout>
      <div v-if="notif" class="fixed top-20 right-4 bg-green-500 text-white px-4 py-2 rounded shadow z-50">{{ notif }}</div>
      <router-view :setNotif="setNotif" />
    </DefaultLayout>
  </template>
  <template v-else>
    <div v-if="notif" class="fixed top-4 right-4 bg-green-500 text-white px-4 py-2 rounded shadow z-50">{{ notif }}</div>
    <router-view :setNotif="setNotif" />
  </template>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

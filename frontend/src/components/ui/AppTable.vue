<template>
  <div class="overflow-x-auto">
    <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
      <thead class="bg-gray-50 dark:bg-slate-800/90 backdrop-blur-sm sticky top-0 z-10">
        <tr>
          <th v-for="col in columns" :key="col.key" class="px-4 md:px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
            {{ col.label }}
          </th>
          <th v-if="$slots.action" class="px-4 md:px-6 py-3 text-right text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
            Aksi
          </th>
        </tr>
      </thead>
      <tbody class="bg-white dark:bg-slate-900 divide-y divide-gray-200 dark:divide-gray-700">
        <tr v-for="(row, index) in data" :key="row.id || index" class="group hover:bg-gray-50 dark:hover:bg-slate-800/70 transition-all duration-200">
          <td v-for="col in columns" :key="col.key" class="px-4 md:px-6 py-3.5 text-sm">
            <slot :name="col.key" :row="row">
              {{ row[col.key] }}
            </slot>
          </td>
          <td v-if="$slots.action" class="px-4 md:px-6 py-3.5 text-right">
            <div class="opacity-80 group-hover:opacity-100 transition-opacity duration-200">
              <slot name="action" :row="row" />
            </div>
          </td>
        </tr>
        <tr v-if="data.length === 0">
          <td :colspan="$slots.action ? columns.length + 1 : columns.length" class="px-6 py-16 text-center">
            <div class="flex flex-col items-center justify-center text-gray-500 dark:text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mb-4 text-gray-300 dark:text-gray-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
              </svg>
              <slot name="empty">
                <p class="text-base font-medium">Tidak ada data yang tersedia</p>
                <p class="text-sm mt-1">Coba ubah filter atau tambahkan data baru</p>
              </slot>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
const props = defineProps({
  columns: Array,
  data: Array,
})
</script>

<style scoped>
th {
  position: relative;
  letter-spacing: 0.05em;
  font-weight: 600;
}

th:after {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 1px;
  @apply bg-gray-200 dark:bg-gray-700;
}

/* Subtle zebra-striping for better readability */
tbody tr:nth-child(even) {
  @apply bg-gray-50/50 dark:bg-slate-800/20;
}

/* Animation for hovering rows */
@keyframes pulse-subtle {
  0% { transform: scale(1); }
  50% { transform: scale(1.002); }
  100% { transform: scale(1); }
}

tbody tr:hover {
  animation: pulse-subtle 0.3s ease-in-out;
  @apply shadow-sm;
}
</style> 
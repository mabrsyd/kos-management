<template>
  <div class="overflow-x-auto rounded-xl shadow">
    <table class="min-w-full bg-white dark:bg-slate-900 rounded-xl">
      <thead class="bg-blue-100 dark:bg-slate-800 sticky top-0 z-10">
        <tr>
          <th v-for="col in columns" :key="col.key" class="p-3 text-left font-semibold text-blue-700 dark:text-blue-300 whitespace-nowrap">{{ col.label }}</th>
          <th v-if="$slots.action" class="p-3 text-left font-semibold text-blue-700 dark:text-blue-300">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in data" :key="row.id" class="hover:bg-blue-50 dark:hover:bg-slate-800 transition" :class="(data.indexOf(row) % 2 === 0) ? 'bg-white dark:bg-slate-900' : 'bg-blue-50 dark:bg-slate-800'">
          <td v-for="col in columns" :key="col.key" class="p-3 whitespace-nowrap">{{ row[col.key] }}</td>
          <td v-if="$slots.action" class="p-3">
            <slot name="action" :row="row" />
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
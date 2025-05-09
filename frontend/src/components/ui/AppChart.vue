<template>
  <div class="bg-white dark:bg-slate-900 rounded-xl shadow-lg p-6">
    <apexchart :options="options" :series="series" :type="chartType" :height="height" />
  </div>
</template>
<script setup>
import { ref, onMounted, watch } from 'vue'

const props = defineProps({
  chartType: { type: String, default: 'bar' },
  options: Object,
  series: Array,
  height: { type: [String, Number], default: 260 },
})
</script> 
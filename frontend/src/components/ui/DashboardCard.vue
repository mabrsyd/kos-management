<template>
  <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm p-6 border border-gray-100 dark:border-slate-700 hover:shadow-md transition-all duration-200">
    <div class="flex items-start">
      <div :class="`p-3 rounded-xl ${iconBgClass}`">
        <component :is="icon" :class="`w-6 h-6 ${iconColorClass}`" />
      </div>
      <div class="ml-4">
        <div class="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">{{ label }}</div>
        <div class="text-2xl font-bold text-gray-900 dark:text-white">
          <span>{{ displayValue }}</span>
        </div>
        <div v-if="description" class="mt-1 text-xs text-gray-500 dark:text-gray-400">
          {{ description }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, computed } from 'vue'

const props = defineProps({
  label: String,
  value: Number,
  icon: [Object, Function],
  color: { type: String, default: 'blue' },
  description: { type: String, default: '' }
})

// Computed colors based on the color prop
const iconBgClass = computed(() => {
  const colors = {
    blue: 'bg-blue-50 dark:bg-blue-900/30',
    green: 'bg-green-50 dark:bg-green-900/30',
    red: 'bg-red-50 dark:bg-red-900/30',
    yellow: 'bg-yellow-50 dark:bg-yellow-900/30',
    indigo: 'bg-indigo-50 dark:bg-indigo-900/30',
    purple: 'bg-purple-50 dark:bg-purple-900/30',
    pink: 'bg-pink-50 dark:bg-pink-900/30',
  }
  return colors[props.color] || colors.blue
})

const iconColorClass = computed(() => {
  const colors = {
    blue: 'text-blue-600 dark:text-blue-400',
    green: 'text-green-600 dark:text-green-400',
    red: 'text-red-600 dark:text-red-400',
    yellow: 'text-yellow-600 dark:text-yellow-400',
    indigo: 'text-indigo-600 dark:text-indigo-400',
    purple: 'text-purple-600 dark:text-purple-400',
    pink: 'text-pink-600 dark:text-pink-400',
  }
  return colors[props.color] || colors.blue
})

const displayValue = ref(0)

function animateValue() {
  let start = 0
  const end = props.value || 0
  const duration = 800
  const step = Math.ceil(end / (duration / 16))
  
  function update() {
    if (start < end) {
      start += step
      if (start > end) start = end
      displayValue.value = start
      requestAnimationFrame(update)
    } else {
      displayValue.value = end
    }
  }
  
  update()
}

onMounted(animateValue)
watch(() => props.value, animateValue)
</script> 
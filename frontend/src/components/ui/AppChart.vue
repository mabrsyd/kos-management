<template>
  <div class="bg-white dark:bg-slate-800 rounded-xl shadow-sm p-6 border border-gray-100 dark:border-slate-700 overflow-hidden">
    <div v-if="title" class="flex justify-between items-center mb-4">
      <h3 class="text-lg font-medium text-gray-700 dark:text-gray-300">{{ title }}</h3>
      <slot name="actions"></slot>
    </div>
    <div v-if="subtitle" class="text-sm text-gray-500 dark:text-gray-400 mb-4">{{ subtitle }}</div>
    <div :class="{'pt-2': title || subtitle}">
      <apexchart :options="mergedOptions" :series="series" :type="chartType" :height="height" :width="width" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue'

const props = defineProps({
  title: { type: String, default: '' },
  subtitle: { type: String, default: '' },
  chartType: { type: String, default: 'bar' },
  options: { type: Object, default: () => ({}) },
  series: { type: Array, default: () => [] },
  height: { type: [String, Number], default: 300 },
  width: { type: [String, Number], default: '100%' },
  theme: { type: String, default: 'light' }
})

// Merge default options with provided options
const defaultOptions = {
  chart: {
    fontFamily: 'inherit',
    toolbar: { show: false },
    zoom: { enabled: false },
    animations: { enabled: true },
  },
  grid: {
    borderColor: 'rgba(0, 0, 0, 0.05)',
    padding: { left: 0, right: 0, top: 0, bottom: 0 }
  },
  tooltip: {
    theme: 'light',
    enabled: true,
    shared: true
  },
  legend: {
    position: 'top',
    horizontalAlign: 'right',
    fontSize: '14px',
    fontFamily: 'inherit',
    offsetY: 8
  },
  colors: ['#3b82f6', '#10b981', '#ef4444', '#f59e0b', '#8b5cf6']
}

// Apply dark theme overrides if in dark mode
const darkModeOptions = {
  chart: {
    foreColor: '#9ca3af' // text-gray-400
  },
  grid: {
    borderColor: 'rgba(255, 255, 255, 0.1)'
  },
  tooltip: {
    theme: 'dark'
  }
}

// Merge options based on theme
const mergedOptions = computed(() => {
  const isDark = document.documentElement.classList.contains('dark');
  
  // Deep merge of default options and user options
  const baseOptions = { ...defaultOptions };
  
  // Apply dark mode options if needed
  if (isDark || props.theme === 'dark') {
    Object.assign(baseOptions, darkModeOptions);
  }
  
  // Merge with user options (with user options taking precedence)
  return Object.assign({}, baseOptions, props.options);
});

// Monitor theme changes in the DOM
onMounted(() => {
  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.attributeName === 'class' && 
          mutation.target === document.documentElement) {
        // Force refresh of computed property when theme changes
        mergedOptions.value = { ...mergedOptions.value };
      }
    });
  });
  
  observer.observe(document.documentElement, { attributes: true });
});
</script> 
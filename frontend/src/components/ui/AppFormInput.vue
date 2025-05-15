<template>
  <div class="mb-4">
    <label :for="id" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">{{ label }}</label>
    <div class="relative rounded-md shadow-sm">
      <div v-if="prefix" class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
        <span class="text-gray-500 dark:text-gray-400">{{ prefix }}</span>
      </div>
      <input
        :id="id"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        :type="type"
        :min="min"
        :max="max"
        :placeholder="placeholder"
        :required="required"
        :disabled="disabled"
        :class="[
          'block w-full rounded-lg transition-colors duration-200',
          'focus:ring-2 focus:ring-blue-500 focus:border-blue-500 focus:ring-opacity-30',
          'border border-gray-300 dark:border-slate-600 dark:bg-slate-800 dark:text-white',
          prefix ? 'pl-10' : 'pl-3',
          disabled ? 'opacity-70 cursor-not-allowed' : '',
          error ? 'border-red-300 focus:ring-red-500 focus:border-red-500 dark:border-red-600' : ''
        ]"
        class="py-2 pr-3"
      />
      <div v-if="error" class="mt-1 text-sm text-red-600 dark:text-red-400">{{ error }}</div>
      <div v-if="helperText && !error" class="mt-1 text-xs text-gray-500 dark:text-gray-400">{{ helperText }}</div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  modelValue: {
    type: [String, Number],
    default: ''
  },
  label: {
    type: String,
    required: true
  },
  id: {
    type: String,
    required: true
  },
  type: {
    type: String,
    default: 'text'
  },
  placeholder: {
    type: String,
    default: ''
  },
  prefix: {
    type: String,
    default: ''
  },
  required: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  },
  min: {
    type: [String, Number],
    default: null
  },
  max: {
    type: [String, Number],
    default: null
  },
  error: {
    type: String,
    default: ''
  },
  helperText: {
    type: String,
    default: ''
  }
})

defineEmits(['update:modelValue'])
</script>

<style scoped>
/* Custom styling for date inputs */
input[type="date"]::-webkit-calendar-picker-indicator {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='%236B7280'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z'/%3E%3C/svg%3E");
  cursor: pointer;
}

.dark input[type="date"]::-webkit-calendar-picker-indicator {
  filter: invert(0.8);
}

/* For webkit browsers */
input[type=number]::-webkit-inner-spin-button, 
input[type=number]::-webkit-outer-spin-button { 
  -webkit-appearance: none; 
  margin: 0; 
}

/* For Firefox */
input[type=number] {
  -moz-appearance: textfield;
}
</style> 
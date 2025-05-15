<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    :class="[
      'flex items-center justify-center gap-2 transition-all duration-200 font-medium rounded-lg',
      'focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-opacity-50',
      sizeClasses,
      {
        'bg-blue-600 hover:bg-blue-700 focus:ring-blue-500 text-white': variant === 'primary' && !disabled && !loading,
        'bg-red-600 hover:bg-red-700 focus:ring-red-500 text-white': variant === 'danger' && !disabled && !loading,
        'bg-gray-200 hover:bg-gray-300 focus:ring-gray-500 text-gray-700 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-gray-200': variant === 'secondary' && !disabled && !loading,
        'bg-green-600 hover:bg-green-700 focus:ring-green-500 text-white': variant === 'success' && !disabled && !loading,
        'bg-gray-300 text-gray-500 cursor-not-allowed dark:bg-gray-700 dark:text-gray-400': disabled || loading,
        'shadow-sm hover:shadow': !flat && !disabled && !loading,
        'border border-blue-500 bg-transparent text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/20 dark:text-blue-400': variant === 'outline' && !disabled && !loading,
      },
      { 'opacity-75': loading }
    ]"
    @click="onClick"
  >
    <svg v-if="loading" class="animate-spin -ml-1 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
    </svg>
    <slot></slot>
  </button>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: value => ['primary', 'secondary', 'danger', 'success', 'outline'].includes(value)
  },
  size: {
    type: String,
    default: 'md',
    validator: value => ['sm', 'md', 'lg'].includes(value)
  },
  type: {
    type: String,
    default: 'button'
  },
  disabled: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  },
  flat: {
    type: Boolean,
    default: false
  }
})

const emits = defineEmits(['click'])

const sizeClasses = computed(() => {
  switch (props.size) {
    case 'sm': return 'text-xs px-3 py-1.5'
    case 'lg': return 'text-base px-6 py-3'
    default: return 'text-sm px-4 py-2'
  }
})

const onClick = (event) => {
  if (!props.disabled && !props.loading) {
    emits('click', event)
  }
}
</script> 
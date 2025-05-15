<template>
  <transition name="fade">
    <div v-if="modelValue" class="fixed inset-0 z-50 overflow-y-auto" @click.self="close">
      <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
        <transition name="scale">
          <div v-if="modelValue" class="relative transform overflow-hidden rounded-xl bg-white dark:bg-slate-800 border border-gray-200 dark:border-slate-700 px-4 pb-4 pt-5 text-left shadow-xl transition-all w-full max-w-lg sm:p-6">
            <div class="absolute right-0 top-0 pr-4 pt-4">
              <button
                @click="close"
                class="rounded-md bg-white dark:bg-slate-800 text-gray-400 dark:text-gray-500 hover:text-gray-500 dark:hover:text-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                <span class="sr-only">Close</span>
                <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            
            <div class="mb-6">
              <h3 class="text-lg md:text-xl font-bold text-gray-900 dark:text-gray-100">
                {{ title }}
              </h3>
              <p v-if="subtitle" class="mt-1 text-sm text-gray-500 dark:text-gray-400">
                {{ subtitle }}
              </p>
            </div>
            
            <div class="mt-4 max-h-[70vh] overflow-y-auto">
              <slot></slot>
            </div>
            
            <div v-if="$slots.footer" class="mt-5 sm:mt-6 flex justify-end gap-3">
              <slot name="footer"></slot>
            </div>
          </div>
        </transition>
      </div>
    </div>
  </transition>
  
  <transition name="fade">
    <div v-if="modelValue" class="fixed inset-0 z-40 bg-black bg-opacity-50 backdrop-blur-sm transition-opacity"></div>
  </transition>
</template>

<script setup>
defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    required: true
  },
  subtitle: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])

const close = () => {
  emit('update:modelValue', false)
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.scale-enter-active,
.scale-leave-active {
  transition: all 0.3s ease;
}

.scale-enter-from,
.scale-leave-to {
  transform: scale(0.95);
  opacity: 0;
}
</style> 
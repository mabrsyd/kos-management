<template>
  <div :class="['rounded-xl shadow-lg p-6 flex flex-col items-center w-full transition', bgClass, textClass]">
    <component :is="icon" class="w-10 h-10 mb-2" />
    <div class="text-lg font-semibold mb-1">{{ label }}</div>
    <div class="text-3xl font-bold">
      <span>{{ displayValue }}</span>
    </div>
  </div>
</template>
<script setup>
import { ref, watch, onMounted } from 'vue'
const props = defineProps({
  label: String,
  value: Number,
  icon: [Object, Function],
  color: { type: String, default: 'blue' },
})
const bgClass = `bg-gradient-to-br from-${props.color}-400 to-${props.color}-600 dark:from-${props.color}-700 dark:to-${props.color}-900 text-white`
const textClass = 'text-white'
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
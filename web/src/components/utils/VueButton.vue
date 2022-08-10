<script setup>
// Vue
import { computed } from 'vue';
const props = defineProps({
  message: String,
  buttonType: {
    type: String,
    default: 'button',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})

const buttonClass = computed(() => {
  if (props.disabled) return ['disabled'];

  return [];
});
</script>

<template>
  <button :type="buttonType" :disabled="disabled" :class="buttonClass">
    {{ message }}
  </button>
</template>

<style scoped>
button {
  color: var(--vt-c-indigo);
  background: #FFFFFF;
  border: 1px solid var(--em-c-pink-1);
  border-radius: 5px;
  padding: 5px 20px;
  cursor: pointer;
}

button:hover {
  filter: drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25));
}

button.disabled {
  background-color: var(--em-c-disabled-button);
  cursor: not-allowed;
}
</style>
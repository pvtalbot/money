<script setup>
defineProps({
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
</script>

<template>
  <transition name="button-shadow">
    <button :type="buttonType" :disabled="disabled">
      {{ message }}
    </button>
  </transition>
</template>

<style scoped>
button {
  color: #002D4F;
  background: #FFFFFF;
  border: 1px solid #D01378;
  border-radius: 5px;
  padding: 5px 20px;
  cursor: pointer;
}

button:hover {
  filter: drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25));
}
</style>
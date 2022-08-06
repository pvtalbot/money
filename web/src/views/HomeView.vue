<script setup>
// Pinia
import { useUserStore } from '@/stores/user.js'
import { computed } from '@vue/reactivity';

const userStore = useUserStore()
const firstName = computed(() => userStore.user.firstName)
</script>

<template>
  <h1>Hello {{firstName}}!</h1>
</template>
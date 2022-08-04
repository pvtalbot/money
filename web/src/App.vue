<script setup>
import { RouterView } from 'vue-router'
import { useUserStore } from '@/stores/user.js'

const userStore = useUserStore();
</script>

<template>
  <RouterView />
</template>

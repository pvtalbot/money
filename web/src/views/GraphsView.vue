<script setup>
//Pinia
import { useUserStore } from '@/stores/user.js'

//Vue
import { computed } from 'vue';
import ExpensesSum from '@/components/graphs/ExpensesSum.vue';

const userStore = useUserStore();
const firstName = computed(() => userStore.user.firstName);
</script>

<template>
  <div class="graphs">
    <h1>Hello {{ firstName }}, this is sums!</h1>
    <ExpensesSum/>
  </div>
</template>

<style scoped>
</style>
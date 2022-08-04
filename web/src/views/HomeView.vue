<script setup>
import { useUserStore } from '@/stores/user.js';
import { computed } from 'vue';
import ExpensesList from '../components/ExpensesList.vue';

const userStore = useUserStore();
const firstName = computed(() => userStore.user.firstName)
</script>

<template>
  <div class="home">
    <h1>Hello {{firstName}}, this is Expenses manager!</h1>
  </div>
  <ExpensesList/>
</template>
<script setup>
import { useUserStore } from '@/stores/user.js';
import { useDrawerStore } from '@/stores/drawer.js';

import { computed, ref } from 'vue';
import ExpensesList from '@/components/expenses-list/ExpensesList.vue';
import VueButton from '@/components/utils/VueButton.vue';
import Drawer from '../components/utils/Drawer.vue';
import CreateExpenseForm from '../components/expenses-list/CreateExpenseForm.vue';

const userStore = useUserStore();
const drawerStore = useDrawerStore();
const firstName = computed(() => userStore.user.firstName)

const openCloseDrawer = () => {
  drawerStore.setDrawerOpen(!drawerStore.drawerOpen);
}
</script>

<template>
  <div class="home">
    <Drawer>
      <template #component>
        <CreateExpenseForm />
      </template>
    </Drawer>
    <h1>Hello {{firstName}}, this is Expenses manager!</h1>
    <div class="expenses-list__create-expense" @click="openCloseDrawer">
      <VueButton message="Add an expense"/>
    </div>
    <ExpensesList/>
  </div>
</template>

<style scoped>
.expenses-list__create-expense {
  margin: 0 15px;
  width: fit-content;
}
</style>
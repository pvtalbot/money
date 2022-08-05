<script setup>
import { useUserStore } from '@/stores/user.js';
import { useDrawerStore } from '@/stores/drawer.js';

import { computed, ref } from 'vue';
import ExpensesList from '@/components/home/ExpensesList.vue';
import VueButton from '@/components/utils/VueButton.vue';
import Drawer from '@/components/layout/Drawer.vue';
import CreateExpenseForm from '@/components/home/CreateExpenseForm.vue';

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
    <div class="create-expense" @click="openCloseDrawer">
      <VueButton message="Add an expense"/>
    </div>
    <ExpensesList/>
  </div>
</template>

<style scoped>
.create-expense {
  margin: 0 15px;
  width: fit-content;
}
</style>
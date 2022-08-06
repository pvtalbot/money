<script setup>
// Pinia
import { useUserStore } from '@/stores/user.js';
import { useDrawerStore } from '@/stores/drawer.js';

// Vue
import { computed } from 'vue';
import ExpensesList from '@/components/expenses/ExpensesList.vue';
import VueButton from '@/components/utils/VueButton.vue';
import CreateExpenseForm from '@/components/expenses/CreateExpenseForm.vue';

const userStore = useUserStore();
const drawerStore = useDrawerStore();

const firstName = computed(() => userStore.user.firstName)
const COMPONENT_TO_DRAWER = "CreateExpenseForm"

</script>

<template>
  <div class="expenses">
      <Teleport to="#teleport-component-to-drawer">
        <CreateExpenseForm v-if="drawerStore.isCurrentComponentDisplayed(COMPONENT_TO_DRAWER)"/>
      </Teleport>
    <h1>Hello {{firstName}}, this is Expenses manager!</h1>
    <div class="create-expense" @click="drawerStore.registerComponent(COMPONENT_TO_DRAWER)">
      <VueButton message="Add an expense" />
    </div>
    <ExpensesList />
  </div>
</template>

<style scoped>
.expenses {
  display: flex;
  flex-flow: column nowrap;
}

.create-expense {
  margin: auto;
  width: fit-content;
}
</style>
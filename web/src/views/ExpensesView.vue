<script setup>
// Pinia
import { useUserStore } from '@/stores/user.js';
import { useDrawerStore } from '@/stores/drawer.js';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Vue
import { computed, ref } from 'vue';
import ExpensesList from '@/components/expenses/ExpensesList.vue';
import RevenueList from '@/components/revenues/RevenueList.vue';
import VueButton from '@/components/utils/VueButton.vue';
import CreateExpenseForm from '@/components/expenses/CreateExpenseForm.vue';
import CreateRevenueForm from '@/components/revenues/CreateRevenueForm.vue';
import DatePicker from '@/components/utils/DatePicker.vue';

dayjs.extend(utc);
const userStore = useUserStore();
const drawerStore = useDrawerStore();

const mainDatePicker = ref(null)

const firstName = computed(() => userStore.user.firstName)
const COMPONENT_TO_DRAWER = "CreateExpenseForm";
const COMPONENT_TWO_TO_DRAWER = "CreateRevenueForm";

const initialDate = dayjs().utc().date(1).hour(0).minute(0).second(0).millisecond(0);
const dateToProp = computed(() => mainDatePicker == null || mainDatePicker.value == null ? initialDate : mainDatePicker.value.date);
</script>

<template>
  <h1>Hello {{firstName}}, this is Expenses manager!</h1>
  <DatePicker :initialDate="initialDate" ref="mainDatePicker"/>
  <div class="transfers">
    <Teleport to="#teleport-component-to-drawer">
      <Transition name="component">
        <CreateExpenseForm v-if="drawerStore.isCurrentComponentDisplayed(COMPONENT_TO_DRAWER)"/>
      </Transition>
    </Teleport>
    <Teleport to="#teleport-component-to-drawer">
      <Transition name="component">
        <CreateRevenueForm v-if="drawerStore.isCurrentComponentDisplayed(COMPONENT_TWO_TO_DRAWER)"/>
      </Transition>
    </Teleport>
    <div class="revenues">
      <div class="create-revenue">
        <VueButton message="Add a revenue" @click="drawerStore.registerComponent(COMPONENT_TWO_TO_DRAWER)"/>
      </div>
      <RevenueList :initialDate="dateToProp"/>
    </div>
    <div class="expenses">
      <div class="create-expense" @click="drawerStore.registerComponent(COMPONENT_TO_DRAWER)">
        <VueButton message="Add an expense" />
      </div>
      <ExpensesList :initialDate="dateToProp"/>
    </div>
  </div>
</template>

<style scoped>
.transfers {
  display: flex;
  flex-flow: row;
}

.revenues, .expenses {
  margin-left: auto;
  margin-right: auto;
}

.create-expense, .create-revenue {
  margin: auto;
  width: fit-content;
}

.component-leave-active, .component-enter-active {
  transition: all 0.5s ease;
}

.component-leave-to, .component-enter-from {
  transform: translateX(320px);
}
</style>
<script setup>
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

import { ref, computed } from 'vue';

import { useMutation } from '@vue/apollo-composable';
import CreateExpenseMutation from '@/graphql/mutations/CreateExpenseMutation.gql';

import { useExpenseStore } from '@/stores/expense.js'
import { useDrawerStore } from '@/stores/drawer';
import VueButton from '../utils/VueButton.vue';

dayjs.extend(utc);

const amount = ref(0);
const disabled = ref(false);
const date = ref(dayjs().format('YYYY-MM-DD'));
const formatedDate = computed(() => date.value + 'T00:00:00Z')
const expenseStore = useExpenseStore();
const drawerStore = useDrawerStore();

const { mutate: createExpenseMutation } = useMutation(CreateExpenseMutation)
const createExpense = function() {
  disabled.value = true;
  createExpenseMutation({
    input: {
      amount: amount.value,
      date: formatedDate.value,
    }
  }).then((r) => {expenseStore.updateExpenses([r.data.createExpense]); })
  .then(() => {drawerStore.setDrawerOpen(false);})
  .catch(e => {console.log(e);})
  .finally(() => {disabled.value = false;})
}

</script>

<template>
  <div class="create-expense-form">
    <h2>Create a new expense</h2>
    <form @submit.prevent="createExpense">
      <div class="create-expense-form__item-container">
        <label for="create-expense-form__amount">Amount</label>
        <input type="number"
        min=0
        step="1"
        id="create-expense-form__amount"
        v-model.number="amount"/>
      </div>
      <input class="create-expense-form__datepicker" type="date" v-model="date" />
      <VueButton button-type="submit" message="Submit" :disabled="disabled" />
    </form>
  </div>

</template>

<style scoped>
form {
  display: flex;
  flex-flow: column nowrap;
}

.create-expense-form {
  color: black;
}

.create-expense-form__item-container {
  padding: 0 0 10px 0;
  display: flex;
  flex-flow: column nowrap;
}

.create-expense-form__datepicker {
  margin-bottom: 10px;
}

label {
  align-self: start;
}
</style>
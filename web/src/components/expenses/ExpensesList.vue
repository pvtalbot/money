<script setup>
// Apollo
import { useMutation, useQuery } from '@vue/apollo-composable';
import Expenses from '@/graphql/queries/ExpenseList.gql';
import DeleteExpense from '@/graphql/mutations/DeleteExpenseMutation.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Pinia
import { useExpenseStore } from '@/stores/expense.js'

// Vue
import { computed, ref, watch } from 'vue'
import ExpenseCard from '@/components/expenses/ExpenseCard.vue';
import DatePicker from '@/components/utils/DatePicker.vue';

dayjs.extend(utc)
const expenseStore = useExpenseStore();

// Ref to get the date from the datepicker
const datePicker = ref(null)

// initial date, start date, end date
const initialDate = dayjs().utc().day(1).hour(0).minute(0).second(0).millisecond(0);
const startDate = computed(() => datePicker.value == null ? initialDate : datePicker.value.date);
const endDate = computed(() => startDate.value.add(1, 'month'))

const displayedExpenses = computed(() => expenseStore.getCurrentExpenses(startDate.value));
const sortedExpenses = computed(() => {
  return Object.values(displayedExpenses.value).sort((a, b) => {
    if (a.date.isBefore(b.date)) return 1;
    else if (a.date.isAfter(b.date)) return -1;

    return 0;
  })
})

// Apollo Query to get expenses list, hook, watcher
const { result: expenses, onResult: onExpenseListSucceeded, refetch: refetchExpenses, loading: expensesLoading} = 
  useQuery(Expenses, 
    { input: { startDate: startDate.value.toISOString(), endDate: endDate.value.toISOString()}},
  );

onExpenseListSucceeded(() => { expenseStore.updateExpenses(expenses.value.expenses); })

watch(startDate, () => {
  refetchExpenses({
    input: {
      startDate: startDate.value.toISOString(),
      endDate: endDate.value.toISOString(),
    }
  });
})

// Apollo Mutation to delete an expense
const {mutate: deleteExpenseMutation} = useMutation(DeleteExpense)
// Wrapper function for the mutation
const deleteExpense = function(expense) {
  deleteExpenseMutation({input: {id: expense.id}})
    .then(() => expenseStore.deleteExpense(expense))
    .catch(e => { console.log(e); });
}

</script>

<template>
  <div class="expenses-list">
    <DatePicker class="datepicker" ref="datePicker" :initialDate="initialDate"/>
      <div v-if="expensesLoading.value" class="expenses-list__loader" key="waiting">
        <h2>A minute please, I'm gathering everything!</h2>
      </div>
      <div v-else class="expenses-list__list" key="loaded">
        <p v-if="sortedExpenses.length > 0">Expenses of the month:</p>
        <div v-for="expense in sortedExpenses" :key="expense.id" class="expenses-list__expense">
          <ExpenseCard :expense="expense" @delete-expense="deleteExpense(expense)"/>
        </div>
      </div>
  </div>
</template>

<style scoped>
.expenses-list {
  display: flex;
  flex-flow: column nowrap;
}
.expenses-list__list {
  display: flex;
  flex-flow: column nowrap;
  max-height: 80vh;
  overflow: auto;
}

.expenses-list__list .expenses-list__expense, .expenses-list__list p {
  margin: auto;
}

</style>
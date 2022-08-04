<script setup>
// External libraries
import dayjs from 'dayjs';

// Apollo
import { useMutation, useQuery } from '@vue/apollo-composable';
import Expenses from '@/graphql/queries/ExpenseList.gql';
import DeleteExpense from '@/graphql/mutations/DeleteExpenseMutation.gql';

// Pinia
import { useExpenseStore } from '@/stores/expense.js'

// Vue
import { computed, ref } from 'vue'
import ExpenseCard from '@/components/expenses-list/ExpenseCard.vue';
import MonthPicker from '@/components/utils/MonthPicker.vue';

const expenseStore = useExpenseStore();

const monthPicker = ref(null)

const startDate = ref(dayjs().day(1).hour(0).minute(0).second(0).millisecond(0))
const endDate = computed(() => startDate.value.add(1, 'month'))

const displayedExpenses = computed(() => expenseStore.getCurrentExpenses(startDate.value));
const sortedExpenses = computed(() => {
  return Object.values(displayedExpenses.value).sort((a, b) => {
    if (a.date.isBefore(b.date)) return 1;
    else if (a.date.isAfter(b.date)) return -1;

    return 0;
  })
})


const { result: expenses, onResult: onExpenseListSucceeded, loading: expensesLoading,
  } = 
  useQuery(Expenses, {
  input: {
    startDate: startDate.value.toISOString(),
    endDate: endDate.value.toISOString(),
  }
});

onExpenseListSucceeded(() => {
  expenseStore.updateExpenses(expenses.value.expenses);
})

const {mutate: deleteExpenseMutation} = useMutation(DeleteExpense)
const deleteExpense = function(expense) {
  deleteExpenseMutation({input: {id: expense.id}})
    .then(() => expenseStore.deleteExpense(expense))
    .catch(e => { console.log(e); });
}


</script>

<template>
  <div class="expenses-list">
    <MonthPicker ref="monthPicker"/>
    <transition name="slide-fade" tag="div" mode="out-in">
      <div v-if="expensesLoading" class="expenses-list__loader" key="waiting">
        <h1>A minute please, I'm gathering everything!</h1>
      </div>
      <div v-else class="expenses_list__list" key="loaded">
        <p>Expenses of the month:</p>
        <div v-for="expense in sortedExpenses" :key="expense.id" class="expenses-list__expense">
          <ExpenseCard :expense="expense" @delete-expense="deleteExpense(expense)"/>
        </div>
      </div>
    </transition>
  </div>
</template>
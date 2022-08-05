<script setup>
// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Apollo
import { useMutation, useQuery } from '@vue/apollo-composable';
import Expenses from '@/graphql/queries/ExpenseList.gql';
import DeleteExpense from '@/graphql/mutations/DeleteExpenseMutation.gql';

// Pinia
import { useExpenseStore } from '@/stores/expense.js'

// Vue
import { computed, ref, watch } from 'vue'
import ExpenseCard from '@/components/home/ExpenseCard.vue';
import DatePicker from '@/components/utils/DatePicker.vue';

dayjs.extend(utc)

const expenseStore = useExpenseStore();
const datePicker = ref(null)

const initialDate = dayjs().utc().day(1).hour(0).minute(0).second(0).millisecond(0);

const startDate = computed(() => {
  if (datePicker.value == null) return initialDate;

  return datePicker.value.date;
})
const endDate = computed(() => startDate.value.add(1, 'month'))

const displayedExpenses = computed(() => expenseStore.getCurrentExpenses(startDate.value));
const sortedExpenses = computed(() => {
  return Object.values(displayedExpenses.value).sort((a, b) => {
    if (a.date.isBefore(b.date)) return 1;
    else if (a.date.isAfter(b.date)) return -1;

    return 0;
  })
})

const { result: expenses, onResult: onExpenseListSucceeded, refetch: refetchExpenses, loading: expensesLoading} = 
  useQuery(Expenses, 
    { input: { startDate: startDate.value.toISOString(), endDate: endDate.value.toISOString()}},
  );

onExpenseListSucceeded(() => { expenseStore.updateExpenses(expenses.value.expenses); })

const {mutate: deleteExpenseMutation} = useMutation(DeleteExpense)
const deleteExpense = function(expense) {
  deleteExpenseMutation({input: {id: expense.id}})
    .then(() => expenseStore.deleteExpense(expense))
    .catch(e => { console.log(e); });
}

watch(startDate, () => {
  refetchExpenses({
    input: {
      startDate: startDate.value.toISOString(),
      endDate: endDate.value.toISOString(),
    }
  });
})
</script>

<template>
  <div class="expenses-list">
    <DatePicker ref="datePicker" :initialDate="initialDate"/>
    <transition name="slide-fade" tag="div" mode="out-in">
      <div v-if="expensesLoading.value" class="expenses-list__loader" key="waiting">
        <h2>A minute please, I'm gathering everything!</h2>
      </div>
      <div v-else class="expenses_list__list" key="loaded">
        <p v-if="sortedExpenses.length > 0">Expenses of the month:</p>
        <div v-for="expense in sortedExpenses" :key="expense.id" class="expenses-list__expense">
          <ExpenseCard :expense="expense" @delete-expense="deleteExpense(expense)"/>
        </div>
      </div>
    </transition>
  </div>
</template>
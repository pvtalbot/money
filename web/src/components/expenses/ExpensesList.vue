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
import { useDrawerStore } from '@/stores/drawer.js';

// Vue
import { computed, ref, watch } from 'vue'
import ExpenseCard from '@/components/expenses/ExpenseCard.vue';
import UpdateExpenseForm from '@/components/expenses/UpdateExpenseForm.vue';

dayjs.extend(utc);
const expenseStore = useExpenseStore();
const drawerStore = useDrawerStore();

const props = defineProps({
  initialDate: {
    required: true,
  }
})

// initial date, start date, end date
const startDate = computed(() => props.initialDate);
const endDate = computed(() => startDate.value.add(1, 'month'))

const expenseToUpdate = ref(null);
const COMPONENT_TO_DRAWER = 'UpdateExpenseForm';

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

onExpenseListSucceeded(() => { expenseStore.cacheExpenses(expenses.value.expenses); })

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

const updateExpense = (expense) => {
  drawerStore.registerComponent(COMPONENT_TO_DRAWER);
  expenseToUpdate.value = expense;
}

</script>

<template>
  <div class="expenses-list">
    <div v-if="expensesLoading.value" class="expenses-list__loader" key="waiting">
      <h2>A minute please, I'm gathering everything!</h2>
    </div>
    <div v-else class="expenses-list__list" key="loaded">
      <p v-if="sortedExpenses.length > 0">Expenses of the month:</p>
      <div v-for="expense in sortedExpenses" :key="expense.id" class="expenses-list__expense">
        <ExpenseCard :expense="expense" 
                      @delete-expense="deleteExpense(expense)"
                      @click.self="updateExpense(expense)" />
      </div>
    </div>
    <Teleport to="#teleport-component-to-drawer">
      <Transition name="component">
        <UpdateExpenseForm v-if="drawerStore.isCurrentComponentDisplayed(COMPONENT_TO_DRAWER)" :expense="expenseToUpdate"/>
      </Transition>
    </Teleport>
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

.component-leave-active, .component-enter-active {
  transition: all 0.5s ease;
}

.component-leave-to, .component-enter-from {
  transform: translateX(320px);
}
</style>
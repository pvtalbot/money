<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import CreateExpenseMutation from '@/graphql/mutations/CreateExpenseMutation.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Pinia
import { useExpenseStore } from '@/stores/expense.js'
import { useDrawerStore } from '@/stores/drawer';

// Vue
import { ref, computed } from 'vue';
import VueButton from '../utils/VueButton.vue';

dayjs.extend(utc);
const expenseStore = useExpenseStore();
const drawerStore = useDrawerStore();

const amount = ref(0);
const date = ref(dayjs().format('YYYY-MM-DD'));
const categoryId = ref(null)
const formatedDate = computed(() => date.value + 'T00:00:00Z')
const disabled = ref(false);

const getCategories = computed(() => expenseStore.expensesCategories)

// Apollo Mutation to create an expense
const { mutate: createExpenseMutation } = useMutation(CreateExpenseMutation)

// Wrapper function
const createExpense = function() {
  disabled.value = true;
  createExpenseMutation({
    input: {
      amount: amount.value,
      date: formatedDate.value,
      categoryId: categoryId.value,
    }
  }).then(r => {
    expenseStore.cacheExpenses([r.data.createExpense]); 
    drawerStore.closeDrawer();
    })
  .catch(e => {console.log(e);})
  .finally(() => {disabled.value = false;})
}

</script>

<template>
  <div class="create-expense-form">
    <h2>Create a new expense</h2>
    <form @submit.prevent="createExpense">
      <div class="item-container">
        <label for="amount">Amount</label>
        <input type="number"
        min=0
        step="1"
        id="create-expense-form__amount"
        v-model.number="amount"/>
      </div>
      <div class="item-container">
        <label for="create-expense-form__date">Date</label>
        <input class="create-expense-form__datepicker" type="date" v-model="date" id="create-expense-form__date" />
      </div>
      <div class="item-container">
        <label for="create-expense-form__category">Category</label>
        <select class="select-category" id="create-expense-form__category" v-model="categoryId" required>
          <option v-for="c in getCategories" :key="c.id" :value="c.id">{{ c.name }}</option>
        </select>
      </div>
      <VueButton button-type="submit" message="Submit" :disabled="disabled" class="submit-button"/>
    </form>
  </div>

</template>

<style scoped>
form {
  display: flex;
  flex-flow: column nowrap;
}

.item-container {
  padding: 0 0 10px 0;
  display: flex;
  flex-flow: column nowrap;
}

.select-category {
  background-color: var(--vt-c-white);
}

label {
  align-self: start;
}

.submit-button {
  margin-top: 10px;
}
</style>
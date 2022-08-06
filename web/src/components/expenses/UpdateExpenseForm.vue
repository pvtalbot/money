<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import UpdateExpenseMutation from '@/graphql/mutations/UpdateExpenseMutation.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Pinia
import { useExpenseStore } from '@/stores/expense.js'
import { useDrawerStore } from '@/stores/drawer.js';

// Vue
import VueButton from '@/components/utils/VueButton.vue';
import { ref, onMounted } from 'vue';
import { computed } from '@vue/reactivity';

dayjs.extend(utc);
const expenseStore = useExpenseStore();
const drawerStore = useDrawerStore();

const props = defineProps({
  expense: {
    required: true,
  }
})

const amount = ref(0)
const date = ref(null)
const formatedDate = computed(() => date.value + 'T00:00:00Z')
const disabled = ref(false)
onMounted(() => {
  amount.value = props.expense.amount;
  date.value = props.expense.date.format('YYYY-MM-DD');
})

// Apollo Mutation to update an expense
const { mutate: updateExpenseMutation } = useMutation(UpdateExpenseMutation)
//Wrapper function
const updateExpense = () => {
  disabled.value = true;
  updateExpenseMutation({
    input: {
      id: props.expense.id,
      amount: amount.value,
      date: formatedDate.value,
    }
  }).then(r => {expenseStore.cacheExpenses([r.data.updateExpense])})
  .then(() => expenseStore.deleteExpense(props.expense))
  .then(() => {drawerStore.closeDrawer();})
  .catch(e => {console.log(e);})
  .finally(() => {disabled.value = false;})
}

</script>

<template>
  <div class="update-expense-form">
    <h2>Update expense</h2>
    <form @submit.prevent="updateExpense">
      <div class="item-container">
        <label for="amount">Amount</label>
        <input type="number"
        min="0"
        step="1"
        id="update-expense-form__amount"
        v-model="amount"/>
      </div>
      <input class="datepicker" type="date" v-model="date"/>
      <VueButton button-type="submit" message="Submit" :disabled="disabled"/>
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

.datepicker {
  margin-bottom: 10px;
}

label {
  align-self: start;
}
</style>
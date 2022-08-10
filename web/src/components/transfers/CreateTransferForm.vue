<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import CreateRevenueMutation from '@/graphql/mutations/CreateRevenueMutation.gql';
import CreateExpenseMutation from '@/graphql/mutations/CreateExpenseMutation.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Pinia
import { useRevenueStore } from '@/stores/revenue';
import { useExpenseStore } from '@/stores/expense';
import { useDrawerStore } from '@/stores/drawer';

// Vue
import { ref, computed } from 'vue';
import VueButton from '@/components/utils/VueButton.vue';

const props = defineProps({
  mode: {
    type: String,
    required: true,
    validator: value => ['expense', 'revenue'].includes(value),
  }
})

dayjs.extend(utc);
const drawerStore = useDrawerStore();

const amount = ref(0);
const date = ref(dayjs().format('YYYY-MM-DD'));
const formatedDate = computed(() => date.value + 'T00:00:00Z');
const disabled = ref(false);
const categoryId = ref(null);

const config = (() => {
  if (props.mode == 'expense') {
    const store = useExpenseStore();
    return {
      getCategories: computed(() => store.expensesCategories),
      mutation: CreateExpenseMutation,
      cacheTransfer: store.cacheExpenses,
      resultName: 'createExpense',
    }
  } else {
    const store = useRevenueStore();
    return {
      getCategories: null,
      mutation: CreateRevenueMutation,
      cacheTransfer: store.cacheRevenues,
      resultName: 'createRevenue',
    }
  }
})();

const getInput = () => {
  const input = {amount: amount.value, date: formatedDate.value}
  if (props.mode == 'expense') input['categoryId'] = categoryId.value;
  return input;
}

// Apollo mutation to create a revenue
const { mutate: createTransferMutation } = useMutation(config.mutation);
// Wrapper function
const createTransfer = () => {
  disabled.value = true;
  createTransferMutation({
    input: getInput()
  }).then(r => {
    config.cacheTransfer([r.data[config.resultName]]);
    drawerStore.closeDrawer();
  }).catch(e => {console.log(e);})
  .finally(() => {disabled.value = false;})

}
</script>

<template>
  <div class="create-revenue-form">
    <h2>Create a new {{ mode }}</h2>
    <form @submit.prevent="createTransfer">
      <div class="item-container">
        <label for="create-revenue-form__amount">Amount</label>
        <input type="number" min=0 step="1" id="create-revenue-form__amount" v-model.number="amount" />
      </div>
      <div class="item-container">
        <label for="create-expense-form__date">Date</label>
        <input class="datepicker" type="date" id="create-expense-form__date" v-model="date" />
      </div>
      <div class="item-container" v-if="mode == 'expense'">
        <label for="create-expense-form__cagegory">Category</label>
        <select class="select-category" id="create-expense-form__category" v-model="categoryId" required>
          <option v-for="c in config.getCategories.value" :key="c.id" :value="c.id">{{ c.name }}</option>
        </select>
      </div>
      <VueButton button-type="submit" message="Submit" :disabled="disabled" class="submit-button"></VueButton>
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
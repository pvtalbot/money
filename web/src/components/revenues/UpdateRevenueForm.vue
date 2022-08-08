<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import UpdateExpenseMutation from '@/graphql/mutations/UpdateExpenseMutation.gql';
import UpdateRevenueMutation from '@/graphql/mutations/UpdateRevenueMutation.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Pinia
import { useExpenseStore } from '@/stores/expense.js';
import { useRevenueStore } from '@/stores/revenue.js';
import { useDrawerStore } from '@/stores/drawer';

// Vue
import VueButton from '@/components/utils/VueButton.vue';
import { ref, computed, onMounted } from 'vue';

dayjs.extend(utc);

const props = defineProps({
  mode: {
    type: String,
    required: true,
    validator: (value) => ['expense', 'revenue'].includes(value),
  },
  updatedObject : {
    required: true,
  }
});

let getCategories;
const store = props.mode == 'expense' ? useExpenseStore() : useRevenueStore();
const drawerStore = useDrawerStore();

const amount = ref(0);
const date = ref(null);
const formatedDate = computed(() => date.value + 'T00:00:00Z');
const categoryId = ref(null);
const disabled = ref(false);

onMounted(() => {
  amount.value = props.updatedObject.amount;
  date.value = props.updatedObject.date.format('YYYY-MM-DD');

  if (props.mode == 'expense') {
    categoryId.value = props.updatedObject.categoryId;
    getCategories = computed(() => store.expensesCategories);
  }
});

const getInput = computed(() => {
  const input = {
    id: props.updatedObject.id,
    amount: amount.value,
    date: formatedDate.value,
  };

  if (props.mode == 'expense') {
    input['categoryId'] == categoryId.value;
  }

  return input;
})
const mutation = props.mode == 'expense' ? UpdateExpenseMutation : UpdateRevenueMutation;
const deleteObject = props.mode == 'expense' ? store.deleteExpense : store.deleteRevenue;
const cacheObject = props.mode == 'expense' ? store.cacheExpenses : store.cacheRevenues;
const resultName = props.mode == 'expense' ? 'updateExpense' : 'updateRevenue';

const { mutate: updateObjectMutation } = useMutation(mutation);
// Wrapper function
const update = () => {
  disabled.value = true;
  updateObjectMutation({ input: getInput.value })
    .then(r => {
      deleteObject(props.updatedObject);
      cacheObject([r.data[resultName]]);
      drawerStore.closeDrawer();
    })
    .catch(e => { console.log(e); })
    .finally(() => {disabled.value = false;})
  ;
}
</script> 

<template>
  <div class="update-object-form">
    <h2>Update</h2>
    <form @submit.prevent="update">
      <div class="item-container">
        <label for="update-object-form__amount">Amount</label>
        <input type="number" 
        min="0" 
        step="1" 
        id="update-object-form__amount" 
        v-model="amount" />
      </div>
      <div class="item-container">
        <label for="update-object-form__date">Date</label>
        <input class="datepicker" type="date" v-model="date" id="update-object-form__date"/>
      </div>
      <div class="item-container" v-if="props.mode =='expense'">
        <label for="update-object-form__category">Category</label>
        <select class="select-category" id="update-object-form__category" v-model="categoryId">
          <option v-for="c in getCategories" :key="c.id" :value="c.id">{{ c.name }}</option>
        </select>
      </div>
      <VueButton button-type="submit" message="Submit" :disabled="disabled" class="submit-butotn"/>
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

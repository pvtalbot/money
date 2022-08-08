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

const drawerStore = useDrawerStore();

const amount = ref(0);
const date = ref(null);
const formatedDate = computed(() => date.value + 'T00:00:00Z');
const categoryId = ref(null);
const disabled = ref(false);

// Link refs to props
onMounted(() => {
  amount.value = props.updatedObject.amount;
  date.value = props.updatedObject.date.format('YYYY-MM-DD');

  if (props.mode == 'expense') {
    categoryId.value = props.updatedObject.categoryId;
  }
});

const getInput = computed(() => {
  const input = { id: props.updatedObject.id, amount: amount.value, date: formatedDate.value };

  if (props.mode == 'expense') {
    input['categoryId'] = categoryId.value;
  }
  return input;
})

const config = (function() {
  if (props.mode == 'expense') {
    const store = useExpenseStore();
    return {
      store: store,
      getCategories: computed(() => store.expensesCategories),
      mutation: UpdateExpenseMutation,
      deleteObject: store.deleteExpense,
      cacheObject: store.cacheExpenses,
      resultName: 'updateExpense',
    };
  } else {
    const store = useRevenueStore();
    return {
      store: store,
      getCategories: null,
      mutation: UpdateRevenueMutation,
      deleteObject: store.deleteRevenue,
      cacheObject: store.cacheRevenues,
      resultName: 'updateRevenue',
    };
  }
})();

// Apollo mutation to update the object
const { mutate: updateObjectMutation } = useMutation(config.mutation);
// Wrapper function
const update = () => {
  disabled.value = true;
  updateObjectMutation({ input: getInput.value })
    .then(r => {
      config.deleteObject(props.updatedObject);
      config.cacheObject([r.data[config.resultName]]);
      drawerStore.closeDrawer();
    })
    .catch(e => { console.log(e); })
    .finally(() => {disabled.value = false;})
  ;
}
</script> 

<template>
  <div class="update-object-form">
    <h2>Update {{ props.mode }}</h2>
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
          <option v-for="c in config.getCategories.value" :key="c.id" :value="c.id">{{ c.name }}</option>
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

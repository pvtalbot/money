<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import UpdateExpenseMutation from '@/graphql/mutations/UpdateExpenseMutation.gql';
import UpdateRevenueMutation from '@/graphql/mutations/UpdateRevenueMutation.gql';
import DeleteExpenseMutation from '@/graphql/mutations/DeleteExpenseMutation.gql';
import DeleteRevenueMutation from '@/graphql/mutations/DeleteRevenueMutation.gql';

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
    validator: value => ['expense', 'revenue'].includes(value),
  },
  updatedTransfer : {
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
  amount.value = props.updatedTransfer.amount;
  date.value = props.updatedTransfer.date.format('YYYY-MM-DD');

  if (props.mode == 'expense') {
    categoryId.value = props.updatedTransfer.categoryId;
  }
});

const getInput = computed(() => {
  const input = { id: props.updatedTransfer.id, amount: amount.value, date: formatedDate.value };

  if (props.mode == 'expense') {
    input['categoryId'] = categoryId.value;
  }
  return input;
});

const config = (function() {
  if (props.mode == 'expense') {
    const store = useExpenseStore();
    return {
      store: store,
      getCategories: computed(() => store.expensesCategories),
      mutation: UpdateExpenseMutation,
      deleteMutation: DeleteExpenseMutation,
      deleteTransfer: store.deleteExpense,
      cacheTransfer: store.cacheExpenses,
      resultName: 'updateExpense',
    };
  } else {
    const store = useRevenueStore();
    return {
      store: store,
      getCategories: null,
      mutation: UpdateRevenueMutation,
      deleteMutation: DeleteRevenueMutation,
      deleteTransfer: store.deleteRevenue,
      cacheTransfer: store.cacheRevenues,
      resultName: 'updateRevenue',
    };
  }
})();

// Apollo mutation to update the transfer
const { mutate: updateTransferMutation, onDone: onUpdateSuccess, onError: onUpdateFailed } = useMutation(config.mutation);
// Wrapper function
const update = () => {
  disabled.value = true;
  updateTransferMutation({ input: getInput.value })
    .catch(e => { console.log(e); })
    .finally(() => {disabled.value = false;})
  ;
}
onUpdateSuccess(({data}) => {
  config.deleteTransfer(props.updatedTransfer);
  config.cacheTransfer([data[config.resultName]]);
  drawerStore.closeDrawer();
})
onUpdateFailed(e => {console.log(e)});

// Apollo mutation to delete the transfer
const { mutate: deleteTransferMutation, onDone: onDeleteSuccess, onError: onDeleteFailed } = useMutation(config.deleteMutation);
// Wrapper function for the mutation
const deleteTransfer = () => {
  deleteTransferMutation({input: {id: props.updatedTransfer.id}});
}
onDeleteSuccess(() => {
  config.deleteTransfer(props.updatedTransfer);
  drawerStore.closeDrawer();
})
onDeleteFailed(e => {console.log(e);})
</script> 

<template>
  <div class="update-transfer-form">
    <h2>Update {{ props.mode }}</h2>
    <form @submit.prevent="update">
      <div class="item-container">
        <label for="update-transfer-form__amount">Amount</label>
        <input type="number" min="0" step="1" id="update-transfer-form__amount" v-model="amount" />
      </div>
      <div class="item-container">
        <label for="update-transfer-form__date">Date</label>
        <input class="datepicker" type="date" v-model="date" id="update-transfer-form__date" />
      </div>
      <div class="item-container" v-if="props.mode =='expense'">
        <label for="update-transfer-form__category">Category</label>
        <select class="select-category" id="update-transfer-form__category" v-model="categoryId">
          <option v-for="c in config.getCategories.value" :key="c.id" :value="c.id">{{ c.name }}</option>
        </select>
      </div>
      <div class="buttons-container">
        <VueButton class="button-delete" button-type="button" message="Delete" @click.stop="deleteTransfer" />
        <VueButton button-type="submit" message="Submit" :disabled="disabled" class="submit-button" />
      </div>
    </form>
  </div>
</template>

<style scoped>
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

.buttons-container {
  margin-top: 10px;
  display: flex;
  flex-flow: row wrap;
}

.buttons-container button {
  margin: auto;
}

.button-delete {
  border: 0;
}
</style>

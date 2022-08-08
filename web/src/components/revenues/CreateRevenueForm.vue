<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import CreateRevenueMutation from '@/graphql/mutations/CreateRevenueMutation.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Pinia
import { useRevenueStore } from '@/stores/revenue';
import { useDrawerStore } from '@/stores/drawer';

// Vue
import { ref, computed } from 'vue';
import VueButton from '@/components/utils/VueButton.vue';

dayjs.extend(utc);
const revenueStore = useRevenueStore();
const drawerStore = useDrawerStore();

const amount = ref(0);
const date = ref(dayjs().format('YYYY-MM-DD'));
const formatedDate = computed(() => date.value + 'T00:00:00Z');
const disabled = ref(false);

// Apollo mutation to create a revenue
const { mutate: createRevenueMutation } = useMutation(CreateRevenueMutation);
// Wrapper function
const createRevenue = () => {
  disabled.value = true;
  createRevenueMutation({
    input: {
      amount: amount.value,
      date: formatedDate.value,
    }
  }).then(r => {
    revenueStore.cacheRevenues([r.data.createRevenue]);
    drawerStore.closeDrawer();
  }).catch(e => {console.log(e);})
  .finally(() => {disabled.value = false;})

}
</script>

<template>
  <div class="create-revenue-form">
    <h2>Create a new revenue</h2>
    <form @submit.prevent="createRevenue">
      <div class="item-container">
        <label for="create-revenue-form__amount">Amount</label>
        <input type="number"
        min=0
        step="1"
        id="create-revenue-form__amount"
        v-model.number="amount"/>
      </div>
      <div class="item-container">
        <label for="create-expense-form__date">Date</label>
        <input class="datepicker" type="date" id="create-expense-form__date" v-model="date" />
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

label {
  align-self: start;
}

.submit-button {
  margin-top: 10px;
}
</style>
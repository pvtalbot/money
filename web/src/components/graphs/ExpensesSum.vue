<script setup>
// Apollo
import { useQuery } from '@vue/apollo-composable';
import ExpensesSum from '@/graphql/queries/ExpensesSum.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Vue
import { computed, ref, watch } from 'vue';
import DatePicker from '@/components/utils/DatePicker.vue';

dayjs.extend(utc);

// Ref to get the date from the datepicker
const datePicker = ref(null);

// initial date, start date, end date
const initialDate = dayjs((new Date()).getFullYear() + '-01-01T00:00:00Z').utc()
const startDate = computed(() => datePicker.value == null ? initialDate : datePicker.value.date )
const endDate = computed(() => startDate.value.add(1, 'year'));


// Apollo Query from expenses sum, hook, watcher
const {result: expensesSum, refetch: refetchSums, loading: loadingSums, onResult: onExpensesSumSucceed} = useQuery(ExpensesSum, {
  input: {startDate: startDate.value.toISOString(), endDate: endDate.value.toISOString(), groupBy: 'MONTH'}
}, {fetchPolicy: 'no-cache'});

onExpensesSumSucceed(() => cleanAndSort(expensesSum))

watch(startDate, () => {
  refetchSums({
    input: {
      startDate: startDate.value.toISOString(),
      endDate: endDate.value.toISOString(),
      groupBy: 'MONTH',
    }
  });
})

// helper function, transforms the result from Apollo into something usable + sort
const cleanAndSort = function(s) {
  const res = []

  if (!s.value.expensesSum) return res;
  for (const e in s.value.expensesSum) {
    res.push({
      startDate: dayjs(s.value.expensesSum[e].startDate).utc(),
      endDate: dayjs(s.value.expensesSum[e].endDate).utc(),
      amount: s.value.expensesSum[e].amount,
    });
  }

  res.sort((a, b) => {
    if (a.startDate.isBefore(b.startDate)) return -1;
    if (a.startDate.isAfter(b.startDate)) return 1;

    return 0;
  });

  s.value = res;
}
</script>

<template>
  <div class="expenses-sum">
    <DatePicker ref="datePicker" mode="YEAR" :initialDate="initialDate"/>
    <transition name="slide-fade" tag="div" mode="out-in">
      <div v-if="!loadingSums" class="expenses-sum__sum" key="loaded">
        <p>Expenses of the year:</p>
        <div v-for="expense in expensesSum" key="expense.startDate">
          {{ expense.startDate.format('MMMM') }} : {{ expense.amount }} 
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.expenses-sum {
  display: flex;
  flex-flow: column nowrap;
}

.expenses-sum__sum {
  margin: auto;
}

</style>
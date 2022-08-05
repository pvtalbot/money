<script setup>
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

import { computed, ref, watch } from 'vue';

import { useQuery } from '@vue/apollo-composable';
import ExpensesSum from '@/graphql/queries/ExpensesSum.gql';
import DatePicker from '@/components/utils/DatePicker.vue';

dayjs.extend(utc);

const datePicker = ref(null);
const initialDate = dayjs((new Date()).getFullYear() + '-01-01T00:00:00Z').utc()

const startDate = computed(() => {
  if (datePicker.value == null) return initialDate;

  return datePicker.value.date;
})
const endDate = computed(() => startDate.value.add(1, 'year'));


const {result: expensesSum, refetch: refetchSums, loading: loadingSums, onResult: onExpensesSumSucceed} = useQuery(ExpensesSum, {
  input: {startDate: startDate.value.toISOString(), endDate: endDate.value.toISOString(), groupBy: 'MONTH'}
}, {fetchPolicy: 'no-cache'});

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
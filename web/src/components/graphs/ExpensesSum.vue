<script setup>
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

import { computed, ref } from 'vue';

import { useQuery } from '@vue/apollo-composable';
import ExpensesSum from '@/graphql/queries/ExpensesSum.gql';

dayjs.extend(utc);

const startDate = ref(dayjs((new Date()).getFullYear() + '-01-01T00:00:00Z').utc());
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

//onExpensesSumSucceed(() => {listToDisplay.value = cleanAndSort(expensesSum.value)})
onExpensesSumSucceed(() => cleanAndSort(expensesSum))

</script>

<template>
  <div class="expenses-sum">
    <transition name="slide-fade" tag="div" mode="out-in">
      <div v-if="loadingSums" class="expenses-list__loader" key="waiting">
        <h1>A minute please, I'm gathering everything!</h1>
      </div>
      <div v-else class="expenses-sum__sum" key="loaded">
        <p>Expenses of the year:</p>
        <div v-for="expense in expensesSum" key="expense.startDate">
          {{ expense.startDate.format('MMMM') }} : {{ expense.amount }} 
        </div>
      </div>
    </transition>
  </div>
</template>
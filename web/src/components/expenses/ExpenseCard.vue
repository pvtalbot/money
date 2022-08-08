<script setup>
// Pinia
import { useExpenseStore } from '@/stores/expense';

const expenseStore = useExpenseStore()

const props = defineProps({
  expense: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <div class="expense-card">
    <div class="expense-card__expense-details">
      <p>{{ expense.date.format('D MMM YYYY')}}</p>
      <p>{{ expense.amount}}</p>
      <p>{{ expenseStore.getCategory(expense).name }}</p>
    </div>
  </div>
</template>

<style scoped>
.expense-card {
  border-radius: 5px;
  border: 1px solid black;
  width: 200px;
  margin: 5px;
  background: white;
  display: flex;
  flex-flow: row nowrap;
  justify-content: space-between;
  cursor: pointer;
}

.expense-card__expense-details {
  margin: 5px;
}

.expense-card__expense-details p {
  margin: 5px;
}

.expense-card:hover {
  filter: drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25))
}

</style>
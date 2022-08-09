<script setup>
// Pinia
import { useExpenseStore } from '@/stores/expense';
const expenseStore = useExpenseStore();

const props = defineProps({
  displayedObject: {
    type: Object,
    required: true,
  },
  mode: {
    type: String,
    required: true,
    validator: value => ['expense', 'revenue'].includes(value),
  }
});
</script>

<template>
  <div class="card">
    <div class="details">
      <p>{{ displayedObject.date.format('D MMM YYYY')}}</p>
      <p>{{ displayedObject.amount }}</p>
      <p v-if="mode == 'expense'">{{ expenseStore.getCategory(displayedObject).name }}</p>
    </div>
  </div>
</template>

<style scoped>
.card {
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

.card:hover {
  filter: drop-shadow(0 4px 4px rgba(0, 0, 0, 0.25))
}

.details {
  margin: 5px;
}

.details p {
  margin: 5px;
}
</style>
<script setup>
// Pinia
const props = defineProps({
  revenue: {
    type: Object,
    required: true,
  }
});
</script>

<template>
  <div class="revenue-card">
    <div class="revenue-details">
      <p>{{ revenue.date.format('D MMM YYYY')}}</p>
      <p>{{ revenue.amount }}</p>
    </div>
  </div>
</template>

<style scoped>
.revenue-card {
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

.revenue-details {
  margin: 5px;
}

.revenue-details p {
  margin: 5px;
}
</style>
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
    <div class="delete-revenue" @click="$emit('delete-revenue')">
      <img alt="delete revenue" src="@/assets/bin_icon.svg"/> 
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

.delete-revenue {
  margin: 10px 10px 0 0;
}

.delete-revenue img {
  width: 18px;
}
</style>
<script setup>
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

import { computed, ref, watch } from 'vue';


const date = ref(dayjs().utc().day(1).hour(0).minute(0).second(0).millisecond(0));

const currentDate = computed(() => date.value.format('MMMM YYYY'));

const addMonth = function() {
  date.value = date.value.add(1, 'month');
}
const subtractMonth = function() {
  date.value = date.value.subtract(1, 'month');
}
defineExpose({date})
</script>

<template>
  <div class="month_picker">
    <div class="arrow left_arrow" @click="subtractMonth">&lt;</div>
    <div class="current_date">{{ currentDate }}</div>
    <div class="arrow right_arrow" @click="addMonth">&gt;</div>
  </div>
</template>

<style scoped>
.month_picker {
  display: flex;
  flex-flow: row nowrap;
  margin:auto;
}

.current_date {
  min-width: 150px;
  text-align: center;
}

.arrow {
  cursor: pointer;
}
</style>
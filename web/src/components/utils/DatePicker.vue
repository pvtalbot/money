<script setup>
// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Vue
import { ref, onMounted } from 'vue';

dayjs.extend(utc);

const props = defineProps({
  mode: {
    type: String,
    default: "MONTH",
    validator: value => ['MONTH', 'YEAR'].includes(value),
  },
  initialDate: {
    required: true,
  },
})

const date = ref(dayjs());
onMounted(() => {date.value = props.initialDate;})
defineExpose({date})

// Add or subtract the duration of the datepicker to current date
const add = function() {
  date.value = date.value.add(1, config.duration);
}
const subtract = function() {
  date.value = date.value.subtract(1, config.duration);
}

const config = (function() {
  return {
    duration: props.mode.toLowerCase(),
    formatDate: props.mode === 'MONTH' ? 'MMMM YYYY' : 'YYYY'
  }
})();
</script>

<template>
  <div class="date_picker">
    <div class="arrow left_arrow" @click="subtract">
      <img src="@/assets/chevron-compact-left.svg" alt="left"/>
    </div>
    <div class="current_date">{{ date.format(config.formatDate) }}</div>
    <div class="arrow right_arrow" @click="add">
      <img src="@/assets/chevron-compact-right.svg" alt="right"/>
    </div>
  </div>
</template>

<style scoped>
.date_picker {
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
  display: flex;
}

.arrow img {
  margin: auto;
}
</style>
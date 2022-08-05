<script setup>
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

import { computed, ref, onMounted } from 'vue';

dayjs.extend(utc);

const props = defineProps({
  mode: {
    type: String,
    default: "MONTH",
    validator: (value) => ['MONTH', 'YEAR'].includes(value),
  },
  initialDate: {
    required: true,
  },
})

const config = (function() {
  const res = {duration: props.mode.toLowerCase()}

  if (props.mode === 'MONTH') {
    res.formatDate = 'MMMM YYYY';
  } else {
    res.formatDate = 'YYYY';
  }

  return res
})();

const date = ref(dayjs());
onMounted(() => {date.value = props.initialDate;})

const currentDate = computed(() => date.value.format(config.formatDate));

const add = function() {
  date.value = date.value.add(1, config.duration);
}
const subtract = function() {
  date.value = date.value.subtract(1, config.duration);
}
defineExpose({date})
</script>

<template>
  <div class="date_picker">
    <div class="arrow left_arrow" @click="subtract">&lt;</div>
    <div class="current_date">{{ currentDate }}</div>
    <div class="arrow right_arrow" @click="add">&gt;</div>
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
}
</style>
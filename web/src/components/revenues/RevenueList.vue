<script setup>
// Apollo
import { useMutation, useQuery } from '@vue/apollo-composable';
import Revenues from '@/graphql/queries/RevenueList.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc'

// Pinia
import { useRevenueStore } from '@/stores/revenue.js'
import { useDrawerStore } from '@/stores/drawer.js'

// Vue
import { computed, ref, watch } from 'vue';

dayjs.extend(utc);
const revenueStore = useRevenueStore();
const drawerStore = useDrawerStore();

const props = defineProps({
  initialDate: {
    required: true,
  },
})

const startDate = computed(() => props.initialDate);
const endDate = computed(() => startDate.value.add(1, 'month'));

const displayedRevenues = computed(() => revenueStore.getCurrentRevenues(startDate.value))
const sortedRevenues = computed(() => {
  return Object.values(displayedRevenues.value).sort((a, b) => {
    if (a.date.isBefore(b.date)) return 1;
    else if (a.date.isAfter(b.date)) return -1;

    return 0;
  })
});

// Apollo Query to get revenues list, hook, watcher
const { result: revenues, onResult: onRevenueListSucceeded, refetch: refetchRevenues, loading: revenuesLoading } =
  useQuery(Revenues,
    { input: { startDate: startDate.value.toISOString(), endDate: endDate.value.toISOString()}}
  );

onRevenueListSucceeded(() => { revenueStore.cacheRevenues(revenues.value.revenues); })

watch(startDate, () => {
  refetchRevenues({
    input: {
      startDate: startDate.value.toISOString(),
      endDate: endDate.value.toISOString(),
    }
  })
})
</script>

<template>
  <div class="revenues-list">
    <div v-if="revenuesLoading.value" class="loader" key="waiting">
      <h2>A minute please, I'm gathering everything!</h2>
    </div>
    <div v-else class="list" key="loaded">
      <p v-if="sortedRevenues.length > 0">Revenues of the month:</p>
      <div v-for="revenue in sortedRevenues" :key="revenue.id" class="revenue">
        {{revenue.id}} - {{revenue.amount}}
      </div>
    </div>
  </div>
</template>

<style scoped>
.revenues-list {
  display: flex;
  flex-flow: column nowrap;
}

.list {
  display: flex;
  flex-flow: column nowrap;
  max-height: 80vh;
  overflow: auto;
}

.list .expense, .list p {
  margin: auto;
}

.component-leave-active, .component-enter-active {
  transition: all 0.5s ease;
}

.component-leave-to, .component-enter-from {
  transform: translate(320px);
}
</style>
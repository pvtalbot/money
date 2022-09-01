<script setup>
// Apollo
import { useLazyQuery } from '@vue/apollo-composable';
import Revenues from '@/graphql/queries/RevenueList.gql';
import Expenses from '@/graphql/queries/ExpenseList.gql';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc'

// Pinia
import { useRevenueStore } from '@/stores/revenue.js'
import { useExpenseStore } from '@/stores/expense.js'
import { useDrawerStore } from '@/stores/drawer.js'

// Vue
import { computed, watch, ref } from 'vue';
import TransferCard from '@/components/transfers/TransferCard.vue';
import UpdateTransferForm from '@/components/transfers/UpdateTransferForm.vue';

dayjs.extend(utc);
const drawerStore = useDrawerStore();

const props = defineProps({
  initialDate: {
    required: true,
  },
  mode: {
    required: true,
    type: String,
    validator: value => ['expense', 'revenue'].includes(value),
  }
})

const config = (() => {
  if (props.mode == 'expense') {
    const store = useExpenseStore();
    return {
      query: Expenses,
      cacheTransfers: store.cacheExpenses,
      resultName: 'expenses',
      displayedTransfers: computed(() => store.getCurrentExpenses(startDate.value)),
      title: 'Expenses of the month',
    };
  } else {
    const store = useRevenueStore();
    return {
      query: Revenues,
      cacheTransfers: store.cacheRevenues,
      resultName: 'revenues',
      displayedTransfers: computed(() => store.getCurrentRevenues(startDate.value)),
      title: 'Revenues of the month',
    };
  }
})();

const startDate = computed(() => props.initialDate);
const endDate = computed(() => startDate.value.add(1, 'month'));

const transferToUpdate = ref(null)
const COMPONENT_TO_DRAWER = 'UpdateTransferForm' + '__' + props.mode;

const sortedTransfers = computed(() => {
  return Object.values(config.displayedTransfers.value).sort((a, b) => {
    if (a.date.isBefore(b.date)) return 1;
    else if (a.date.isAfter(b.date)) return -1;

    return 0;
  })
});

// Apollo Query to get transfers list, hook, watcher
const { result: transfers, load: loadTransfers, onResult: onTransferListSucceeded, loading: transfersLoading } =
  useLazyQuery(config.query,
    { input: { startDate: startDate.value.toISOString(), endDate: endDate.value.toISOString()}}
  );

onTransferListSucceeded(() => { config.cacheTransfers(transfers.value[config.resultName]); })

watch(startDate, () => {
  loadTransfers(undefined, {
    input: {
      startDate: startDate.value.toISOString(),
      endDate: endDate.value.toISOString(),
    }
  })
})

const updateTransfer = transfer => {
  drawerStore.registerComponent(COMPONENT_TO_DRAWER);
  transferToUpdate.value = transfer;
}
</script>

<template>
  <div class="transfers-list">
    <div v-if="transfersLoading.value" class="loader" key="waiting">
      <h2>A minute please, I'm gathering everything!</h2>
    </div>
    <div v-else class="list" key="loaded">
      <p v-if="sortedTransfers.length > 0">{{ config.title }}</p>
      <div v-for="t in sortedTransfers" :key="t.id" class="transfer">
        <TransferCard :displayedTransfer="t"
                      :mode="props.mode"
                      @click="updateTransfer(t)"/>
      </div>
    </div>
    <Teleport to="#teleport-component-to-drawer">
      <Transition name="component">
        <UpdateTransferForm v-if="drawerStore.isCurrentComponentDisplayed(COMPONENT_TO_DRAWER)"
        :updatedTransfer="transferToUpdate"
        :mode="props.mode"/>
      </Transition>
    </Teleport>
  </div>
</template>

<style scoped>
.list {
  display: flex;
  flex-flow: column nowrap;
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

@media (min-width:550px) {
  .list {
    max-height: 80vh;
    overflow: auto;
    /* IE and Edge */
    -ms-overflow-style: none;
    /* Firefox */
    scrollbar-width: none;
  }

  /* Chrome */
  .list::-webkit-scrollbar {
    display: none;
  }

}
</style>
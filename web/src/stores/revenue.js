import dayjs from 'dayjs';

import { defineStore, acceptHMRUpdate } from 'pinia';
import { reactive, computed } from 'vue';

export const useRevenueStore = defineStore('revenues', () => {
  const revenues = reactive({});
  
  const cacheRevenues = newRevenues => {
    for (const revenue of newRevenues) {
      const date = dayjs(revenue.date);
      const year = date.year();
      const month = date.month();
      if (!(year in revenues)) {
        revenues[year] = {};
      }
      if (!(month in revenues[year])) {
        revenues[year][month] = {};
      }
      if (revenue.id in revenues[year][month]) {
        continue;
      }

      revenues[year][month][revenue.id] = {
        id: revenue.id,
        amount: revenue.amount,
        date: date,
      };
    }
  }

  const deleteRevenue = revenue => {
    delete revenues[revenue.date.year()][revenue.date.month()][revenue.id]
  }

  const getCurrentRevenues = computed(() => {
    return date => {
      if (!revenues[date.year()] || !revenues[date.year()][date.month()]) return {};

      return revenues[date.year()][date.month()];
    }
  });

  return {
    revenues,
    getCurrentRevenues,
    cacheRevenues,
    deleteRevenue,
  };
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useRevenueStore, import.meta.hot));
}

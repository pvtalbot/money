import dayjs from 'dayjs';

import { defineStore, acceptHMRUpdate } from 'pinia';
import { reactive, computed } from 'vue';

export const useExpenseStore = defineStore('expenses', () => {
  const expenses = reactive({});
  const expensesCategories = reactive({})

  const cacheExpensesCategories = categories => {
    for (const c of categories) {
      expensesCategories[c.id] = {
        id: c.id,
        name: c.name
      }
    }
  }

  const cacheExpenses = newExpenses => {
    for (const expense of newExpenses) {
      const date = dayjs(expense.date);
      const year = date.year();
      const month = date.month();
      if (!(year in expenses)) {
        expenses[year] = {};
      }
      if (!(month in expenses[year])) {
        expenses[year][month] = {};
      }
      if (expense.id in expenses[year][month]) {
        continue;
      }

      expenses[year][month][expense.id] =  {
        id: expense.id,
        amount: expense.amount,
        date: date,
        categoryId: expense.category.id,
      };
    }
  }

  const deleteExpense = expense => {
    delete expenses[expense.date.year()][expense.date.month()][expense.id];
  }

  const getCurrentExpenses = computed(() => {
    return (date) => {
      if (!expenses[date.year()] || !expenses[date.year()][date.month()]) return {}

      return expenses[date.year()][date.month()];
    }
  });

  const getCategory = computed(() => (expense) => expensesCategories[expense.categoryId]);

  return {
    expenses,
    expensesCategories,
    getCurrentExpenses,
    getCategory,
    cacheExpenses,
    cacheExpensesCategories,
    deleteExpense,
  };

})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useExpenseStore, import.meta.hot));
}
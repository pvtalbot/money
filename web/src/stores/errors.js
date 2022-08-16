import { defineStore, acceptHMRUpdate } from 'pinia';
import { reactive } from 'vue';

export const useErrorStore = defineStore('errors', () => {
  const errors = reactive({});

  const cacheErrors = newErrors => {
    for (const e of newErrors) {
      errors[e.code] = e.translation;
    }
  }

  const getError = code => errors[code] ?? "Unknown error";

  return {
    errors,
    cacheErrors,
    getError,
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useErrorStore, import.meta.hot));
}
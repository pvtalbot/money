import { defineStore, acceptHMRUpdate } from 'pinia';
import { ref } from 'vue';

export const useDrawerStore = defineStore('drawer', () => {
  const drawerOpen = ref(false);

  const setDrawerOpen = (newDrawerOpen) => {
    drawerOpen.value = newDrawerOpen;
  }

  return {
    drawerOpen,
    setDrawerOpen,
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useDrawerStore, import.meta.hot));
}
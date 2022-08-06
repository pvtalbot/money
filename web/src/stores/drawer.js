import { defineStore, acceptHMRUpdate } from 'pinia';
import { computed, ref } from 'vue';

export const useDrawerStore = defineStore('drawer', () => {
  const isDrawerOpen = ref(false);
  const currentComponent = ref("")

  const openDrawer = () => {
    isDrawerOpen.value = true;
  }

  const registerComponent = (componentName) => {
    currentComponent.value = componentName;
    openDrawer();
  }

  const closeDrawer = () => {
    currentComponent.value = "";
    isDrawerOpen.value = false;
  }

  const isCurrentComponentDisplayed = computed(() => (componentName) => componentName === currentComponent.value);

  return {
    isDrawerOpen,
    registerComponent,
    closeDrawer,
    isCurrentComponentDisplayed,
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useDrawerStore, import.meta.hot));
}
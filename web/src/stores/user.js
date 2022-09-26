import { defineStore, acceptHMRUpdate } from 'pinia';
import { ref, computed } from 'vue';

export const useUserStore = defineStore('users', () => {
  const user = ref({
    firstName: "",
    lastName: "",
    id: "",
  });

  const userLoggedIn = computed(() => !!user.value.firstName)

  return {
    user,
    userLoggedIn,
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot));
}
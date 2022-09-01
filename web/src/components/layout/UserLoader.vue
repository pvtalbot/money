<script setup>
// Money
import { 
  useCheckLocalStorageForToken,
  useLoadCurrentUser, 
  useGetAllErrors,
 } from '@/components/login/LoadCurrentUser';

// Vue
import { useRouter } from 'vue-router';

const router = useRouter();
const loadCurrentUser = useLoadCurrentUser();
const loadAllErrors = useGetAllErrors();

useCheckLocalStorageForToken()
 .then(() => { 
  loadCurrentUser();
  loadAllErrors();
 })
 .catch(() => {console.log('oups');router.push({name: 'login'})});
</script>

<template>
  <div class="user-loader-grid-wrapper">
    <div class="user-loader-grid-cell">
      Loading ...
    </div>
  </div>
</template>

<style scoped>
.user-loader-grid-wrapper {
  display: grid;
  grid-template-columns: 1fr;
  height: 25%;
}

.user-loader-grid-cell {
  margin: auto;
}
</style>
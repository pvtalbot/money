<script setup>
// Pinia
import { useUserStore } from '@/stores/user.js';

// Vue
import VueHeader from '@/components/layout/VueHeader.vue';
import Drawer from '@/components/layout/Drawer.vue';
import UserLoader from '@/components/layout/UserLoader.vue';

const userStore = useUserStore();

// Drawer self-teleports to body
</script>

<template>
  <div class="layout">
    <VueHeader/>
    <main>
      <UserLoader v-if="!userStore.userLoggedIn"/>
      <RouterView v-else/>
    </main>
    <Drawer />
  </div>
</template>

<style scoped>
body {
  margin: 0;
}

main {
  position: fixed;
  top: 0;
  padding-left: 111px;
  height: 100%;
  width: 100%;
  max-width: 100vw;
}
</style>
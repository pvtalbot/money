<script setup>
import VueHeader from '@/components/layout/VueHeader.vue';
import Drawer from '../components/layout/Drawer.vue';

// Drawer self-teleports to body
</script>

<template>
  <div class="layout">
    <VueHeader/>
    <main>
      <RouterView/>
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
}
</style>
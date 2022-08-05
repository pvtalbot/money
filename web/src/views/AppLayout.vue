<script setup>
import VueHeader from '@/components/layout/VueHeader.vue';
import HomeView from './HomeView.vue';

</script>
<template>
  <div class="layout">
    <VueHeader/>
    <main>
      <HomeView/>
    </main>
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
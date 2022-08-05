<script setup>
// Pinia
import { useDrawerStore } from '@/stores/drawer.js'

// Vue
import { watch, ref } from 'vue';

const drawerStore = useDrawerStore()

const drawerOpen = ref(false)

const closeDrawer = () => {
  drawerStore.setDrawerOpen(false);
}

watch(() => drawerStore.drawerOpen, (newValue, _) => {
  drawerOpen.value = newValue;
});
</script>

<template>
  <div class="transition_container">
    <transition name="mask_transition">
      <div class="mask" v-if="drawerOpen" @click="closeDrawer">
      </div>
    </transition>
    <transition name="drawer">
      <div class="container" v-if="drawerOpen">
        <div class="icon_close" @click="closeDrawer">
          <img id="icon_close__icon" alt="close" src="@/assets/cross_icon.svg" />
        </div>
        <slot name="component"></slot>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.container {
  background: white;
  height: 100vh;
  width: 300px;
  position: absolute;
  top: -1px;
  right: 0;
  padding: 30px 20px;
  z-index: 2000;
  box-shadow: -6px 0 16px -8px rgba(0,0,0,.08),-9px 0 28px 0 rgba(0,0,0,.05),-12px 0 48px 16px rgba(0,0,0,.03);
}

.icon_close {
  position: absolute;
  top: 10px;
  right: 15px;
  cursor: pointer;
}

.drawer-leave-active, .drawer-enter-active {
  transition: all 0.7s ease;
}

.drawer-leave-to, .drawer-enter {
  transform: translateX(320px);
}

.mask-transition-leave-active, .mask-transition-enter-active {
  transition: all 0.7s ease;
}

.mask-transition-leave-to, .mask-transition-enter {
  opacity: 0;
}
</style>
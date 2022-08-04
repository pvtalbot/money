<script setup>
</script>

<template>
  <header>
    <div class="logo-container logo-container--head">
      <img src="@/assets/euro-money-bag.svg" alt="logo"/>
    </div>

    <div class="logo-container">
      <img src="@/assets/hamburger-menu.svg" alt="expense list"/>
    </div>

    <div class="logo-container">
      <img src="@/assets/chart.svg" alt="charts"/>
    </div>

    <div class="divider"></div>

  </header>
</template>

<style scoped>
header {
  position: fixed;
  width: 71px;
  height: 100vh;
  border-right: 1px solid var(--color-border);
  z-index: 1000;
  display: flex;
  flex-flow: column nowrap;
  top: 0;
}

.logo-container {
  display: flex;
}

.logo-container--head {
  border-bottom: 1px solid var(--color-border);
  margin-bottom: 10px;
}

.logo-container img {
  width: 40px;
  margin: 10px auto;
}

.divider {
  flex: 1;
}
</style>
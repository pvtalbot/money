<script setup>
import LoginForm from "@/components/LoginForm.vue"
</script>

<template>
  <div class="main-wrapper">
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125"/>

    <div class="wrapper">
      <div class="greetings">
        <h1 class="green">Welcome to expenses manager!</h1>
      </div>
    </div>
  </header>
  <main>
    <LoginForm />
  </main>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

.greetings h1 {
  text-align: center;
}

main {
    height: 100%;
    width: 100%;
    display: flex;
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  .main-wrapper {
    display: grid;
    grid-template-columns: 1fr 1fr;
    padding: 0 4rem;
  }

  header {
    display: flex;
    place-items: center;
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  .greetings h1 {
    text-align: left;
  }
}
</style>
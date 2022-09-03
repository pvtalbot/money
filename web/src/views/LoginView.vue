<script setup>
// Money
import { useCheckLocalStorageForToken, useGetAllErrors } from "@/components/login/LoginUtils";

// Vue
import { ref, computed } from 'vue';
import LoginForm from "@/components/login/LoginForm.vue"
import SignupForm from '@/components/login/SignupForm.vue';
import { useRouter } from 'vue-router';

const router = useRouter();

useGetAllErrors()();
useCheckLocalStorageForToken()
  .then(() => {router.push({name: 'home'})})
  .catch(() => {}); // user stays on current page as there is no token

const signup = ref(false)

const switchForm = () => {signup.value = !signup.value};
const message = computed(() => signup.value ? 'Go to login' : 'Go to signup')
</script>

<template>
  <div class="main-wrapper">
    <header>
      <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

      <div class="wrapper">
        <div class="greetings">
          <h1 class="green">Welcome to expenses manager!</h1>
        </div>
      </div>
    </header>
    <main>
      <div class="switch" @click="switchForm">
        <p>{{ message }}</p>
        <div class="arrow right_arrow">
          <img src="@/assets/chevron-compact-right.svg" alt="right" />
        </div>
      </div>
      <div class="forms">
        <Transition mode="out-in" name="login-form">
          <SignupForm v-if="signup" />
          <LoginForm v-else />
        </Transition>
      </div>
    </main>
  </div>
</template>

<style scoped>
.login-form-leave-active, .login-form-enter-active {
  transition: all 0.3s ease;
}

.login-form-leave-to, .login-form-enter-from {
  transform: translate(100px);
  opacity: 0.5;
}

h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

.main-wrapper {
  height: 100vh;
}

.greetings h1 {
  text-align: center;
}

main {
    width: 100%;
    display: flex;
    flex-flow: column;
    justify-content: center;
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

.switch {
  margin-right: auto;
  margin-left: auto;
  display: flex;
  cursor: pointer;
}

.switch p, .switch .arrow {
  margin: auto;
  display: flex;
}

.arrow img {
  margin: auto;
}

.forms {
  margin-right: auto;
  margin-left: auto;
}

@media (min-width: 1024px) {
  .main-wrapper {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 200px 1fr 200px;
    grid-template-areas: ". ." "left right" ". .";
    padding: 0 4rem;
  }

  header {
    display: flex;
    place-items: center;
    grid-area: left;
  }

  main {
    grid-area: right;
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
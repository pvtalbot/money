<script setup>
// Apollo
import { useLazyQuery } from "@vue/apollo-composable";
import ValidateAccessToken from '@/graphql/queries/ValidateAccessToken.gql';

// Gasp
import { gsap } from "gsap";

// Money
import { useLoadCurrentUser, getUserToken, removeUserToken, useGetAllErrors } from "@/components/login/LoadCurrentUser";

// Vue
import { ref, onMounted, computed } from 'vue';
import LoginForm from "@/components/login/LoginForm.vue"
import SignupForm from '@/components/login/SignupForm.vue';

const loadCurrentUser = useLoadCurrentUser();
const getAllErrors = useGetAllErrors();

// Apollo Query to check the validity of a token. Used on mounted
const {result: tokenValidity, load: loadTokenValidity, onResult: onTokenValidated} = useLazyQuery(ValidateAccessToken);
// Checks if there is a token in local storage. If yes, checks validity. If the token is still valid, logs the user in
onMounted(() => {
  const accessToken = getUserToken();
  if (accessToken == null) return;

  loadTokenValidity(undefined, {accessToken: accessToken});
})
onTokenValidated(() => {
  if (!tokenValidity.value.validateAccessToken) {
    removeUserToken();
    return;
  }

  loadCurrentUser();
})


const signup = ref(false)

const switchForm = () => {signup.value = !signup.value};
const message = computed(() => signup.value ? 'Go to login' : 'Go to signup')

const onEnter = (el, done) => {
  gsap.from(el, {
    transform: 'translateX(-100px)',
    opacity: 0,
    onComplete: done,
    duration: 0.3,
  })
}

const onLeave = (el, done) => {
  gsap.to(el, {
    transform: 'translateX(100px)',
    opacity: 0,
    onComplete: done,
    duration: 0.3,
  })
}
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
      <div class="forms">
        <Transition mode="out-in" @enter="onEnter" @leave="onLeave" :css="false">
          <SignupForm v-if="signup" />
          <LoginForm v-else />
        </Transition>
      </div>
      <div class="switch" @click="switchForm">
        <p>{{ message }}</p>
        <div class="arrow right_arrow">
          <img src="@/assets/chevron-compact-right.svg" alt="right" />
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped>
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
  margin: auto;
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
  margin: auto;
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
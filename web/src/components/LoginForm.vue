<script setup>
import { ref, watch } from 'vue';
import { useUserStore } from '../stores/user';
import { useRouter } from 'vue-router';
import VueButton from './VueButton.vue';

import { useMutation, useLazyQuery } from '@vue/apollo-composable'
import LoginMutation from '../graphql/mutations/LoginMutation.gql'
import Me from '@/graphql/queries/CurrentUser.gql';

const userStore = useUserStore();
const router = useRouter();

const username = ref("");
const password = ref("");
const disabled = ref(false);


const {mutate: loginMutation} = useMutation(LoginMutation);
const {result: currentUser, load: loadCurrentUser, onResult: onCurrentUserSucceeded} = useLazyQuery(Me);

const login = () => {
  disabled.value = false;
  let response = loginMutation({login: {username: username.value, password: password.value}});
  response.then(({data: {login}}) => {
    if (typeof localStorage !== 'undefined' && login) {
      localStorage.setItem('accessToken', login)
    }
  })
  .then(() => loadCurrentUser())

  onCurrentUserSucceeded(() => {
    if (currentUser.value) {
      userStore.$patch((store) => {
        store.user.firstName = currentUser.value.me.firstName;
        store.user.lastName = currentUser.value.me.lastName;
      })
    }
  })
}

watch(() => userStore.userLoggedIn, () => {router.push({name: 'app'})});
</script>

<template>
  <div class="login-form">
    <form @submit.prevent="login">
      <div class="login-form__item-container">
        <label for="login-form__login">Login</label>
        <input id="login-form__login" type="text" v-model="username">
      </div>
      <div class="login-form__item-container">
        <label for="login-form__password">Password</label>
        <input id="login-form__password" type="password" v-model="password">
      </div>
      <VueButton button-type="submit" message="Submit" :disabled="disabled"/>
    </form>
  </div>
</template>

<style scoped>
.login-form {
    margin: auto;
}

.login-form__item-container {
    padding: 0 0 10px 0;
    display: flex;
    flex-flow: column nowrap;
}

form {
    display: flex;
    flex-flow: column nowrap;
}

label {
    align-self: start;
}
</style>
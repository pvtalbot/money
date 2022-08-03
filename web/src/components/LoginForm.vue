<script setup>
import { ref } from 'vue';
import VueButton from './VueButton.vue';

import { useMutation } from '@vue/apollo-composable'
import LoginMutation from '../graphql/mutations/LoginMutation.gql'

const username = ref("");
const password = ref("");
const disabled = ref(false);

const {mutate: login} = useMutation(LoginMutation);

</script>

<template>
  <div class="login-form">
    <form @submit.prevent="login({login: {username: username, password: password}})">
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
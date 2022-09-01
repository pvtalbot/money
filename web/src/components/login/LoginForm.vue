<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable'
import LoginMutation from '@/graphql/mutations/LoginMutation.gql'

// Money
import { storeUserToken } from './LoginUtils';

// Vue
import { ref } from 'vue';
import VueButton from '@/components/utils/VueButton.vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const username = ref("");
const password = ref("");
const disabled = ref(false);

// Apollo Mutation to log the user in
const {mutate: loginMutation, onDone: onLoginSuccess } = useMutation(LoginMutation);

// Wrapper function for login mutation
const login = () => {
  disabled.value = true;
  loginMutation({login: {username: username.value, password: password.value}})
  .finally(() => {disabled.value = false;});
}
onLoginSuccess(({data: {login}}) => {
  storeUserToken(login);
  router.push({name: 'home'});
})
</script>

<template>
  <div class="login-form">
    <form @submit.prevent="login">
      <div class="login-form__item-container">
        <label for="login-form__login">Login</label>
        <input id="login-form__login" type="text" v-model="username" required>
      </div>
      <div class="login-form__item-container">
        <label for="login-form__password">Password</label>
        <input id="login-form__password" type="password" v-model="password" required>
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
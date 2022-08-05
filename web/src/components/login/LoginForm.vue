<script setup>
// Apollo
import { useMutation, useLazyQuery } from '@vue/apollo-composable'
import LoginMutation from '@/graphql/mutations/LoginMutation.gql'
import ValidateAccessToken from '@/graphql/queries/ValidateAccessToken.gql'
import Me from '@/graphql/queries/CurrentUser.gql';

// Vue
import { ref, watch, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';
import { useRouter } from 'vue-router';
import VueButton from '@/components/utils/VueButton.vue';

const userStore = useUserStore();
const router = useRouter();

const username = ref("");
const password = ref("");
const disabled = ref(false);


// Apollo Mutation to log the user in
const {mutate: loginMutation} = useMutation(LoginMutation);
// Apollo Query to get info on the current user (defered after log in)
const {result: currentUser, load: loadCurrentUser, onResult: onCurrentUserSucceeded} = useLazyQuery(Me);
// Apollo Query to check the validity of a token. Used when the component is mounted
const {result: tokenValidity, load: loadTokenValidity, onResult: onTokenValidated} = useLazyQuery(ValidateAccessToken)

// Wrapper function for login mutation
const login = () => {
  disabled.value = false;
  let response = loginMutation({login: {username: username.value, password: password.value}});
  response.then(({data: {login}}) => {
    if (typeof localStorage !== 'undefined' && login) {
      localStorage.setItem('accessToken', login)
    }
  })
  .then(() => loadCurrentUser())

}

onCurrentUserSucceeded(() => {
  if (!currentUser.value) return;
  userStore.$patch((store) => {
    store.user.firstName = currentUser.value.me.firstName;
    store.user.lastName = currentUser.value.me.lastName;
  })
})

// If loadCurrentUser is called and succeed, will then call onCurrentUserSucceeded
onTokenValidated(() => {
  if (!tokenValidity.value) return;
  if (!tokenValidity.value.validateAccessToken) {localStorage.removeItem('accessToken'); return;}
  loadCurrentUser();
})

watch(() => userStore.userLoggedIn, () => {router.push({name: 'home'})});

// Check if there is a token in local storage. If yes, checks validity. If the token is still valid, logs the user in.
onMounted(() => {
  const accessToken = localStorage.getItem('accessToken');
  if (accessToken == null) return;
  loadTokenValidity(undefined, {accessToken: accessToken});
})
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
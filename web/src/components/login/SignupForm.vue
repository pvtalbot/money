<script setup>
// Apollo
import { useMutation } from '@vue/apollo-composable';
import CreateUserMutation from '@/graphql/mutations/CreateUserMutation.gql';

// Money
import { useLoadCurrentUser, storeUserToken } from './LoadCurrentUser';

// Vue
import { reactive, ref } from 'vue';
import VueButton from '@/components/utils/VueButton.vue';

const username = ref("");
const password = ref("");
const firstName = ref("");
const lastName = ref("");
const disabled = ref(false);

const errors = reactive([]);

const mapErrors = {
  'createuser-1': 'An account with this username already exists',
}
const getError = code => mapErrors[code] ?? 'Unknown error';

const loadCurrentUser = useLoadCurrentUser();

// Apollo Mutation to create the user
const {mutate: createUserMutation, onError: onCreateUserError, onDone: onCreateUserSuccess} = useMutation(CreateUserMutation);
// Wrapper function for the mutation
const createNewUser = () => {
  disabled.value = true;
  errors.length = 0;
  createUserMutation({input: {
    firstName: firstName.value,
    lastName: lastName.value,
    password: password.value,
    name: username.value,
  }}).finally(() => {disabled.value = false;})
}

onCreateUserError(e => {
  for (const k of e.graphQLErrors) {
    if (k.extensions && k.extensions.code) errors.push(k.extensions.code)
  }
});

onCreateUserSuccess(({data: {createUser}}) => {
  storeUserToken(createUser);
  loadCurrentUser();
});

</script>

<template>
  <div class="signup-form">
    <form @submit.prevent="createNewUser">
      <div class="item-container">
        <label for="signup-form__username">Username</label>
        <input id="signup-form__username" type="text" v-model="username" required />
      </div>
      <div class="item-container">
        <label for="signup-form__password">Password</label>
        <input id="signup-form__password" type="text" v-model="password" required />
      </div>
      <div class="item-container">
        <label for="signup-form__firstName">First name</label>
        <input id="signup-form__firstName" type="text" v-model="firstName" required />
      </div>
      <div class="item-container">
        <label for="signup-form__lastName">Last name</label>
        <input id="signup-form__lastName" type="text" v-model="lastName" required />
      </div>
      <VueButton button-type="submit" message="Submit" :disabled="disabled" />
    </form>
    <div class="item-container errors" v-if="errors.length > 0">
      <p v-for="e in errors">{{ getError(e) }}</p>
    </div>
  </div>
</template>

<style scoped>
.signup-form {
  margin: auto;
  width: 200px;
}

.item-container {
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

.errors {
  color: var(--em-c-pink-1);
}

</style>
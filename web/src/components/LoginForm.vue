<script setup lang="ts">

</script>

<template>
  <div class="login-form">
    <form>
      <div class="login-form__item-container">
        <label for="login-form__login">Login</label>
        <input id="login-form__login" type="text">
      </div>
      <div class="login-form__item-container">
        <label for="login-form__password">Password</label>
        <input id="login-form__password" type="password">
      </div>
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
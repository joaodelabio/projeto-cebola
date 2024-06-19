<template>
  <div class="login-card">
    <h2>Login</h2>
    <form class="login-form" @submit.prevent="handleLogin">
      <input type="text" placeholder="Username or email" v-model="usernameOrEmail" />
      <input type="password" placeholder="Password" v-model="password" />
      <RouterLink to="/register" class="signup-link">Not a member? Sign up now</RouterLink>
      <button type="submit">Login</button>
    </form>
    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const usernameOrEmail = ref('');
const password = ref('');
const errorMessage = ref('');

const handleLogin = () => {
  if (usernameOrEmail.value === '' || password.value === '') {
    errorMessage.value = 'cade os dados, bixo?';
  } else {

    // puxando dados da localStorage

    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const userData = JSON.parse(storedUser);
      if (
        (usernameOrEmail.value === userData.username || usernameOrEmail.value === userData.email) &&
        password.value === userData.password
      ) {
        router.push('/library');
      } else {
        errorMessage.value = 'não te conheço nao bixo.';
      }
    }
  }
};
</script>

<style scoped>
* {
  box-sizing: border-box;
}

@keyframes rotate {
  100% {
    background-position: 15% 50%;
  }
}

:root {
  --color-primary: #7656f0;
}

body {
  display: grid;
  place-items: center;
  margin: 0;
  height: 100vh;
  padding: 0 24px;
  background-color: var(--color-primary);
  background-image: url("../assets/images/login-bg.svg");
  background-repeat: no-repeat;
  background-size: cover;
  color: #f9f9f9;
  animation: rotate 6s infinite alternate linear;
}

body::after {
  content: "";
  position: fixed;
  inset: 0;
}

.login-card {
  position: relative;
  z-index: 3;
  width: 100%;
  margin: 0 20px;
  padding: 70px 30px 44px;
  border-radius: 20px;
  background: rgba(24, 21, 36, 0.66);
  backdrop-filter: blur(10px);
  text-align: center;
}

.login-card > h2 {
  font-size: 36px;
  font-weight: 400;
  margin: 0 0 12px;
}

.login-card > h3 {
  color: rgba(255, 255, 255, 0.38);
  margin: 0 0 48px;
  font-weight: 400;
  font-size: 16px;
}

.login-form {
  width: 100%;
  display: grid;
  gap: 16px;
  margin: 0;
}

.login-form > input,
.login-form > button {
  width: 100%;
  height: 56px;
  border-radius: 6px;
  border: 0;
  font-family: inherit;
  font-size: 16px;
  padding: 0 16px;
}

.login-form > input {
  color: rgba(255, 255, 255, 0.96);
  background: rgba(255, 255, 255, 0.08);
}

.login-form > input::placeholder {
  color: rgba(255, 255, 255, 0.38);
}

.login-form > button {
  cursor: pointer;
  background: var(--color-primary);
  color: #f9f9f9;
  font-size: 18px;
  font-weight: 400;
  text-align: center;
  transition: all 0.375s;
}

.signup-link {
  color: var(--color-primary);
  font-size: 16px;
  text-align: left;
  text-decoration: none;
  margin-bottom: 30px;
}

.error-message {
  color: red;
  margin-top: 20px;
}

@media (min-width: 448px) {
  .login-card {
    margin: 0;
    width: 70%;
    min-width: 400px;
  }
}

@media (min-width: 628px) {
  .login-card {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    padding: 0 44px;
    width: 50%;
    max-width: 500px;
    min-width: auto;
    display: flex;
    border-radius: 0;
    flex-direction: column;
    justify-content: center;
  }
}
</style>
<template>
  <div class="register-card">
    <h2>Register</h2>
    <form class="register-form" @submit.prevent="handleRegister">
      <input type="text" placeholder="Username" v-model="username" />
      <input type="email" placeholder="Email" v-model="email" />
      <input type="password" placeholder="Password" v-model="password" />
      <button type="submit">Register</button>
    </form>
    <RouterLink to="/login" class="login-link">Already have an account? Log in</RouterLink>
    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const username = ref('');
const email = ref('');
const password = ref('');
const errorMessage = ref('');

const handleRegister = () => {
  if (username.value === '' || email.value === '' || password.value === '') {
    errorMessage.value = 'Cadê os dados irmão?';
  } else {

    // salvar localmente por enquanto

    const userData = {
      username: username.value,
      email: email.value,
      password: password.value,
    };
    localStorage.setItem('user', JSON.stringify(userData));
    router.push('/login');
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

.register-card {
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

.register-card > h2 {
  font-size: 36px;
  font-weight: 400;
  margin: 0 0 12px;
}

.register-card > h3 {
  color: rgba(255, 255, 255, 0.38);
  margin: 0 0 48px;
  font-weight: 400;
  font-size: 16px;
}

.register-form {
  width: 100%;
  display: grid;
  gap: 16px;
  margin: 0;
}

.register-form > input,
.register-form > button {
  width: 100%;
  height: 56px;
  border-radius: 6px;
  border: 0;
  font-family: inherit;
  font-size: 16px;
  padding: 0 16px;
}

.register-form > input {
  color: rgba(255, 255, 255, 0.96);
  background: rgba(255, 255, 255, 0.08);
}

.register-form > input::placeholder {
  color: rgba(255, 255, 255, 0.38);
}

.register-form > button {
  cursor: pointer;
  background: var(--color-primary);
  color: #f9f9f9;
  font-size: 18px;
  font-weight: 400;
  text-align: center;
  transition: all 0.375s;
}

.login-link {
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
  .register-card {
    margin: 0;
    width: 70%;
    min-width: 400px;
  }
}

@media (min-width: 628px) {
  .register-card {
    position: fixed;
    top: 0;
    left: 0; 
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

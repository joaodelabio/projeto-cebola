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
import axios from 'axios';

const router = useRouter();
const usernameOrEmail = ref('');
const password = ref('');
const errorMessage = ref('');

const handleLogin = async () => {
  if (usernameOrEmail.value === '' || password.value === '') {
    errorMessage.value = 'Preencha todos os campos';
  } else {
    try {
      const response = await axios.post('http://localhost:8080/api/login', {
        email: usernameOrEmail.value,
        password: password.value,
      });
      if (response.data.success) {
        router.push('/library');
      } else {
        errorMessage.value = 'Dados inválidos.';
      }
    } catch (error) {
      errorMessage.value = 'Erro ao tentar login.';
    }
  }
};
</script>

<style scoped>
.login-card {
  width: 100%;
  max-width: 400px;
  margin: auto;
  padding: 2rem;
  border-radius: 0.5rem;
  background: rgba(0, 0, 0, 0.7);
  text-align: center;
  color: #fff;
}

h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
}

.login-form {
  display: grid;
  gap: 1rem;
}

.login-form input {
  width: 100%;
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
}

.login-form button {
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
  background-color: #FBB7C7;
  color: #fff;
  cursor: pointer;
}

.signup-link {
  display: block;
  margin-top: 1rem;
  color: #FBB7C7;
}

.error-message {
  color: red;
  margin-top: 1rem;
}
</style>

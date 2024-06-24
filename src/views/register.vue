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
import axios from 'axios';

const router = useRouter();
const username = ref('');
const email = ref('');
const password = ref('');
const errorMessage = ref('');

const handleRegister = async () => {
  if (username.value === '' || email.value === '' || password.value === '') {
    errorMessage.value = 'Preencha todos os campos.';
  } else {
    try {
      const response = await axios.post('http://localhost:8080/api/register', {
        username: username.value,
        email: email.value,
        password: password.value,
      });
      if (response.data.success) {
        router.push('/login');
      } else {
        errorMessage.value = 'Erro ao tentar registrar.';
      }
    } catch (error) {
      errorMessage.value = 'Erro ao tentar registrar.';
    }
  }
};
</script>

<style scoped>
.register-card {
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

.register-form {
  display: grid;
  gap: 1rem;
}

.register-form input {
  width: 100%;
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
}

.register-form button {
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
  background-color: #FBB7C7;
  color: #fff;
  cursor: pointer;
}

.login-link {
  display: block;
  margin-top: 1rem;
  color: #FBB7C7;
}

.error-message {
  color: red;
  margin-top: 1rem;
}
</style>

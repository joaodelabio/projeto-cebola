<!--pagina inicial-->

<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
  <div class="home-container">
    <header class="header">
      <h1>Your personal<br>library</h1>
      <div class="buttons">
        <RouterLink to="/login">
          <button class="login-button">LOGIN</button>
        </RouterLink>
        <RouterLink to="/register">
          <button class="register-button">REGISTER</button>
        </RouterLink>
      </div>
    </header>
    <aside class="image-container">
        <img src="@/assets/images/bigbook.png" alt="Book Image" class="book-image"/>
    </aside>
  </div>
</template>

<style scoped>
@font-face {
  font-family: 'League Spartan';
  src: url('@/assets/fonts/LeagueSpartan-Regular.ttf') format('truetype');
}

.home-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100vh;
  color: white;
  padding: 2rem;
  box-sizing: border-box;
}

.header {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}

h1 {
  font-family: 'League Spartan', sans-serif;
  font-size: 4rem;
  margin: 0;
  line-height: 1.2;
}

.buttons {
  display: flex;
  gap: 1rem;
  margin-top: 2rem;
}

.login-button {
  background-color: #FBB7C7;
  color: white;
  padding: 1rem 2rem;
  border: none;
  border-radius: 0.5rem;
  font-size: 1rem;
  cursor: pointer;
}

.register-button {
  background-color: transparent;
  color: white;
  padding: 1rem 2rem;
  border: 2px solid white;
  border-radius: 0.5rem;
  font-size: 1rem;
  cursor: pointer;
}

.image-container {
  display: flex;
  justify-content: flex-end;
  align-items: flex-end;
  width: 50%;
  position: relative;
  bottom: -5rem; 
  right: -10rem; 
}

.book-image {
  width: 200%; 
  height: auto;
  position: relative;
}
</style>
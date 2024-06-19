<template>
  <div class="library-background">
    <header>
      <h2>nao sei</h2>
    </header>
    <main>
      <section class="books-section">
        <img src="@/assets/images/bookguy.png" alt="Book Image" class="book-image">
      </section>
    </main>
    <footer></footer>
  </div>
</template>

<script setup>
</script>

<style scoped>


header {
  display: flex;
  justify-content: center;
  padding: 20px;
  text-align: center;
  color: #fff;
}

header h2 {
  font-size: 36px;
  margin: 0;
}

main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 60px);
  padding: 350px;
  width: 100%;
}

.books-section {
  position: relative;
  background-color: #FBB7C7;
  border-radius: 10px;
  padding: 200px;
  width: 100%;
  max-width: 1200px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.book-image {
  position: absolute;
  left: -500px; /* Ajuste conforme necessário para posicionar a imagem fora do retângulo */
  top: 50%;
  transform: translateY(-50%);
  width: 400px;  /* Ajuste a largura da imagem conforme necessário */
  height: auto;
}

header {
  position: absolute;
  top: 50px; /* Ajuste conforme necessário para posicionar o header */
}

header h2 {
  margin-bottom: 10px; /* Ajuste conforme necessário para a distância entre o texto e a imagem */
}
</style>

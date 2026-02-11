<script>
import Hero from '../components/Hero.vue'
import Trademark from '../components/Trademark.vue'

export default {
  name: 'IndexView',
  components: {
    Hero,
    Trademark
  }
}
</script>

<template>
  <Hero />
  <Trademark />
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

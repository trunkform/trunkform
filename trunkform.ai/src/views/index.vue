<script>
import TrunkformHero from "../components/TrunkformHero.vue";
import TrunkformTrademark from "../components/TrunkformTrademark.vue";

export default {
  name: "IndexView",
  components: {
    TrunkformHero,
    TrunkformTrademark,
  },
};
</script>

<template>
  <main>
    <nav>
      <a href="/">trunkform</a>
      <!-- <a href="#">log in</a> -->
    </nav>

    <TrunkformHero />

    <TrunkformTrademark />
  </main>
</template>


<style scoped>
main {
  min-height: 100vh;
  padding: 1rem 2rem;
  overflow: hidden;
}

nav {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-shrink: 0;
}

nav a:first-child {
  display: inline-flex;
  align-items: center;
  padding: 0;
  margin: 0;
  background: transparent;
  border: none;
  border-radius: 0;
  box-shadow: none;
  font-size: 1.875rem;
  font-weight: 700;
  line-height: 1.5rem;
}

nav a:first-child::before {
  content: "";
  width: 22px;
  height: 22px;
  margin-right: 0.7rem;
  margin-top: 0.15rem;
  box-sizing: border-box;
  border: 5px solid #00bfff;
  border-radius: 9999px;
  background: transparent;
  box-shadow:
    0 0 0 1px rgba(0, 191, 255, 0.15),
    0 0 18px rgba(0, 191, 255, 0.35);
}

nav a:last-child:not(:first-child) {
  padding: 1rem 1.5rem;
  border: 1px solid rgba(255, 255, 255, 0.18);
  border-radius: 14px;
  background: rgba(10, 14, 22, 0.4);
  color: rgba(255, 255, 255, 0.9);
  font-size: 1rem;
}

:deep(section) {
  position: relative;
  z-index: 1;
  max-width: 900px;
  margin-left: auto;
  margin-right: auto;
}

:deep(h1) {
  margin: 0 auto 1rem;
  color: var(--text);
  font-size: clamp(1.5rem, 7vw, 5.5rem);
  line-height: 1;
  font-weight: 700;
}

:deep(h2) {
  color: var(--text);
  font-size: clamp(1.25rem, 7vw, 4rem);
  line-height: 1;
  font-weight: 700;
}

:deep(h3) {
  color: var(--text);
  font-size: 1rem;
  line-height: 1;
  font-weight: 700;
}

:deep(p) {
  max-width: 900px;
  margin: 0 auto;
  line-height: 1.4;
}

main :deep(p:last-child) {
  padding-bottom: 0;
}

@media (max-width: 768px) {
  main {
    padding: 1.5rem;
  }

  nav a:first-child {
    font-size: 1.25rem;
    padding-left: 0;
  }
}

@media (orientation: landscape) and (max-width: 1024px) {
  :deep(h1),
  :deep(p),
  :deep(div) {
    padding-bottom: 0.4rem;
  }

  :deep(h1) {
    font-size: clamp(1.5rem, 6vw, 2.5rem);
    width: 100%;
  }

  :deep(p) {
    font-size: clamp(.75rem, 4vw, 1.25rem);
  }

  :deep(div a) {
    font-size: 0.875rem;
    padding: 0.5rem 1rem;
  }

  nav a:first-child {
    font-size: 1.25rem;
    padding-left: 0;
  }
}
</style>

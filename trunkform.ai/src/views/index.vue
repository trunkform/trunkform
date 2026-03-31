<script>
import Hero from "../components/Hero.vue";
import Trademark from "../components/Trademark.vue";

export default {
  name: "IndexView",
  components: {
    Hero,
    Trademark,
  },
};
</script>

<template>
  <main>
    <img src="/hero.png" alt="" />

    <nav>
      <a href="/">trunkform</a>
      <!-- <a href="#">log in</a> -->
    </nav>

    <section id="hero">
      <h1>Designed for Speed, Built for Trust</h1>
      <p>
        <b>trunkform</b>&trade; <b>MCP</b> helps developers remove process bottlenecks,
        standardize delivery, and build release ready software.
      </p>

      <div>
        <a href="https://github.com/trunkform/trunkform/blob/trunk/mcp/README.md">learn more</a>
        <a href="https://github.com/trunkform/trunkform">view the source</a>
      </div>
    </section>

    <section id="trademark">
      <Trademark />
    </section>
  </main>
</template>


<style scoped>
main {
  min-height: 100vh;
  padding: 1rem 2rem;
  overflow: hidden;
}

#hero {
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 100dvh;
  padding-bottom: env(safe-area-inset-bottom, 1rem);
}

#trademark {
  text-align: center;
  position: relative;
  z-index: 1;
}

#trademark :deep(p) {
  font-size: 1rem;
  color: var(--muted);
  text-shadow:
    0px 0px 30px #000000ed,
    0px 0px 60px #000000ed;
}

#trademark :deep(a),
#trademark :deep(a:visited) {
  color: var(--muted);
}

img {
  position: fixed;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.82;
  pointer-events: none;
}

h1,
h2,
#hero p,
#hero > div {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: 1rem;
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

h1 {
  max-width: 900px;
  margin: 0 auto 1rem;
  color: var(--text);
  font-size: clamp(1.5rem, 7vw, 5.5rem);
  line-height: 1;
  font-weight: 700;
}

h2 {
  max-width: 900px;
  margin: 0 auto;
  color: var(--text);
  font-size: clamp(1.25rem, 7vw, 4rem);
  line-height: 1;
  font-weight: 700;
}

h3 {
  max-width: 900px;
  margin: 0 auto;
  color: var(--text);
  font-size: 1rem;
  line-height: 1;
  font-weight: 700;
}

p {
  max-width: 900px;
  margin: 0 auto;
  font-size: clamp(1.3rem, 3.5vw, 2rem);
  padding-bottom: clamp(0.5rem, 7vw, 1.5rem);
  line-height: 1.4;
}

p:last-child {
  padding-bottom: 0;
}

#hero p {
  max-width: 900px;
  margin: 0 auto;
  font-size: clamp(1.125rem, 2.5vw + 0.5rem, 2.5rem);
  line-height: 1.4;
  text-shadow:
    0px 0px 30px #000000ed,
    0px 0px 60px #00000045,
    0px 0px 90px #000000ed;
  color: white;
}

#hero p:last-child {
  padding-bottom: 0;
}

#hero > div {
  margin-top: 1.5rem;
  text-align: center;
}

#hero > div a {
  display: inline-block;
  margin-right: 1rem;
  padding: 0.85rem 1.5rem;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  text-align: center;
}

#hero > div a:first-child {
  background: linear-gradient(180deg, #1ec8ff 0%, #009ee6 100%);
  color: #fff;
  box-shadow: 0 0 24px rgba(0, 191, 255, 0.3);
}

#hero > div a:last-child {
  border: 1px solid rgba(255, 255, 255, 0.24);
  background: rgba(4, 16, 24, 0.85);
  color: #fff;
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
  h1,
  #hero p,
  #hero > div {
    padding-bottom: 0.4rem;
  }

  h1 {
    font-size: clamp(1.5rem, 6vw, 2.5rem);
    width: 100%;
  }

  #hero p {
    font-size: clamp(1rem, 4vw, 1.25rem);
    width: 100%;
  }

  #hero > div a {
    font-size: 0.875rem;
    padding: 0.5rem 1rem;
  }

  #trademark :deep(p) {
    font-size: 0.75rem;
  }

  nav a:first-child {
    font-size: 1.25rem;
    padding-left: 0;
  }
}
</style>

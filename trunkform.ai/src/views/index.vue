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

    <section>
      <nav>
        <a href="/">trunkform</a>
        <!-- <a href="#">log in</a> -->
      </nav>

      <div class="content">
        <h1>Designed for Speed, Built for Trust</h1>
        <p>
          <b>trunkform</b>&trade; <b>MCP</b> helps developers remove process bottlenecks,
          standardize delivery, and build release ready software.
        </p>

        <div>
          <a href="https://github.com/trunkform/trunkform/blob/trunk/mcp/README.md">learn more</a>
          <a href="https://github.com/trunkform/trunkform">view the source</a>
        </div>
      </div>
    </section>

    <section>
      <Trademark style="text-align: center; font-size: 1.3rem !important" />
    </section>
  </main>
</template>

<style scoped>
main {
  min-height: 100vh;
  padding: clamp(0.5rem, 3vw, 2rem) clamp(0.75rem, 5vw, 4rem);
  overflow: hidden;
}

section:first-of-type {
  display: flex;
  flex-direction: column;
  min-height: 100dvh;
  padding-bottom: env(safe-area-inset-bottom, 1rem);
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  min-height: 80vh;
  max-height: 80vh;
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

p {
  /* a big dark black text drop shadow */
  text-shadow:
    0px 0px 30px #000000ed,
    0px 0px 60px #00000045,
    0px 0px 90px #000000ed;
  color: white;
}

h1,
h2,
p,
.content > div {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: clamp(0.5rem, 2vw, 2rem);
}

nav {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: clamp(0.5rem, 2vw, 2rem);
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0rem;
}

nav a:first-child {
  display: inline-flex;
  align-items: center;
  padding: 0;
  margin: 0;

  background: transparent !important;
  border: none !important;
  border-radius: 0;
  box-shadow: none !important;

  color: #fff !important;
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

nav a:last-child {
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
  font-size: clamp(1.75rem, 7vw, 5.5rem);
  line-height: 1;
  font-weight: 700;
}

h2 {
  max-width: 900px;
  margin: 0 auto;
  color: var(--text);
  font-size: clamp(2rem, 7vw, 4rem);
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

/* all additional p after h2 should style this way*/
#why {
  max-width: 56rem;
  margin: 0 auto 25vh auto;
  padding: 2rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 18px;
  background: rgba(12, 16, 26, 0.55);
  backdrop-filter: blur(18px);
  text-align: left;
}

.content > div {
  margin-top: clamp(0.75rem, 3vw, 3rem);
  text-align: center;
}

div a {
  display: inline-block;
  margin-right: 1rem;
  padding: clamp(0.65rem, 2vw, 1.1rem) clamp(1rem, 3vw, 2rem);
  border-radius: 12px;
  font-size: clamp(0.875rem, 2.5vw, 1.1rem);
  font-weight: 600;
  text-align: center;
}

div a:first-child {
  background: linear-gradient(180deg, #1ec8ff 0%, #009ee6 100%);
  color: #fff;
  -webkit-text-fill-color: #fff;
  box-shadow: 0 0 24px rgba(0, 191, 255, 0.3);
}

div a:last-child {
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
</style>

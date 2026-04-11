<script>
export default {
  name: 'TrunkformHero',
  props: {}
}
</script>

<template>
  <img
    src="/hero.png"
    alt=""
  >

  <section>
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
</template>

<style scoped>
img {
  position: fixed;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.82;
  pointer-events: none;
}

section {
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 100dvh;
  padding-bottom: env(safe-area-inset-bottom, 1rem);
}

h1,
p,
div {
  max-width: 900px;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: 1rem;
}

p {
  max-width: 900px;
  margin: 0 auto;
  font-size: clamp(1.125rem, 2.5vw + 0.5rem, 2.5rem);
  width: 100%;
  line-height: 1.4;
  text-shadow:
    0px 0px 30px #000000ed,
    0px 0px 60px #00000045,
    0px 0px 90px #000000ed;
  color: white;
}

p:last-child {
  padding-bottom: 0;
}

div {
  margin-top: 1.5rem;
  text-align: center;
}

div a {
  display: inline-block;
  margin-right: 1rem;
  padding: 0.85rem 1.5rem;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  text-align: center;
}

div a:first-child {
  background: linear-gradient(180deg, #1ec8ff 0%, #009ee6 100%);
  color: #fff;
  box-shadow: 0 0 24px rgba(0, 191, 255, 0.3);
}

div a:last-child {
  border: 1px solid rgba(255, 255, 255, 0.24);
  background: rgba(4, 16, 24, 0.85);
  color: #fff;
}
</style>

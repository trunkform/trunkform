<script>
export default {
  name: 'TrunkformTrademark',
  props: {},
  data() {
    return {}
  }
}
</script>

<template>
  <section>
    <p>
      <a href="https://tsdr.uspto.gov/#caseNumber=99631818&caseSearchType=US_APPLICATION&caseType=DEFAULT&searchType=statusSearch">
        <span>trunkform is a</span> <span>trademark of</span> <span>Richard James Minchuk</span>
      </a>
    </p>
  </section>
</template>

<style scoped>

p {
  text-align: center;
  color: var(--muted);
  text-shadow:
    0px 0px 30px #000000ed,
    0px 0px 60px #000000ed;
}

a,
a:visited {
  color: var(--muted) !important;
  text-decoration: none;
}

span {
  white-space: nowrap;
}
</style>

<script>
export default {
  name: 'Trademark',
  props: {},
  data() {
    return {}
  }
}
</script>

<template>
  <p>
    <a href="https://tsdr.uspto.gov/#caseNumber=99631818&caseSearchType=US_APPLICATION&caseType=DEFAULT&searchType=statusSearch">
      <span>trunkform is a</span> <span>trademark of</span> <span>Richard James Minchuk</span>
    </a>
  </p>
</template>

<style scoped>
p {
  font-size: 0.75rem;
  margin: 0 0 1rem 0;
  letter-spacing: 0.2rem;
  width: 100%;
  text-align: center;
}
p a {
  color: #888;
  text-decoration: none;
}
span {
  white-space: nowrap;
}
</style>

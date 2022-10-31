<template>
  <Index></Index>
</template>

<script>
import Index from '@/view/Index';
export default {
  name: 'App',
  created: function () {
    this.axios.get('http://localhost:8081/').then((res) => {
      console.log(res);
    });
  },
  components: {
    Index
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

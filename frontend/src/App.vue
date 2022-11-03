<template>
  <div id="app">
    <router-view/>
    <Footer></Footer>
  </div>
</template>

<style>
  #app {
    max-width: 960px;
    margin: 0 auto;
  }
</style>
<script>
import Footer from './components/Footer';
export default {
  components: {Footer}
};
</script>
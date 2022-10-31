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
</style>

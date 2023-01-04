<template>
  <div class="content-container">
    <component :is="tabName"></component>
  </div>
</template>

<script>
import Articles from "./Articles.vue";
import Aside from "./Aside.vue";
import About from "./About.vue";
export default {
  name: 'Content',
  computed: {
    tabName: function () {
      let route = this.$route.path
      switch (route) {
        case '/home':
          return Articles
        case '/about':
          return About
        default:
          return Aside
      }
    }
  }
};
</script>

<style lang="less" scoped>
</style>
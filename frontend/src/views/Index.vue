<template>
  <div class="container">
    <el-container>
      <el-header class="header"><Header></Header></el-header>
      <el-container class="main-container">
        <div style="width: 20%"></div>
        <el-main class="main-body"><Content></Content></el-main>
        <div style="width: 20%"></div>
      </el-container>
      <el-footer class="footer"><Footer></Footer></el-footer>
    </el-container>
  </div>
</template>

<script>
import Header from '../components/Header.vue';
import Footer from '../components/Footer.vue';
import Content from '../components/Content.vue';
import Aside from '../components/Aside.vue';
export default {
  name: 'Index',
  components: {
    Aside,
    Content,
    Footer,
    Header
  }
};
</script>

<style lang="less" scoped>
.container {
  width: 100%;
}
.main-body {
  width: 60%;
}
.header {
  padding: 0;
  width: 100%;
}
.footer {
  margin: 0;
  padding: 0;
  width: 100%;
}
</style>
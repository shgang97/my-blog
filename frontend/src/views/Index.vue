<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header"><Header></Header></el-header>
      <el-container>
        <el-aside><Aside></Aside></el-aside>
        <el-main><Content></Content></el-main>
        <el-aside><Aside></Aside></el-aside>
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
.header {
  padding: 0;
  width: 100%;
}
.footer {
  padding: 0;
  width: 100%;
}
</style>
<template>
  <div class="common-layout">
    <el-container class="container">
      <el-header class="header">
        <template v-slot>
          <Header></Header>
        </template>
      </el-header>
      <el-container>
        <el-aside width="200px" class="aside">
          <Aside></Aside>
        </el-aside>
        <el-container>
          <el-main>
            <Content></Content>
          </el-main>
          <el-footer><Footer></Footer></el-footer>
        </el-container>
      </el-container>
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
}
</style>
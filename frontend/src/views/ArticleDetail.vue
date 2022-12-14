<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">
        <Header></Header>
      </el-header>
      <el-main class="main-container">
        <div class="preview-container">
          <v-md-editor height="800px" :model-value="text" mode="preview"></v-md-editor>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Header from '../components/Header.vue';

export default {
  name: 'ArticleDetail',
  data() {
    return {
      text: '### hello',
    };
  },
  components: {Header}
};
</script>

<style lang="less" scoped>
.container {
  width: 100%;
}

.header {
  padding: 0;
  width: 100%;
}

.preview-container {
  width: 60%;
  margin-left: 20%;
}

</style>
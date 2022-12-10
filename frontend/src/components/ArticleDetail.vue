<template>
  <div class="article-layout">
    <el-container>
      <el-header>
        <Header></Header>
      </el-header>
      <el-main>
        <div>
          <h1>article detail--{{ this.$route.params.id }}</h1>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Header from './Header.vue';

export default {
  name: 'ArticleDetail',
  components: {Header}
};
</script>

<style lang="less" scoped>
.article-layout {
  width: 100%;
}

</style>
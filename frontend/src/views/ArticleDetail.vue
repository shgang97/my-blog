<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">
        <Header></Header>
      </el-header>
      <el-main class="main-container">
        <div class="preview-container">
          <v-md-editor height="800px" :model-value="article.article.content" mode="preview"></v-md-editor>
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
      article: {article: '', tags: [], category: ''},
    };
  },
  components: {Header},
  methods: {
    async getArticle(id) {
      const {data: res} = await this.$http.get('/articles/' + id);
      this.article = res.data
    }
  },
  created() {
    const id = this.$route.params.id;
    this.getArticle(id)
  }
};
</script>

<style lang="less" scoped>

.header {
  padding: 0;
  width: 100%;
}

.preview-container {
  width: 60%;
  margin-left: 20%;
}

</style>
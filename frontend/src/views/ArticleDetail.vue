<template>
  <div class="container">
    <el-container>
      <el-header class="header">
        <Header></Header>
      </el-header>
      <el-container>
        <el-aside class="aside-container">
          <div class="article-suspended-panel">
            <div class="badge">
              <el-badge :value="article.view_count" type="info" class="badge-item center-item">
                <i class="iconfont icon-view1"></i>
              </el-badge>
            </div>
            <div class="badge">
              <el-badge :value="article.like_count" type="info" class="badge-item center-item">
                <i class="iconfont icon-good"></i>
              </el-badge>
            </div>
            <div class="badge">
              <el-badge type="info" class="badge-item center-item">
                <i class="iconfont icon-share"></i>
              </el-badge>
            </div>
          </div>

        </el-aside>
        <el-main class="main-container">
          <div class="article-header">
            <h1 class="title">{{ article.title }}</h1>
            <div class="meta-info">
              <div>本文最后编辑于：{{ dateFormat(article.update_at, "yyyy年MM月dd日 hh:mm") }}
                <router-link class="edit-button" :to="'/writing/' + article.id">编辑</router-link>
              </div>
              <div>
                <router-link class="tag vertical-line"
                             :to="'/tag/' + tag.id"
                             :key="tag.id"
                             v-for="tag in tags">#{{ tag.name }}
                </router-link>
                <router-link class="category vertical-line"
                             :to="'/category/' + category.id"
                             :key="category.id">{{ category.name }}
                </router-link>
              </div>

            </div>
          </div>
          <div class="preview-container">
            <v-md-editor height="800px" :model-value="article.content" mode="preview"></v-md-editor>
          </div>
        </el-main>
      </el-container>

    </el-container>
  </div>
</template>

<script>
import Header from '../components/Header.vue';

export default {
  name: 'ArticleDetail',
  data() {
    return {
      article: {},
      tags: [],
      category: ''
    };
  },
  components: {Header},
  methods: {
    async getArticle(id) {
      const {data: res} = await this.$http.get('/articles/' + id);
      this.article = res.data.article
      this.tags = res.data.tags
      this.category = res.data.category
    }
  },
  created() {
    const id = this.$route.params.id;
    this.getArticle(id)
  },
  computed: {
    dateFormat() {
      return this.$dataFormat
    }
  }
};
</script>

<style lang="less" scoped>
.container {
  width: 100%;
  height: 50px;
}

.header {
  padding: 0;
  width: 100%;
  height: 50px;
}

.article-header {
  /*text-align: center;*/
  background-color: white;
  font-size: 1.5em;
  font-weight: normal;
  margin: 10px 0;
  overflow-wrap: break-word;
  word-wrap: break-word;

}

.main-container {
  margin-right: 20%;
}

.title {
  margin: 5px 0;
}

.meta-info {
  color: #8a919f;
  font-size: 1.2rem;
}
.edit-button {
  text-decoration: none;
}
.vertical-line {
  border-left: 1px #86909c;
  padding: 0 5px;
  color: inherit;
  text-decoration: none;
}

.aside-container {
  padding-top: 40px;
}

.article-suspended-panel {
  position: fixed;
  margin-left: 15%;
  top: 130px;
  z-index: 2;
}

.badge {
  margin-top: 20px;
  width: 44px;
  height: 44px;
  border-radius: 50%;
  background-color: white;
}

.iconfont {
  font-size: 26px;
}

.center-item {
  top: 9px;
  left: 9px;
}

i {
  color: #8a919f;
}
</style>
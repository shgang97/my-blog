<template>
  <div class="article-page-container">
    <div class="article-list">
      <div class="article-container" v-for="article in articles" :key="article.article.id">
        <div class="meta-container">
          <div class="date vertical-line">{{dateFormat(article.article.update_at, "yyyy年MM月dd日 hh:mm")}}</div>
          <div class="tag-list ">
            <div class="tag vertical-line">
              <router-link class="tag vertical-line"
                           :to="'/tag/' + tag.id"
                           :key="tag.id"
                           v-for="tag in article.tags">#{{ tag.name }}
              </router-link>
            </div>
          </div>
          <div class="category vertical-line">
            <router-link :to="'/category/' + article.category.id">{{ article.category.name }}</router-link>
          </div>
          <div class="delete vertical-line">
            <el-button v-if="canDelete" class="delete-button" @click="deleteArticle(article.article.id)">删除</el-button>
          </div>
        </div>
        <div class="content-wrapper">
          <div class="title">
            <router-link :to="'/articles/' + article.article.id">{{ article.article.title }}</router-link>
          </div>
          <div class="abstract">
            {{ article.article.content }}
          </div>
          <ul class="action-list">
            <li class="item view"><i class="iconfont icon-view1"></i><span>{{ article.article.view_count }}</span></li>
            <li class="item comment"><i class="iconfont icon-comment_fill_light"></i><span>{{ article.article.comment_count }}</span></li>
            <li class="item like"><i class="iconfont icon-good"></i><span>{{ article.article.like_count }}</span></li>
          </ul>
        </div>
      </div>
    </div>
    <div class="pagination-container">
      <el-pagination
          small
          layout="prev, pager, next"
          :current-page="currentPage"
          :page-size="pageSize"
          :total="total"
          @current-change="postPage"
          class="mt-4"
      />
    </div>
  </div>
</template>

<script>

export default {
  name: 'Articles',
  data() {
    return {
      articles: [],
      currentPage: 1,
      pageSize: 10,
      total: 0,
      canDelete: false
    };
  },
  methods: {
    async postPage(page) {
      this.currentPage = page;
      const {data: res} = await this.$http.get('/articles', {params: {page: this.currentPage, pageSize: this.pageSize}});
      this.articles = res.data.data;
      this.total = res.data.total;
    },
    async deleteArticle(id) {
      const {data: res} = await this.$http.delete("/articles/" + id, {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      await this.postPage(1);
    }
  },
  async created() {
    await this.postPage(1);
    if (this.$store.getters.getUser.username) {
      this.canDelete = true
    } else {
      this.canDelete = false
    }
  },
  computed: {
    dateFormat() {
      return this.$dataFormat
    },
  }
};
</script>

<style lang="less" scoped>
.meta-container {
  display: flex;
  color: #86909c;

}

a {
  color: inherit;
  text-decoration: none;
}

.vertical-line {
  border-left: 1px #86909c;
  padding: 0 5px;
}
.delete-button {
  border: 0 none;
  outline: none;
  padding: 0 5px;
  color: #86909C;
}

.title {
  font-weight: 700;
  font-size: 16px;
  line-height: 24px;
  color: #1d2129;
  width: 100%;
  display: -webkit-box;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;

}

ul {
  display: flex;
  align-items: center;
}

li {
  list-style: none;
  padding: 0 10px;
}

li > i {
  padding-right: 5px;
}

.article-container {
  padding: 5px;
  width: 100%;
  background-color: white;
  margin-bottom: 5px;
}

.pagination-container {
  margin-top: 5px;
}

.action-list {
  padding: 0;
  color: #8a919f;
}
</style>
<template>
  <div class="article-page-container">
    <div class="article-list">
      <div class="article-container" v-for="article in articles">
        <div class="meta-container">
          <div class="date vertical-line">1个月前</div>
          <div class="tag-list ">
            <div class="tag vertical-line">
              <router-link class="tag vertical-line"
                           :to="'/tag/' + id"
                           :key="id"
                           v-for="(id, index) in article.tag_ids">#{{ article.tag_names[index] }}
              </router-link>
            </div>
          </div>
          <div class="category vertical-line">
            <router-link :to="'/category/' + article.category_id">{{ article.category_name }}</router-link>
          </div>
        </div>
        <div class="content-wrapper">
          <div class="title">
            <router-link :to="'/article/' + article.id">{{ article.title }}</router-link>
          </div>
          <div class="abstract">{{ article.description }}</div>
          <ul class="action-list">
            <li class="item view"><i class="iconfont icon-yuedu"></i><span>{{ article.viewCount }}</span></li>
            <li class="item comment"><i class="iconfont icon-pinglun"></i><span>{{ article.commentCount }}</span></li>
            <li class="item like"><i class="iconfont icon-dianzan"></i><span>{{ article.likeCount }}</span></li>
          </ul>
        </div>
      </div>
    </div>
    <div>
      <el-pagination
          small
          background
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
      total: 0
    };
  },
  methods: {
    async postPage(page) {
      this.currentPage = page;
      const {data: res} = await this.$http.get('/articles', {params: {page: this.currentPage, pageSize: this.pageSize}});
      this.articles = res.data.data;
      this.total = res.data.total;
    }
  },
  async created() {
    await this.postPage(1);
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
}

.action-list {
  padding: 0;
}
</style>
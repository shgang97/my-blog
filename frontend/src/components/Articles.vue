<template>
  <div class="article-page-container">
    <div class="article-list">
      <ArticleCard :key="article.id"
                   :id="article.id"
                   :title="article.title + article.id"
                   v-for="article in articles">

      </ArticleCard>
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
import ArticleCard from './ArticleCard.vue'
import Pagination from './Pagination.vue'

export default {
  name: "Articles",
  data() {
    return {
      articles: [
        {id: '001', title: 'hello'},
        {id: '002', title: 'hello'},
        {id: '003', title: 'hello'},
        {id: '004', title: 'hello'},
        {id: '005', title: 'hello'},
        {id: '006', title: 'hello'},
        {id: '007', title: 'hello'},
        {id: '008', title: 'hello'},
        {id: '009', title: 'hello'},
        {id: '010', title: 'hello'},
      ],
      currentPage: 1,
      pageSize: 10,
      total: 15,
    }
  },
  components: {
    ArticleCard,
    Pagination
  },
  methods: {
    postPage(page) {
      this.currentPage = page
      console.log(page)
      if (page === 2) {
        this.articles = [
          {id: '011', title: 'hello'},
          {id: '012', title: 'hello'},
          {id: '013', title: 'hello'},
          {id: '014', title: 'hello'},
          {id: '015', title: 'hello'},
        ]
      }
      if (page === 1) {
        this.articles = [
          {id: '001', title: 'hello'},
          {id: '002', title: 'hello'},
          {id: '003', title: 'hello'},
          {id: '004', title: 'hello'},
          {id: '005', title: 'hello'},
          {id: '006', title: 'hello'},
          {id: '007', title: 'hello'},
          {id: '008', title: 'hello'},
          {id: '009', title: 'hello'},
          {id: '010', title: 'hello'},
        ]
      }
    }
  },
  created() {
    this.postPage(1)
  }
}
</script>

<style lang="less" scoped>
</style>
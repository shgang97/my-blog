<template>
  <div class="mcontaner">
    <Header></Header>

    <div class="block">
      <el-timeline>

        <el-timeline-item :timestamp="post.createTime" placement="top" v-for="post in posts">
          <el-card>
            <h4>
              <router-link :to="{name: 'PostDetail', params: {postId: post.id}}">
                {{post.title}}
              </router-link>
            </h4>
            <p>{{post.description}}</p>
          </el-card>
        </el-timeline-item>

      </el-timeline>

      <el-pagination class="mpage"
                     background
                     layout="prev, pager, next"
                     :current-page="currentPage"
                     :page-size="pageSize"
                     :total="total"
                     @current-change=postPage>
      </el-pagination>

    </div>

  </div>
</template>

<script>
  import Header from "../components/Header";

  export default {
    name: "Posts",
    components: {Header},
    data() {
      return {
        posts: {},
        currentPage: 1,
        total: 0,
        pageSize: 5
      }
    },
    methods: {
      postPage(currentPage) {
        const _this = this
        _this.$axios.get("/api/posts", {params: {page: 1}}).then(res => {
          let PostRes = res.data.data
          _this.posts = PostRes.posts
          _this.currentPage = PostRes.page
          _this.total = PostRes.total
          _this.pageSize = PostRes.pageSize

        })
      }
    },
    created() {
      this.postPage(1)
    }
  }
</script>

<style scoped>

  .mpage {
    margin: 0 auto;
    text-align: center;
  }

</style>
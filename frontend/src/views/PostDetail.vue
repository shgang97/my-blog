<template>
  <div>
    <Header></Header>

    <div class="mblog">
      <h2> {{ post.title }}</h2>
      <el-link icon="el-icon-edit" v-if="ownPost">
        <router-link :to="{name: 'PostEdit', params: {postId: post.id}}" >
        编辑
        </router-link>
      </el-link>
      <el-divider></el-divider>
      <div class="markdown-body" v-html="post.content"></div>

    </div>

  </div>
</template>

<script>
  import 'github-markdown-css'
  import Header from "../components/Header";

  export default {
    name: "PostDetail",
    components: {Header},
    data() {
      return {
        post: {
          id: "",
          title: "",
          content: ""
        },
        ownPost: false
      }
    },
    created() {
      const postId = this.$route.params.postId
      const _this = this
      this.$axios.get('/api/post/' + postId).then(res => {
        console.log("res.data.data = ", res.data.data)
        const post = res.data.data
        _this.post.id = post.id
        _this.post.title = post.title

        let MardownIt = require("markdown-it")
        let md = new MardownIt()

        let result = md.render(post.content)
        _this.post.content = result
        _this.ownPost = (post.userId === _this.$store.getters.getUser.id)

      })
    }
  }
</script>

<style scoped>
  .mblog {
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    width: 100%;
    min-height: 700px;
    padding: 20px 15px;
  }

</style>
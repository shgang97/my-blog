<template>
  <div class="edit-container">
    <el-form :model="article" :rules="rules" ref="subForm">
      <div class="header">
        <el-form-item prop="title" class="input-item">
          <input v-model="article.title" placeholder="输入文章标题..." spellcheck="false" maxlength="80"
                 class="title-input">
        </el-form-item>
        <el-form-item class="button button-item">
          <el-button @click="save(article)" plain>存草稿</el-button>
          <div class="publishPopup">
            <el-button type="primary" @click="display = !display">发布</el-button>
            <el-card class="panel" v-if="display">
              <div>发布文章</div>
              <div class="form-item">
                <div class="label required category-label"> 分类：</div>
                <el-row class="mb-4">
                  <el-button v-model="article.category">后端</el-button>
                  <el-button v-model="article.category">前端</el-button>
                  <el-button v-model="article.category">Android</el-button>
                  <el-button v-model="article.category">iOS</el-button>
                </el-row>
                <el-row class="mb-4">
                  <el-button v-model="article.category">后端</el-button>
                  <el-button v-model="article.category">前端</el-button>
                  <el-button v-model="article.category">Android</el-button>
                  <el-button v-model="article.category">iOS</el-button>
                </el-row>
                <el-form-item prop="tag">
                  <div class="label required category-label"> 标签：</div>
                  <el-input v-model="article.tags" placeholder="请输入标签..." spellcheck="false"></el-input>
                </el-form-item>
                <el-row class="mb-4 button-action">
                  <el-button @click="display = !display">取消</el-button>
                  <el-button type="primary" @click="submit(article)">发布文章</el-button>
                </el-row>
              </div>

            </el-card>
          </div>

        </el-form-item>
      </div>
      <div class="main">
        <el-form-item class="v-md-editor-container" prop="content">
          <v-md-editor v-model="article.content" height="800px"></v-md-editor>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'Writing',
  data() {
    return {
      article: {
        title: 'title',
        content: 'content:hello world',
        tags: ['tag'],
        category: 'category'
      },
      display: false,
      rules: {
        title: [
          {required: true, message: '请输入标题', trigger: 'blur'},
          {min: 3, max: 25, message: '长度在 3 到 25 个字符', trigger: 'blur'}
        ],
        content: [
          {required: true, message: '请输入内容', trigger: 'blur'}
        ]
      }
    };
  },
  methods: {
    submit(article) {
      this.$refs['subForm'].validate((valid) => {

        if (valid) {
          console.log(article)
          if ('/writing' === this.$route.path) {
            this.write(this.article)
          } else if ('/writing' === this.$route.path) {

          }
        } else {
          return false;
        }
      });
    },
    async write(article) {
      const {data: res} = await this.$http.post('/article', article, {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      console.log(res)
      if (res.code === 200) {
        await this.$router.push('/article/' + res.data)
      }
    },
    save() {
      console.log(this.article.title);
    }
  }
};
</script>

<style lang="less" scoped>
.header {
  display: flex;
  height: 60px;
  background: white;
}

.title-input {
  margin: 0;
  padding: 0;
  font-size: 24px;
  font-weight: 500;
  color: #1d2129;
  border: none;
  outline: none;
  width: 100%;
}

.input-item {
  width: 90%;
  height: 40px;
}

.title-input {
  height: 60px;
}

.button-item {
  width: 10%;
  height: 100%;
  padding-left: 80px;
  margin: 0;
}

.panel {
  top: 65px;
  position: fixed;
  z-index: 999;
  right: 10px;
}
</style>
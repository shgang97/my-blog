<template>
  <div class="edit-container">
    <el-form :model="article" :rules="rules" ref="subForm">
      <div class="header">
        <el-form-item prop="title" class="input-item">
          <input v-model="article.article.title" placeholder="输入文章标题..." spellcheck="false" maxlength="80"
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
                <el-radio-group v-model="category" size="large" @change="change">
                  <el-radio-button :key="category.id" :label="category.name" v-for="category in categories" />
                </el-radio-group>
                <div class="label required category-label"> 标签：</div>
                <el-select
                    v-model="article.tags"
                    value-key="id"
                    multiple
                    filterable
                    allow-create
                    default-first-option
                    :reserve-keyword="false"
                    placeholder="Choose tags for your article"
                >
                  <el-option
                      v-for="tag in tags"
                      :key="tag.id"
                      :label="tag.name"
                      :value="tag"
                  />
                </el-select>
                <el-row class="mb-4 button-action">
                  <el-button @click="display = !display">取消</el-button>
                  <el-button type="primary" @click="write(article)" v-if="update">发布文章</el-button>
                  <el-button type="primary" @click="modify(article)" v-else>更新文章</el-button>
                </el-row>
              </div>

            </el-card>
          </div>

        </el-form-item>
      </div>
      <div class="main">
        <el-form-item class="v-md-editor-container" prop="content">
          <v-md-editor v-model="article.article.content" height="800px"></v-md-editor>
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
      article: {article: {}, category: {}, tags: []},
      tags: {},
      categories: [],
      display: false,
      update: false,
      category: '',
      categoryMap: {},
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
    show() {
      console.log(this.article.category.id)
      console.log(this.article.category.name)
    },
    change(value) {
      console.log(value)
      this.article.category = {id: this.categoryMap[value], name: value}
    },
    write(article) {
      this.$refs['subForm'].validate( async (valid) => {
        if (valid) {
          const {data: res} = await this.$http.post('/articles', article, {
            headers: {
              'Authorization': localStorage.getItem('token')
            }
          })
          if (res.code === 200) {
            await this.$router.push('/articles/' + res.data)
          }
        }
      })
    },
    modify(article) {
      console.log(article)
    },
    save() {
      console.log(this.article.title);
    },
    async getTags() {
      const {data: res} = await this.$http.get('/tags')
      if (res.code === 200) {
        this.tags = res.data
      }
    },
    async getCategories() {
      const {data: res} = await this.$http.get('/categories')
      if (res.code === 200) {
        this.categories = res.data
      }
    },
    async getArticle(id) {
      const {data: res} = await this.$http.get('/articles/' + id)
      if (res.code === 200) {
        console.log(res)
        this.article = res.data
      }
    },
  },
  async created() {
    if (this.$route.params.id) {
      await this.getArticle(this.$route.params.id)
    }
    await this.getTags()
    await this.getCategories()
    this.categories.forEach(category => {this.categoryMap[category.name] = category.id})
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
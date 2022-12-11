<template>
  <div class="container">
    <a href="/" class="logo">
      <img src="../assets/images/logo.jpeg" alt="logo-img" class="logo-img">
    </a>

    <div class="nav-container nav-container-left">
      <nav>
        <router-link to="/home">
          <el-icon>
            <HomeFilled/>
          </el-icon>
          首页
        </router-link>
        <router-link to="/category">
          <el-icon>
            <Grid/>
          </el-icon>
          分类
        </router-link>
        <router-link to="/tag">
          <el-icon>
            <CollectionTag/>
          </el-icon>
          标签
        </router-link>
        <router-link to="/archive">
          <el-icon>
            <Calendar/>
          </el-icon>
          归档
        </router-link>
        <router-link to="/about">
          <el-icon>
            <UserFilled/>
          </el-icon>
          关于
        </router-link>
        <router-link to="/link">
          <el-icon>
            <Link/>
          </el-icon>
          友链
        </router-link>
      </nav>
    </div>
    <div class="nav-container nav-container-right">
      <input class="text" type="text" placeholder="请输入关键词~~~" name="search">
      <input class="button" type="button" value="搜索">
      <div class="has-login-wrapper" v-if="hasLogin">
        <img :src="this.userInfo.avatar">
        <nav>
          <el-button type="primary" plain @click="logout">退出</el-button>
        </nav>
        <nav class="write-article">
          <router-link to="/writing" class="write-article-link">
            <el-button type="primary" round><i class="iconfont icon-Writing"></i>发布文章</el-button>
          </router-link>
        </nav>
      </div>
      <div class="login-button-wrapper" v-else>
        <nav class="write-article">
          <router-link to="/login" class="write-article-link">
            <el-button type="primary" plain>登录</el-button>
          </router-link>
        </nav>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Header',
  data() {
    return {
      userInfo: {
        username: '',
        avatar: ''
      },
      hasLogin: false
    };
  },
  methods: {
    async logout() {
      const {data: res} = await this.$http.get('/logout', {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      this.hasLogin = false
      this.$store.commit('REMOVE_INFO')
      await this.$router.push('/login')
    }
  },
  created() {
    if (this.$store.getters.getUser.username) {
      this.userInfo.username = this.$store.getters.getUser.username
      this.userInfo.avatar = this.$store.getters.getUser.avatar
      this.hasLogin = true
    }
  }
};
</script>

<style lang="less" scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  background-color: #292c2f;
  color: #ffffff;
  /*吸顶效果*/
  position: absolute;
  position: -webkit-sticky; /*!*兼容 -webkit 内核的浏览器*!*/
  top: 0; /*!*必须设一个值，否则不生效*!*/

}

.nav-container {
  display: flex;
  align-items: center;
  width: 50%;
}

a.logo {
  display: flex;
  align-items: center;
  font: normal 28px Cookie, Arial, Helvetica, sans-serif;
  padding: 0px 10px;
  width: 10%;
}

img {
  height: 40px;
}

nav {
  display: flex;
  align-items: center;
  margin: 0px 40px;
  font: 16px Arial, Helvetica, sans-serif;
}

.has-login-wrapper, .login-button-wrapper {
  display: flex;
}

.write-article-link {
  width: 120px;
}

/*/deep/ router-link {*/
/*  padding: 0 10px;*/
/*}*/

nav a {
  padding: 0 5px;
  width: 78px;
  text-decoration: none;
  color: #ffffff;
  font-size: 16px;
  font-weight: normal;
  opacity: 0.9;
}

nav a:hover {
  opacity: 1;
}

.router-link-active {
  color: #608bd2;
  pointer-events: none;
  opacity: 1;
}

/*搜索框*/

.text {
  height: 22px;
  font-size: 18px;
  border: 1px solid #ccc;
  padding: 3px 16px;
  border-bottom-left-radius: 20px;
  border-top-left-radius: 20px;
}

.text:focus {
  outline: none;
  border-color: rgba(82, 168, 236, 0.8);
  box-shadow: inset 0 2px 2px rgba(0, 0, 0, 0.075), 0 0 8px rgba(82, 168, 236, 0.6);
}

.button {
  width: 60px;
  height: 30px;
  font-size: 14px;
  margin-right: 35px;
  border: 1px solid #608bd2;
  background-color: #608bd2;
  border-top-right-radius: 20px;
  border-bottom-right-radius: 20px;
}
</style>
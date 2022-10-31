<template>
  <div class="header-wrap">
    <!-- logo -->
    <div class="header-left">
      <h1>
        <a href="/">
          <img class="logo-img" src="/images/logo.png" alt="-LOGO" />
        </a>
      </h1>
      <!-- 导航 -->
      <ul class="header-nav">
        <li><a href=""></a></li>
      </ul>
    </div>
    <div class="header-right">
      <div class="header-search">
        <input id="search-input" type="text" placeholder="搜索一下" />
        <i class="iconfont icon-search"></i>
        <div class="search-drop"></div>
      </div>
      <a class="login-action" href="/login">
        <button class="login-btn">登录</button>
      </a>
      <div class="login-end">
        <p class="iconfont icon-yonghu"><span class="login-username"></span></p>
        <ul class="login-drop">
          <a href="/writing"><li>写作</li></a>
          <li class="login-out">退出登录</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Header'
};
</script>

<style scoped>
/* HEADER */
header {
  position: fixed;
  width: 100%;
  height: 60px;
  background: #fff;
  backdrop-filter: saturate(180%) blur(0.4rem);
  background-color: rgba(255, 255, 255, 0.55);
  border-bottom: 1px solid #f1f1f1;
  box-shadow: 0px 4px 10px rgb(0 0 0 / 5%);
  z-index: 99;
  top: 0;
}
.logo-img {
  /*height: 100px;*/
  max-height: 60px;
}
.header-wrap {
  width: 1080px;
  height: 100%;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
}
.header-wrap ul {
  list-style: none;
}
.header-left,
.header-right {
  height: 100%;
  display: flex;
  align-items: center;
}
.header-wrap h1 {
  font-size: 30px;
  font-weight: bold;
}
.header-wrap a {
  color: #333;
}
.header-nav {
  display: flex;
  margin-left: 20px;
}
.header-nav li {
  margin-left: 20px;
}
.header-nav .active {
  color: #448ef6;
}
.header-search {
  position: relative;
}
.header-search input {
  border: 1px solid #f1f1f1;
  background-color: transparent;
  outline: none;
  padding: 8px 10px;
  padding-right: 30px;
  border-radius: 6px;
  width: 230px;
  color: #86909c;
}
.header-search i {
  position: absolute;
  right: 10px;
  top: 8px;
  color: #86909c;
}
.search-drop {
  display: none;
  position: absolute;
  top: 40px;
  left: 0;
  background-color: #fff;
  padding: 10px 0;
  box-shadow: 0 0 10px #f1f1f1;
  z-index: 9;
  border-radius: 6px;
  min-width: 270px;
  white-space: nowrap;
}
.search-drop a {
  display: block;
  padding: 0 12px;
  line-height: 28px;
  color: #1d2129;
  font-size: 14px;
  cursor: pointer;
  text-decoration: none;
}
.search-drop a:hover {
  background-color: #f1f1f1;
}
</style>
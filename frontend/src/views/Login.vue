<template>
  <div>
    <el-card class="login-form-layout">
      <el-form
          autocomplete="on"
          :model="loginForm"
          ref="loginForm"
          label-position="left"
          :rules="rules"
      >
        <div style="text-align: center">
          <svg-icon icon-class="login-mall" style="width: 56px;height: 56px;color: #409EFF"></svg-icon>
        </div>
        <h2 class="login-title color-main">账号密码登录</h2>
        <el-form-item prop="username" class="input-row">

          <el-input
              name="username"
              type="text"
              v-model="loginForm.username"
              autocomplete="on"
              placeholder="请输入用户名"
          >
            <template #prefix><i class="iconfont icon-user"/></template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password" class="input-row">
          <el-input
              name="password"
              :type="pwdType"
              @keyup.enter.native="login"
              v-model="loginForm.password"
              autocomplete="on"
              placeholder="请输入密码"
          >
            <template #prefix><i class="iconfont icon-mima"/></template>
<!--            <template #suffix v-if="show" @click="showPwd"><i class="iconfont icon-icon_login_input_password_default"/></template>-->
<!--            <template #suffix v-else @click="showPwd"><i class="iconfont icon-icon_login_input_password_pressed"/></template>-->
          </el-input>
        </el-form-item>
        <el-form-item style="margin-bottom: 60px">
          <el-button
              style="width: 100%"
              type="primary"
              :loading="loading"
              @click.native.prevent="login"
          >登录
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>

export default {
  name: 'login',
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 15, message: '长度在 3 到 15 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请选择密码', trigger: 'change' }
        ]
      },
      loading: false,
      pwdType: 'password',
    };
  },
  methods: {
    login() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.handleLogin()
        } else {
          return false;
        }
      })
    },
    async handleLogin() {
      const {data: res} = await this.$http.post('/login', this.loginForm)
      if (res.code !== 200) return alert('登录失败！')
      let data = res.data
      const token = data.token
      const userInfo = data.userInfo
      this.$store.commit('SET_TOKEN', token)
      this.$store.commit('SET_USERINFO', userInfo)
      await this.$router.push('/home')
    }
  }
};
</script>

<style scoped>
.login-form-layout {
  position: absolute;
  left: 0;
  right: 0;
  width: 360px;
  margin: 140px auto;
  border-top: 10px solid #292c2f;
}

.input-row {
  display: flex;
}

.login-title {
  text-align: center;
}

.login-center-layout {
  background: #409eff;
  width: auto;
  height: auto;
  max-width: 100%;
  max-height: 100%;
  margin-top: 200px;
}
</style>
<template>
  <div class="login-wrap" v-if="showPage">
    <div class="ms-login">
      <div class="ms-title">动力电池数据分析系统</div>
      <div class="ms-content">
        <div class="mgb20">
          <el-input
            v-model="loginData.userName"
            prefix-icon="el-icon-user"
            placeholder="帐号"
            maxlength="14"
          ></el-input>
        </div>
        <div class="mgb20">
          <el-input
            type="password"
            placeholder="密码"
            prefix-icon="el-icon-lock"
            v-model="loginData.password"
            @keyup.enter.native="postLogin"
            maxlength="14"
          ></el-input>
        </div>
        <div class="login-btn">
          <el-button type="primary" @click="postLogin">登录</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import globals from '@/globals'

export default {
  name: 'Login',
  data: function() {
    return {
      showPage: false,
      loginData: {
        userName: '',
        password: ''
      }
    }
  },
  methods: {
    setUser(data) {
      this.$store.commit('setUserInfo', data)
      this.$store.commit('changeSidebar', data.userType)
      this.$router.push('/')
    },
    checkForm(form) {
      let err_msg = null
      if (form.userName === '') {
        err_msg = '用户名不能为空！'
      } else if (form.password === '') {
        err_msg = '密码不能为空！'
      } else if (!globals.RE_SIX_CHARACTER_CHECKER.test(form.userName)) {
        err_msg = '用户名须 5～14 位，并且只能是字母和数字！'
      } else if (!globals.RE_SIX_CHARACTER_CHECKER.test(form.password)) {
        err_msg = '密码须 5～14 位，并且只能是字母和数字！'
      }
      return err_msg
    },
    postLogin() {
      const err_msg = this.checkForm(this.loginData)
      if (err_msg !== null) {
        this.$message.error(err_msg)
        return
      }

      this.$axios
        .post(globals.URL_LOGIN, this.loginData)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }

          this.setUser(jd.data)
        })
        .catch(error => {
          this.$message.error(error.message)
        })
    }
  },
  beforeCreate() {
    // 因为是单页面应用，页面一旦刷新，标签页保存的用户信息就会丢失，
    // 通过 cookie 判断是否已经登录过了
    if (this.$store.state.userInfo === null) {
      this.$axios
        .get(globals.URL_LOGIN)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            this.showPage = true
            throw new Error(jd.msg)
          }

          this.setUser(jd.data)
        })
        .catch(() => {})
    }
  }
}
</script>

<style scoped>
.login-wrap {
  position: relative;
  width: 100%;
  height: 100%;
  background-image: url(../assets/img/login-bg.jpg);
  background-size: 100%;
}
.ms-title {
  width: 100%;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
  color: #fff;
  border-bottom: 1px solid #ddd;
}
.ms-login {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 350px;
  margin: -190px 0 0 -175px;
  border-radius: 5px;
  background: rgba(255, 255, 255, 0.3);
  overflow: hidden;
}
.ms-content {
  padding: 30px 30px;
}
.login-btn {
  text-align: center;
}
.login-btn button {
  width: 100%;
  height: 36px;
  margin-bottom: 10px;
}
</style>
<template>
  <el-row :gutter="20">
    <el-col :span="8">
      <el-card shadow="hover" class="mgb20">
        <div class="user-info">
          <img :src="userAvatar" class="user-avator" alt />
          <div class="user-info-cont">
            <div class="user-info-name">{{ userName }}</div>
            <div>{{ userRole }}</div>
          </div>
        </div>
        <div class="user-info-list">
          最后登录：
          <span>{{ userLastLoginTime }}</span>
        </div>
        <div class="user-info-list">
          登录次数：
          <span>{{ userLoginCount }}</span>
        </div>
      </el-card>
      <el-card shadow="hover" style="height:386px;">
        <div slot="header" class="clearfix">
          <h3>系统信息</h3>
        </div>
        <div class="mgb20 text">版本：V2.0</div>
        <div class="mgb20 text">开发：zzzzer91@gmail.com</div>
      </el-card>
    </el-col>
    <el-col :span="16">
      <el-calendar v-model="calendarDate"></el-calendar>
    </el-col>
  </el-row>
</template>

<script>
import globals from '@/globals'

export default {
  name: 'Dashboard',
  data() {
    return {
      calendarDate: new Date()
    }
  },
  computed: {
    userName() {
      return this.$store.state.userInfo.userName
    },
    userRole() {
      return this.$store.state.userInfo.userType === globals.USER_TYPE_SUPER_USER
        ? '超级管理员'
        : '普通用户'
    },
    userAvatar() {
      let avatarName = this.$store.state.userInfo.avatarName
      let imgUrl
      if (avatarName === null) {
        imgUrl = require('../assets/img/null.jpg')
      } else {
        imgUrl = `${globals.URL_AVATAR}/${avatarName}`
      }
      return imgUrl
    },
    userLastLoginTime() {
      return this.$store.state.userInfo.lastLoginTime
    },
    userLoginCount() {
      return this.$store.state.userInfo.loginCount
    }
  },
  methods: {}
}
</script>

<style scoped>
.user-info {
  display: flex;
  align-items: center;
  padding-bottom: 20px;
  border-bottom: 2px solid #ccc;
  margin-bottom: 20px;
}

.user-avator {
  width: 120px;
  height: 120px;
  border-radius: 50%;
}

.user-info-cont {
  padding-left: 50px;
  flex: 1;
  font-size: 14px;
  color: #999;
}

.user-info-cont div:first-child {
  font-size: 30px;
  color: #222;
}

.user-info-list {
  font-size: 14px;
  color: #999;
  line-height: 25px;
}

.user-info-list span {
  margin-left: 10px;
}

.text {
  font-size: 20px;
  color: #999;
}
</style>

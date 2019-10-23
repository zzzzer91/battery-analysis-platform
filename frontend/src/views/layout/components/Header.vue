<template>
  <div class="header">
    <!-- 折叠按钮 -->
    <div class="collapse-btn" @click="collapseChage">
      <i class="el-icon-s-unfold" :class="rotateIcon"></i>
    </div>
    <div class="logo">动力电池数据分析系统</div>
    <div class="header-right">
      <div class="header-user-con">
        <!-- 全屏显示 -->
        <div class="btn-fullscreen" @click="handleFullScreen">
          <el-tooltip effect="dark" :content="fullscreen?`取消全屏`:`全屏`" placement="bottom">
            <i class="el-icon-full-screen"></i>
          </el-tooltip>
        </div>
        <!-- 用户头像 -->
        <div class="user-avator">
          <img :src="userAvatar" />
        </div>
        <!-- 用户名下拉菜单 -->
        <el-dropdown class="user-name" trigger="click" @command="handleCommand">
          <span class="el-dropdown-link">
            {{ userName }}
            <i class="el-icon-caret-bottom"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>设置</el-dropdown-item>
            <el-dropdown-item divided command="logout">注销</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script>
import globals from '@/globals'

export default {
  name: 'Header',
  data() {
    return {
      fullscreen: false,
      rotateIcon: ''
    }
  },
  computed: {
    userName() {
      return this.$store.state.userInfo.userName
    },
    userAvatar() {
      let avatarName = this.$store.state.userInfo.avatarName
      let imgUrl
      if (avatarName === null) {
        imgUrl = require('@/assets/img/null.jpg')
      } else {
        imgUrl = `${globals.URL_AVATAR}/${avatarName}`
      }
      return imgUrl
    }
  },
  methods: {
    // 用户名下拉菜单选择事件
    handleCommand(command) {
      if (command == 'logout') {
        this.$axios
          .post(globals.URL_LOGOUT)
          .then(response => response.data)
          .then(jd => {
            if (jd.code !== globals.SUCCESS) {
              throw new Error('注销失败！')
            }

            this.$store.commit('setUserInfo', null)
            this.$router.push(globals.URL_LOGIN)
          })
          .catch(error => {
            this.$message.error(error.message)
          })
      }
    },
    // 侧边栏折叠
    collapseChage() {
      this.$store.commit('collapseChage')
      this.rotateIcon = this.rotateIcon === 'left-rotate-icon' ? 'right-rotate-icon' : 'left-rotate-icon'
    },
    // 全屏事件
    handleFullScreen() {
      let element = document.documentElement
      if (this.fullscreen) {
        if (document.exitFullscreen) {
          document.exitFullscreen()
        } else if (document.webkitCancelFullScreen) {
          document.webkitCancelFullScreen()
        } else if (document.mozCancelFullScreen) {
          document.mozCancelFullScreen()
        } else if (document.msExitFullscreen) {
          document.msExitFullscreen()
        }
      } else {
        if (element.requestFullscreen) {
          element.requestFullscreen()
        } else if (element.webkitRequestFullScreen) {
          element.webkitRequestFullScreen()
        } else if (element.mozRequestFullScreen) {
          element.mozRequestFullScreen()
        } else if (element.msRequestFullscreen) {
          // IE11
          element.msRequestFullscreen()
        }
      }
      this.fullscreen = !this.fullscreen
    }
  }
}
</script>

<style scoped>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  color: #fff;
  background-color: #242f42;
}
.collapse-btn {
  float: left;
  padding: 0 21px;
  cursor: pointer;
  line-height: 70px;
}
@keyframes leftRotatefresh {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(180deg);
  }
}
.left-rotate-icon {
  animation: 0.5s leftRotatefresh 1;
  animation-fill-mode: forwards; /*动画结束后保持状态*/
}
@keyframes rightRotatefresh {
  from {
    transform: rotate(180deg);
  }
  to {
    transform: rotate(0deg);
  }
}
.right-rotate-icon {
  animation: 0.5s rightRotatefresh 1;
}
.collapse-btn:hover {
  color: rgb(32, 160, 255);
}
.header .logo {
  float: left;
  line-height: 70px;
}
.header-right {
  float: right;
  padding-right: 50px;
}
.header-user-con {
  display: flex;
  height: 70px;
  align-items: center;
}
.btn-bell,
.btn-fullscreen {
  position: relative;
  height: 20px;
  text-align: center;
  border-radius: 15px;
  cursor: pointer;
}
.btn-bell-badge {
  position: absolute;
  right: 0;
  top: -2px;
  width: 8px;
  height: 8px;
  border-radius: 4px;
  background: #f56c6c;
  color: #fff;
}
.btn-bell .el-icon-bell {
  color: #fff;
}
.user-name {
  margin-left: 10px;
}
.user-avator {
  margin-left: 20px;
}
.user-avator img {
  display: block;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.el-dropdown-link {
  color: #fff;
  cursor: pointer;
}
</style>

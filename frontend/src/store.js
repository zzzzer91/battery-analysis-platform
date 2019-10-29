import Vue from 'vue'
import Vuex from 'vuex'
import globals from './globals'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    userInfo: null,
    collapse: true,
    keepAliveRouteList: [], // 组件缓存列表
    // 侧边栏内容
    sidebarItems: []
  },
  mutations: {
    // 设置用户信息
    setUserInfo(state, data) {
      state.userInfo = data
    },
    // 侧边栏折叠
    collapseChage(state) {
      state.collapse = !state.collapse
    },
    // 根据用户类型改变侧边栏
    changeSidebar(state, userType) {
      state.sidebarItems = [
        {
          icon: 'el-icon-house',
          index: 'dashboard',
          title: '系统首页'
        },
        {
          icon: 'el-icon-s-data',
          index: 'mining',
          title: '数据分析',
          subs: [
            {
              index: 'mining-base',
              title: '数据查询'
            },
            {
              index: 'mining-tasks',
              title: '计算任务'
            },
            {
              index: 'mining-test',
              title: '测试'
            }
          ]
        },
        {
          icon: 'el-icon-aim',
          index: 'deep-learning',
          title: '深度学习'
        },
        {
          icon: 'el-icon-document',
          index: 'markdown',
          title: 'Markdown',
        }
      ]
      if (userType === globals.USER_TYPE_SUPER_USER) {
        state.sidebarItems.push({
          icon: 'el-icon-user',
          index: 'user-manager',
          title: '用户管理'
        })
        state.sidebarItems.push({
          icon: 'el-icon-warning-outline',
          index: 'errors',
          title: '错误处理',
          subs: [
            {
              index: 'permission',
              title: '权限测试'
            },
            {
              index: '403',
              title: '403页面'
            },
            {
              index: '404',
              title: '404页面'
            }
          ]
        })
      }
    },
    // 标签
    addKeepAliveRoute(state, name) {
      state.keepAliveRouteList.push(name)
    },
    removeKeepAliveRoute(state, name) {
      const index = state.keepAliveRouteList.indexOf(name)
      state.keepAliveRouteList.splice(index, 1)
    },
    clearAllKeepAliveRoute(state) {
      state.keepAliveRouteList = []
    },
    clearOtherKeepAliveRoute(state, name) {
      let l = state.keepAliveRouteList
      const index = l.indexOf(name)
      state.keepAliveRouteList = l.splice(index, 1)
    },
  },
  actions: {
  }
})

import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import './plugins/axios'
import './plugins/elementui'
import './plugins/vcharts'

Vue.config.productionTip = false

// 路由拦截，在改变任意 url 时执行
router.beforeEach((to, from, next) => {
  // 这里的meta就是我们刚刚在路由里面配置的meta
  if (to.meta.requireAuth) {
    // 下面这个判断有没有登录
    if (store.state.userInfo !== null) {
      // 登录就继续
      next()
    } else {
      // 没有登录跳转到登录页面，登录成功之后再返回到之前请求的页面
      next('/login')
    }
  } else {
    // 不需要登录的，可以继续访问
    next()
  }
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/',
      redirect: '/dashboard',
      component:  () => import('@/views/layout/Layout.vue'),
      children: [
        {
          path: 'dashboard',
          component: () => import('@/views/Dashboard.vue'),
          meta: { title: '系统首页', requireAuth: true }
        },
        {
          path: 'mining-base',
          component: () => import('@/views/mining/Base.vue'),
          meta: { title: '数据查询', requireAuth: true }
        },
        {
          path: 'mining-tasks',
          component: () => import('@/views/mining/Tasks.vue'),
          meta: { title: '计算任务', requireAuth: true }
        },
        {
          path: 'mining-working-condition',
          component: () => import('@/views/mining/WorkingCondition.vue'),
          meta: { title: '工况', requireAuth: true }
        },
        {
          path: 'mining-test',
          component: () => import('@/views/mining/ChartTest.vue'),
          meta: { title: '测试', requireAuth: true }
        },
        {
          path: 'deep-learning',
          component: () => import('@/views/machine_learning/DeepLearning.vue'),
          meta: { title: '深度学习', requireAuth: true }
        },
        {
          path: 'markdown',
          component: () => import('@/views/Markdown.vue'),
          meta: { title: 'Markdown', requireAuth: true }
        },
        {
          path: 'user-manager',
          component: () => import('@/views/UserManager.vue'),
          meta: { title: '用户管理', requireAuth: true }
        },
        {
          // 权限页面
          path: 'permission',
          component: () => import('@/views/Permission.vue'),
          meta: { title: '权限测试', requireAuth: true }
        },
        // 只在单页内部标签中使用
        {
          path: '403',
          component: () => import('@/views/errors/Forbidden.vue'),
          meta: { title: '403', requireAuth: true }
        },
        {
          path: '404',
          component: () => import('@/views/errors/PageNotFound.vue'),
          meta: { title: '404', requireAuth: true }
        }
      ]
    },
    {
      path: '*',
      redirect: '/404'
    }
  ]
})

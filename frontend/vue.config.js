module.exports = {
  publicPath: '/',  // 根路径
  outputDir: 'dist',  // 构建输出目录
  assetsDir: 'assets',  // 静态文件目录
  lintOnSave: false,  // 是否开启 ESLint 保存检测
  devServer: {
    open: true,
    host: 'localhost',
    port: 8081,
    https: false,
    hotOnly: false,
    // proxy: 'http://localhost:3389', // 设置代理，指向生产服务器，获取数据
    proxy: {
      '/media': {
         target: 'http://localhost:3389',
      },
      '/login': {
         target: 'http://localhost:3389',
      },
      '/logout': {
         target: 'http://localhost:3389',
      },
      '/api': {
         target: 'http://localhost:3389',
      },
      '/file': {
         target: 'http://localhost:3389',
      },
      '/websocket': {
         target: 'ws://localhost:3389',
         ws: true
      },
    },
  },
  transpileDependencies: [
    'vue-echarts',
    'resize-detector'
  ]
}
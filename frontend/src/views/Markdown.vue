<template>
  <div>
    <mavon-editor
      v-model="content"
      ref="md"
      @change="change"
      style="min-height: 640px"
      :toolbars="toolbars"
    />
    <el-button class="editor-btn" type="primary" @click="downloadPDF">导出PDF</el-button>
  </div>
</template>

<script>
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

export default {
  name: 'Markdown',
  data: function() {
    return {
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        strikethrough: true, // 中划线
        mark: true, // 标记
        superscript: true, // 上角标
        subscript: true, // 下角标
        quote: true, // 引用
        ol: true, // 有序列表
        ul: true, // 无序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        code: true, // code
        table: true, // 表格
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        htmlcode: false, // 展示html源码
        help: false, // 帮助
        /* 1.3.5 */
        undo: true, // 上一步
        redo: true, // 下一步
        trash: false, // 清空
        save: false, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true, // 导航目录
        /* 2.1.8 */
        alignleft: true, // 左对齐
        aligncenter: true, // 居中
        alignright: true, // 右对齐
        /* 2.2.1 */
        subfield: true, // 单双栏模式
        preview: true // 预览
      },
      content: '',
      html: '',
      configs: {}
    }
  },
  components: {
    mavonEditor
  },
  methods: {
    // 将图片上传到服务器，返回地址替换到md中
    // $imgAdd(pos, $file) {
    //   var formdata = new FormData()
    //   formdata.append('file', $file)
    //   // 这里没有服务器供大家尝试，可将下面上传接口替换为你自己的服务器接口
    //   this.$axios({
    //     url: '/common/upload',
    //     method: 'post',
    //     data: formdata,
    //     headers: { 'Content-Type': 'multipart/form-data' }
    //   }).then(url => {
    //     this.$refs.md.$img2Url(pos, url)
    //   })
    // },
    change(value, render) {
      // render 为 markdown 解析后的结果
      this.html = render
    },
    downloadPDF() {
      let mywindow = window.open('', 'PRINT', 'height=400,width=600')
      const content = document.getElementsByClassName('v-show-content')[0]
        .innerHTML
      mywindow.document.write(
        `<!DOCTYPE html>
        <html lang="en">
          <head>
            <meta charset="utf-8">
            <title>${document.title}</title>
            <link rel="stylesheet" type='text/css' href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.9.0/github-markdown.min.css" />
            <link rel="stylesheet" type='text/css' href="https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.8.3/katex.min.css" />
            <link rel="stylesheet" type='text/css' href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/github.min.css" />
            <style>
              .markdown-body .hljs-left {
                text-align: left;
              }
              .markdown-body .hljs-center {
                text-align: center;
              }
              .markdown-body .hljs-right {
                text-align: right;
              }
            </style>
          </head>
          <body>
            <div class="markdown-body">
              ${content}
            </div>
          </body>
          <script src="https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.8.3/katex.min.js" async><\/script>
          <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js" async><\/script>
        </html>`
      )

      mywindow.document.close() // necessary for IE >= 10
      mywindow.focus() // necessary for IE >= 10*/

      // 等待资源加载完毕
      mywindow.onload = () => {
        mywindow.print()
        mywindow.close()
      }

      return true
    }
  }
}
</script>

<style scoped>
.editor-btn {
  margin-top: 20px;
}
</style>
<template>
  <div class="container">
    <h1>超级管理员权限测试</h1>
    <p>只有拥有超级管理员权限的帐号，才能进到这个页面，其他账号都会跳到403页面。</p>
  </div>
</template>

<script>
export default {
  name: 'Permission',
  data: function() {
    return {}
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
  margin: 30px 0;
}
p {
  line-height: 30px;
  margin-bottom: 10px;
  text-indent: 2em;
}
</style>
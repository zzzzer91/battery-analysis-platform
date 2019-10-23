<template>
  <div class="wrapper">
    <v-header/>
    <v-sidebar/>
    <div class="content-box" :class="{'content-collapse':collapse}">
      <v-tags/>
      <div class="content">
        <transition name="move" mode="out-in">
          <keep-alive :include="keepAliveRouteList">
            <router-view/>
          </keep-alive>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
import vHeader from './components/Header.vue'
import vSidebar from './components/Sidebar.vue'
import vTags from './components/Tags.vue'

export default {
  name: 'Layout',
  components: {
    'v-header': vHeader,
    'v-sidebar': vSidebar,
    'v-tags': vTags
  },
  computed: {
    collapse() {
      return this.$store.state.collapse
    },
    keepAliveRouteList() {
      return this.$store.state.keepAliveRouteList
    }
  }
}
</script>

<style scoped>
.wrapper {
    width: 100%;
    height: 100%;
    overflow: hidden;
}

.content-box {
  position: absolute;
  left: 250px;
  right: 0;
  top: 70px;
  bottom: 0;
  padding-bottom: 30px;
  -webkit-transition: left .3s ease-in-out;
  transition: left .3s ease-in-out;
  background: #f0f0f0;
}

.content {
  width: auto;
  height: 100%;
  padding: 10px;
  overflow-y: scroll;
  box-sizing: border-box;
}

.content-collapse {
  left: 65px;
}
</style>

<script>
export default {
  name: "NaviBar",
};
</script>

<template>
  <el-menu class="navi-bar" mode="horizontal" :ellipsis="false">
    <el-menu-item index="/">
      <router-link to="/">Home</router-link>
    </el-menu-item>
    <el-menu-item index="/upload">
      <router-link to="/upload">Upload</router-link>
    </el-menu-item>
    <el-menu-item index="/about">
      <router-link to="/about">About</router-link>
    </el-menu-item>
  </el-menu>
</template>

<style scoped>
.navi-bar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
}
</style>

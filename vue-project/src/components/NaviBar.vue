<script>
export default {
  name: "NaviBar",
};
</script>

<template>
  <el-menu class="navi-bar" mode="horizontal" :ellipsis="false" router>
    <el-menu-item index="/">Home</el-menu-item>
    <el-menu-item index="/upload">Upload</el-menu-item>
    <el-menu-item index="/about">About</el-menu-item>
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

<script>
import NaviBar from "@/components/NaviBar.vue";
export default {
  name: "App",
  components: { NaviBar },
};
</script>

<template>
  <div id="app">
    <NaviBar />
    <router-view />
  </div>
</template>

<style scoped></style>

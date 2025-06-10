<script>
import {getVideos} from "@/api/video/index.js";

export default {
  name: "home",
  data() {
    return {
      videos: [],
    };
  },
  methods: {
    load() {
      getVideos().then((res) => {
        this.videos = res.data;
      });
    },
  },
  components: {},
  beforeMount() {
    this.load();
  },
};
</script>

<template>
  <div class="home">
    <el-row :gutter="20">
      <el-col v-for="video in videos" :key="video.id" :span="6">
        <el-card>
          <img class="video-cover"
               src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png">
          <span>{{ video.title }}</span>
          <div class="bottom clearfix">
            <span>{{ video.description }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.video-cover {
  width: 100%;
}
</style>

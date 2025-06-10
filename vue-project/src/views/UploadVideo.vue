<script>
import {uploadVideo} from "@/api/video/index.js";

export default {
  name: "UploadVideo",
  data() {
    return {
      form: {
        title: "",
        description: "",
      },
    };
  },
  methods: {
    onSubmit() {
      console.log(import.meta.env.VITE_APP_URL);
      if (!this.form.title.trim()) {
        this.$notify.error({
          title: "Error",
          message: "Video title cannot be empty",
        });
        return;
      }
      uploadVideo(this.form)
          .then((res) => {
            if (res.status > 0) {
              this.$notify.error({
                title: "Error",
                message: "The server is down",
              });
            } else {
              this.$notify({
                title: "Success",
                message: `Video ${res.data.id} Uploaded Successfully!`,
                type: "success",
              });
            }
          })
    },
  },
};
</script>

<template>
  <div class="upload-video">
    <h2>Upload Your Video</h2>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="Title">
        <el-input v-model="form.title"></el-input>
      </el-form-item>

      <el-form-item label="Description">
        <el-input v-model="form.description" type="textarea"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">Submit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style scoped></style>

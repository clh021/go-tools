<template>
  <div>
    <!-- 上传组件 -->
    <el-upload action drag :show-file-list="false" :http-request="handleUploadRequest">
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <!-- <div class="el-upload__tip" slot="tip">大小不超过 200M </div> -->
    </el-upload>

    <!-- 进度显示 -->
    <el-progress class="progress-bar" :text-inside="true" :stroke-width="24" :percentage="parseInt(percent , 10)" />
  </div>
</template>


<script>
import { uploadByPieces } from "./chunkUpload.js";

export default {
  data () {
    return {
      percent: 0,
      videoUrl: '',
      uploading: true,
      percentCount: 0
    }
  },
  methods: {
    handleUploadRequest (options) {
      return this.dealUpload(options);
    },
    dealUpload (options) {
      const file = options.file;
      this.percent = 0;
      return new Promise((resolve, reject) => {
        uploadByPieces({
          file: file,
          pieceSize: 5,
          chunkUrl: "/api/uploadChunk",
          fileUrl: "/api/uploadChunkDone",
          progress: (num) => {
            console.log("num", num);
            this.percent = num;
            options.onProgress({ percent: num });
          },
          success: (data) => {
            this.percent = 100;
            resolve(data);
          },
          error: (e) => {
            this.percent = 0;
            reject(e);
          },
        });
      });
    },
  }
}
</script>

<style scoped>
.progress-bar {
  margin: 5px 0;
}
</style>
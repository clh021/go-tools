<script setup>
defineProps({
  msg: {
    type: String,
    default: '正在检测分析……'
  }
})
</script>
<template>
  <div class="loading-box">
    <div class="loading-txt">
      {{ msg }}
    </div>
    <div class="loading-circle" v-loading="'loading'"></div>
  </div>
</template>
<style scoped>
.loading-box {
  display: flex;
  flex-direction: column;

  .loading-txt {
    text-align: center;
  }
  .loading-circle {
    margin-top:20px;
  }
}
</style>
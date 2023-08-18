<script setup>
// import TheWelcome from '../components/TheWelcome.vue'
import UploadChunk from '../components/UploadChunk.vue'
import AnalySisType from '../components/AnalysisType.vue'
import MyDivider from '../components/MyDivider.vue'
import MyPanel from '../components/MyPanel.vue'
// import MyTransition from '../components/MyTransition.vue'
import MyLoading from '../components/MyLoading.vue'
import { ref } from 'vue'
import axios from "axios";
import { formatBytes } from "@/utils";
// import { ElLoading } from 'element-plus'

const panelUploadResult = ref({})
const step = ref(0)
const fileClassObj = [
  { key: "bin", val: "二进制文件" },
  { key: "java", val: "Java包" },
  { key: "deb", val: "Deb包" },
  { key: "Rpm", val: "Rpm包" },
]
const userFileClass = ref("")
const showLoading = ref(false)
const alertType = ref("success")
const alertTitle = ref("")
const showAlert = ref(false)
const showAlertQuick = (title, type) => {
  alertTitle.value = title
  alertType.value = type
  showAlert.value = true
}
const loadMsg = ref("")
function userFileClassChange (param) {
  let filterUserClass = fileClassObj.filter(e => e.key == param)
  userFileClass.value = filterUserClass[0];
  step.value = 2;
  showAlertQuick("选中文件就开始上传", 'info');
}
const userUploadDone = ref({})
function uploadDone (param) {
  console.log("uploadDone:", param)
  userUploadDone.value = param;
  step.value = 5;
  panelUploadResult.value = {
    '文件名': param.fileName,
    '文件类型': param.fileType,
    '检测类型': userFileClass.value.val,
    '文件体积': formatBytes(param.fileSize),
    '文件MD5': param.fileMD5,
  };
  showAlert.value = false
}
function completeUpload () {
  step.value = 9
  showLoading.value = true;
  // sendToMinIO()
  console.log(sendToMinIO)
  // panelUploadResult.value = { "Username": "kooriookami", "Telephone": "18100000000", "Place": "Suzhou", "Remarks": "School", "Address": "No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province" };
}
function sendToMinIO (fileName, fileMD5, fileType) {
  let makeForm = new FormData();
  makeForm.append("md5", fileMD5);
  makeForm.append("file_name", fileName);
  makeForm.append("type", fileType);
  return axios({
    // 合并文件
    method: "post",
    url: "/api/sendToMinIO",
    data: makeForm,
  })
    .then((res) => {
      console.log(res)
    })
    .catch((e) => {
      console.error(e);
    });
}
</script>

<template>
  <main>
    <!-- step 1-9 -->
    <MyPanel title="上传" :step="step" :data="panelUploadResult" btn-txt="开始检测分析" :step-min="1" :step-max="9"
      :btn-step-min="5" v-on:submit="completeUpload">
      <AnalySisType :file-class-obj="fileClassObj" v-on:update:user-file-class="userFileClassChange" v-if="step < 9" />
      <MyDivider />
      <transition name="el-zoom-in-right">
        <UploadChunk v-show="step > 1" v-on:uploadDone="uploadDone">
          <div class="el-upload__text">
            将 <el-button round>{{ userFileClass.val }}</el-button> 拖到此处，或<em>点击此处</em>，自动上传。
          </div>
        </UploadChunk>
      </transition>
    </MyPanel>

    <el-alert v-show="showAlert" :title="alertTitle" :type="alertType" effect="dark" />
    <MyLoading v-show="showLoading" :msg="loadMsg"/>
  </main>
</template>

<style scoped>
</style>
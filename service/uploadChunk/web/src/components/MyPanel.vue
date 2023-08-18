<template>
  <div>
    <div v-show="step<stepMax">
      <slot></slot>
    </div>
    <el-descriptions v-show="step>=stepMin" class="margin-top" :title="title" :column="3" border>
      <template #extra>
        <el-button type="primary" v-show="step>=btnStepMin && step<stepMax" @click="completePanelOpera">{{ btnTxt }}</el-button>
      </template>
      <el-descriptions-item v-for="(value, key, index) in data" :key="`panel-des-${index}`">
        <template #label>
          <div class="cell-item">
            {{ key }}
          </div>
        </template>
        {{ value }}
      </el-descriptions-item>
    </el-descriptions>
  </div>
</template>

<script setup>
// import { ref } from 'vue'
defineProps({
  step: Number,
  stepMin: Number,
  stepMax: Number,
  data: {},
  title: String,
  btnTxt: String,
  btnStepMin: Number,
})
const emit = defineEmits(['submit'])
function completePanelOpera() {
  emit('submit', 1)
}
</script>

<style scoped>
.el-descriptions {
  margin-top: 20px;
}

.cell-item {
  display: flex;
  align-items: center;
}

.margin-top {
  margin-top: 20px;
}
</style>


<script setup>
import { ref } from 'vue'

// const props = defineProps(["fileClassObj"])// 'modelValue',
defineProps(["fileClassObj"])// 'modelValue',
const emit = defineEmits(['update:userFileClass'])
const mValue = ref("")
const userFileClassChanged = (param) => {
  emit('update:userFileClass', param)
}
</script>

<template>
  <div class="greetings">
    <el-radio-group @change="userFileClassChanged" v-model="mValue" label="label position">
      <el-radio-button v-for="(o, index) in fileClassObj" :key="`analysistype-radio-${index}`" :label="o.key">
        {{ o.val }}
      </el-radio-button>
    </el-radio-group>
  </div>
</template>

<style scoped>
.greetings {
  text-align: center;
}
</style>
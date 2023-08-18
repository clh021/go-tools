<script setup>
defineProps(['transitionName'])
</script>

<template>
  <!-- 包装内置的 Transition 组件 -->
  <Transition :name="transitionName">
    <slot></slot> <!-- 向内传递插槽内容 -->
  </Transition>
</template>

<style>
.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.25);
  }
  100% {
    transform: scale(1);
  }
}
</style>
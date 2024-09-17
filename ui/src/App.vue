<script setup>
import {RouterView} from "vue-router"
</script>

<template>
<!-- 这是根页面：所有的页面都需要挂载在 App.vue 中才能呈现出来 -->
<!--  特别是 Vue Router 的子页面 -->
  <div>

  <!-- 路由出口：路由匹配到的组件将在这里渲染 -->
    <RouterView/>
  </div>
</template>

<style scoped></style>

<script setup>
import {RouterView, RouterLink} from "vue-router"
</script>

<template>
  <!--    使用 RouterLink 组件进行视图切换-->
  <RouterLink :to="{path:'/'}"> Home</RouterLink>
  <RouterLink :to="{name:'PermissionDeny'}"> PermissionDeny</RouterLink>
  <!-- 这是根页面：所有的页面都需要挂载在 App.vue 中才能呈现出来 -->
  <!--  特别是 Vue Router 的子页面 -->
  <div class="content">
    <!-- 路由出口：路由匹配到的组件将在这里渲染 -->
    <RouterView></RouterView>
  </div>
</template>

<style scoped></style>

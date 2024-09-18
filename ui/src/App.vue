<script setup name="App">
import {RouterView} from "vue-router"
import Header from "@/components/Header.vue";
</script>

<template>

  <div class="main-container">

    <!-- 导航区 -->
    <Header></Header>


    <div class="main-page">
      <!-- 路由出口：路由匹配到的组件将在这里渲染 -->
      <RouterView></RouterView>
    </div>
  </div>

</template>

<style scoped>

.main-container {
  width: 90vw;
  height: calc(100vh - 48px);
  margin: auto;
}



.main-page {
  height: 100%;
  width: 100%;
  border-left: solid 2px #ccc;
  border-right: solid 2px #ccc;
  border-bottom: solid 2px #ccc;
}

</style>


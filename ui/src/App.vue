<script setup>
import {RouterView} from "vue-router"
</script>

<template>

  <div class="main-container">

    <!-- 导航区 -->
    <div class="navigation">

      <!-- Logo 显示区域 -->
      <div class="">
        杜宇鹏的个人博客系统
      </div>
      <!--登录操作区域-->
      <div>
        <a-space>
          <a-button>管理(登录后出现)</a-button>
          <a-button>登录(未登录出现)</a-button>
        </a-space>
      </div>
    </div>


    <div class="main-page">
      <!-- 路由出口：路由匹配到的组件将在这里渲染 -->
      <RouterView></RouterView>
    </div>

    <div class="main-footer">
    </div>

  </div>

</template>

<style scoped>

.main-container {
  width: 90vw;
  height: 90vh;
  margin: 0 auto;
}

.navigation {
  /*样式*/
  height: 45px;
  border-bottom: solid 2px #ccc;
  background-color: #F2F3F5;

  /* 布局 */
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-page {
  height: 100%;
  width: 100%;
  border-left: solid 2px #ccc;
  border-right: solid 2px #ccc;
}

.main-footer {
  border: solid 2px #ccc;
}
</style>


<script setup name="BackendLayout">

import {RouterView, useRouter, useRoute} from "vue-router"
import {reactive} from "vue";

const router = useRouter()

const menuItemClickHandler = (key) => {
  router.push({name: key})
}

</script>

<template>
  <div class="backend-container">
    <!-- 后台管理：菜单导航区 -->
    <div class="backend-menu">
      <a-menu
          :style="{height:'100%'}"
          :default-open-keys="['blog', 'comment']"
          :default-selected-keys="[useRoute().name]"
          @menu-item-click="menuItemClickHandler"
      > <!-- default-open-keys：默认打开的菜单 default-selected-key：默认选中的菜单项，通过当前路由的名字进行选择-->
        <a-sub-menu key="blog">
          <template #icon>
            <icon-book/>
          </template>
          <template #title>文章管理</template>
          <a-menu-item key="backend-blogs">文章列表</a-menu-item> <!-- key 使用路由的 name 参数指定，方便后续使用 -->
          <a-menu-item key="backend-tags">标签管理</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="comment">
          <template #icon>
            <icon-align-left/>
          </template>
          <template #title>评论管理</template>
          <a-menu-item key="backend-comments">评论列表</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </div>

    <!-- 后台管理：操作区域 -->
    <div class="backend-content">
      <!-- 后台子路由的页面 -->
      <RouterView></RouterView>
    </div>

  </div>
</template>

<style scoped>
.backend-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: space-between;
}

.backend-menu {
  width: 20%;
  height: 100%;
  border-right: solid 2px #ccc;
}

.backend-content {
  width: 80%;
}
</style>

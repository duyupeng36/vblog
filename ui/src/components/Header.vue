<script setup name="Header">

import {useRouter, useRoute} from "vue-router"
import loginState from "@/stores/login"

const router = useRouter()

const currentRoute = useRoute()

function Logout() {
  loginState.value.isLogin = false
  loginState.value.username = ""
  router.push({name: "frontend"})
}

function Login() {
  router.push({name: "login"})
}

function JumpToBackend() {
  router.push({name: "backend"})
}

function JumpToFrontend() {
  router.push({name: "frontend"})
}

</script>

<template>
<div class="navigation">
  <!-- Logo 显示区域 -->
  <div class="">
    杜宇鹏的个人博客系统
  </div>
  <!--登录操作区域-->
  <div>
    <a-space>
      <a-button v-if="loginState.isLogin && currentRoute.fullPath.startsWith('/frontend')" @click="JumpToBackend">后台管理</a-button>
      <a-button v-if="loginState.isLogin && currentRoute.fullPath.startsWith('/backend')" @click="JumpToFrontend">前台浏览</a-button>
      <a-button v-if="loginState.isLogin" @click="Logout">注销</a-button>
      <a-button v-if="!loginState.isLogin" @click="Login">登录</a-button>
    </a-space>
  </div>
</div>
</template>

<style scoped>
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
</style>
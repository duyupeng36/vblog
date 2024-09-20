<script setup>

import {reactive} from "vue";
import {Message} from "@arco-design/web-vue";
import {useRouter} from "vue-router";

import loginState from "@/stores/login.js"


const form = reactive({
  username: "",
  password: "",
})

const router = useRouter()

const handleSubmit = (data) => {
  if (!data.errors) {

    // 对接后端 API
    if (data.values.username === "admin" && data.values.password === "123456") {
      console.log("登录成功")

      // 保存一个全局状态，最好的方式使用 localStorage 保存登录状态，方便其他标签页或组件
      loginState.value.username = data.values.username
      loginState.value.isLogin = true

      // 跳转到 后台
      router.push({name: "backend"})
    } else {
      Message.error("用户名或密码错误")
      return false;
    }
  }
}

</script>

<template>
 <div class="login-container">
   <a-form :model="form" :style="{width:'600px'}" @submit="handleSubmit">
     <a-form-item>
       <div class="login-title">
         欢迎登录
       </div>
     </a-form-item>
     <a-form-item
         field="username"
         tooltip="请输入用户名"
         label="用户名"
         :rules="[{required:true, message:'用户名是必须的'}]"
         :validate-trigger="['input']"
     >
       <a-input
           v-model="form.username"
           placeholder="请输入用户名..."
       />
     </a-form-item>
     <a-form-item
         field="password"
         tooltip="请输入密码"
         label="密&emsp;码"
         :rules="[{required:true, message:'密码是必须的'}, {minLength:6, message:'密码不能低于 6 位'}]"
         :validate-trigger="['input']"
     >
       <a-input-password
           v-model="form.password"
           placeholder="请输入密码..."
           :defaultVisibility="false"
           allow-clear />
     </a-form-item>
     <a-form-item>
       <a-button html-type="submit" :style="{width: '100%'}">登录</a-button>
     </a-form-item>
   </a-form>
 </div>
</template>

<style scoped>

.login-container {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-title {
  width: 100%;
  text-align: center;
  font-size: 4rem;
  color: #0e0630;
}
</style>
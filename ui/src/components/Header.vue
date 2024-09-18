<script setup name="Header">

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
      <a-button>管理(登录后出现)</a-button>
      <a-button>登录(未登录出现)</a-button>
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
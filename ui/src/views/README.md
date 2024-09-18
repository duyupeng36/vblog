# 视图组件

一个视图就是一个 **Vue 组件**。允许我们将 UI 划分为独立的、可重用的部分，并且可以对每个部分进行单独的思考

使用构建步骤时，一个 Vue 组件可以包含 $3$ 个标签

```html
<script setup>
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  <button @click="count++">You clicked me {{ count }} times.</button>
</template>

<style scoped>
</style>
```

**Vue 组件** 本质上时一个默认导出的 JS 模块

```js
import { ref } from 'vue'

export default {
  setup() {
    const count = ref(0)
    return { count }
  },
  template: `
    <button @click="count++">
      You clicked me {{ count }} times.
    </button>`
  // 也可以针对一个 DOM 内联模板：
  // template: '#my-template-element'
}
```

## 异常视图

vblog 前端部分需要补充下面两种 **异常视图**

+ `403` 视图: 用户 **未登陆** 就访问管理页面时
+ `404` 视图: 用户输入的 **URL 并没有匹配页面**

### 403 视图

使用 [Arco-Design-result#403](https://arco.design/vue/component/result#403) 组件进行封装。详细内容在 [PermissionDeny.vue](errors/PermissionDeny.vue)


### 404 页面

我们使用 [Arco-Design-result#404](https://arco.design/vue/component/result#404) 组件进行封装。详细内容在 [NotFound.vue](errors/NotFound.vue)


只有视图组件还不能在 Vue 中工作，还需要将异常视图 [注册到路由表中](../router/README.md#注册异常组件的路由)

## 布局

### 后台

博客后台使用的布局模版: [BackendLayout.vue](./backend/BackendLayout.vue)

这里我们需要使用到侧边栏导航:[Arco Design菜单 Menu](https://arco.design/vue/component/menu)

#### 布局页面

```html
<script setup>

import {RouterView} from "vue-router"

</script>

<template>
  <div class="backend-container">
    <!-- 后台管理：菜单导航区 -->
    <div class="backend-menu">
      <a-menu
          :style="{height:'100%'}"
          :default-open-keys="['0', '1']"
      > <!-- default-open-keys：默认打开的菜单 default-selected-key：默认选中的菜单项-->
        <a-sub-menu key="0">
           <template #icon>
             <icon-book/>
           </template>
          <template #title>文章管理</template>
          <a-menu-item key="0_0">文章列表</a-menu-item>
          <a-menu-item key="0_1">标签管理</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="1">
           <template #icon>
             <icon-align-left />
           </template>
          <template #title>评论管理</template>
          <a-menu-item key="1_0">评论列表</a-menu-item>
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
```
#### 博客列表页面

后台管理的 [博客列表](./backend/blogs/BlogList.vue)

```html
<script setup>

</script>

<template>
  <div>
    文章列表
  </div>
</template>

<style scoped>

</style>
```

#### 标签列表页面

后台管理的 [标签列表](./backend/tags/TagList.vue)

```html
<script setup>

</script>

<template>
<div>
  标签列表
</div>
</template>

<style scoped>

</style>
```

#### 评论列表页面

后台管理的 [评论列表](./backend/comments/CommentList.vue)

```html
<script setup>

</script>

<template>
<div>
  评论列表
</div>
</template>

<style scoped>

</style>
```

#### 补充后台管理相关路由

```js
        //   后台页面
{
    path: "/backend",
        name: "backend",
    component: () => import('@/views/backend/BackendLayout.vue'),
    children: [
    {
        path: '',
        name: 'statistics',
        component: () => import("@/views/backend/Statistics.vue")

    },
    {
        path: 'blogs',
        name: 'backend-blogs',
        component: () => import("@/views/backend/blogs/BlogList.vue")

    },
    {
        path: 'tags',
        name: 'backend-tags',
        component: () => import("@/views/backend/tags/TagList.vue")

    },
    {
        path: 'comments',
        name: 'backend-comments',
        component: () => import("@/views/backend/comments/CommentList.vue")
    }
]
}
```

> `name` 属性必须全局唯一：因为需要通过 `name` 进行页面切换

#### 路由对接到菜单页

给 `a-menu` 标签绑定一个 `@menu-item-click="menuItemClickHandler"` 事件

```js

import {useRouter} from "vue-router"

const router = useRouter() // 获取路由实例

const menuItemClickHandler = (key) => {
  router.push({ name: key }) // 跳转到点击的标签对应的视图
}
```

### 前台

前台布局模版: [FrontendLayout.vue](./frontend/FrontendLayout.vue)


#### 布局页面

```html
<script setup>

import {RouterView} from "vue-router"

</script>

<template>
  <div class="frontend-container">
    <!-- 前台子路由的页面 -->
    <RouterView></RouterView>
  </div>
</template>

<style scoped>

</style>
```







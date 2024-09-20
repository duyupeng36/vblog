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

#### BlogView 页面

展示博客内容的页面

```html
<script setup name="BlogView">

</script>

<template>
  <main>
    博客页面
  </main>
</template>

<style scoped>

</style>
```

#### BlogList 页面

博客列表页面

```html
<script setup name="BlogList">

</script>

<template>
  <div>
    Blogs List
  </div>
</template>

<style scoped>

</style>
```

#### 前台路由

```js
//   前台页面
{
    path: "/frontend",
    name: "frontend",
    component: () => import('@/views/frontend/FrontendLayout.vue'),
    children: [
        {
            path: '',
            name: 'frontend-blogs',
            component: () => import("@/views/frontend/BlogList.vue")
        },
        {
            path: ':id',
            name: 'blog-content',
            component: () => import("@/views/frontend/BlogView.vue")
        },

    ]
}
```

## 前后台切换

用户未登录的时候是不应该让其查看后台页面的。因此，我们首先补充一个 **登录页面**

```html
<script setup name="Header">

    import {useRouter} from "vue-router"
    import loginState from "@/stores/login"

    const router = useRouter()

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
                <a-button v-if="loginState.isLogin && router.currentRoute.value.fullPath.startsWith('/frontend')" @click="JumpToBackend">后台管理</a-button> <!-- 用户已登录并且当前是在前台，就展示后台管理 -->
                <a-button v-if="loginState.isLogin && router.currentRoute.value.fullPath.startsWith('/backend')" @click="JumpToFrontend">前台浏览</a-button> <!-- 当前是在后台台，就展示前台浏览 -->
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
```

登录状态需要持久化，即 [状态管理](../stores/README.md#vueuse与本地存储)。登录状态定义在 [login.js](../stores/login.js)，这采用全局变量，哪里需要就在那里导入即可

```html
<script setup name="Header">

import {useRouter} from "vue-router"
import {loginState} from "@/stores/login.js"

const router = useRouter()


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
      <a-button v-if="loginState.isLogin" @click="JumpToBackend">后台管理</a-button>
      <a-button v-if="loginState.isLogin" @click="JumpToFrontend">前台浏览</a-button>
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
```


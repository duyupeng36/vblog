# vblog 前端部分

## 项目初始化

这里我们直接使用 vue cli 初始化我们的骨架, 然后在此基础上进行修改

```shell
npm create vue@latest

> npx
> create-vue


Vue.js - The Progressive JavaScript Framework

√ 请输入项目名称： ... ui
√ 是否使用 TypeScript 语法？ ... 否 / 是             # 否
√ 是否启用 JSX 支持？ ... 否 / 是                    # 否
√ 是否引入 Vue Router 进行单页面应用开发？ ... 否 / 是  # 是
√ 是否引入 Pinia 用于状态管理？ ... 否 / 是            # 是
√ 是否引入 Vitest 用于单元测试？ ... 否 / 是           # 否
√ 是否要引入一款端到端（End to End）测试工具？ » 不需要
√ 是否引入 ESLint 用于代码质量检测？ ... 否 / 是        # 是
√ 是否引入 Prettier 用于代码格式化？ ... 否 / 是        # 是
√ 是否引入 Vue DevTools 7 扩展用于调试? (试验阶段) ... 否 / 是  # 否

正在初始化项目 C:\Users\duyup\Documents\vblog\ui...

项目初始化完成，可执行以下命令：

  cd ui
  npm install
  npm run format
  npm run dev
```

设置项目源，我们需要安装依赖，`npm` 默认使用的是国外源，下载依赖可能会失败。使用 `nrm` 查看 `npm` 支持的源

```shell
nrm ls
  npm ---------- https://registry.npmjs.org/
  yarn --------- https://registry.yarnpkg.com/
  tencent ------ https://mirrors.cloud.tencent.com/npm/
  cnpm --------- https://r.cnpmjs.org/
* taobao ------- https://registry.npmmirror.com/
  npmMirror ---- https://skimdb.npmjs.com/registry/
```

通过 `nrm test` 测试所有源

```shell
nrm test
  npm ---------- timeout (Fetch timeout over 5000 ms)
  yarn --------- 2304 ms
  tencent ------ 937 ms
  cnpm --------- 1808 ms
* taobao ------- 56 ms
  npmMirror ---- timeout (Fetch timeout over 5000 ms)

```

通过 `nrm use` 设置最快的源

```shell
nrm use taobao
```

## 清理模版页面

在做 `Home` 页面之前，先清理掉脚手架为我们生成的页面，保留空白页面

删除不需要的组件：删除 `src/components` 中的所有组件，删除 `src/views` 中的 `AboutView.vue`

清理 `App.vue` 和 `HomeView.vue` 中不需要的内容

`App.vue` 中保留的内容

```html
<script setup></script>

<template>
  <div>App</div>
</template>

<style scoped></style>
```

`HomeView.vue` 中保留的内容

```html
<script setup></script>

<template>
  <div>Home</div>
</template>
```

## 清理样式

删除 `src/assets/base.css` 和 `src/assets/main.css` 中的所有内容

添加 [`normalize.css`](https://necolas.github.io/normalize.css/)

设置 `base.css`

```css
* {
  box-sizing: border-box;
}

html,
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-size: 14px;
  background-color: var(--color-bg-1);
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
}

#app {
  width: 100%;
  height: 100%;
}
```

在 `main.css` 中导入

```css
@import 'normalize.css'; /* 导入 normalize.css */
@import 'base.css'; /*导入 base.css*/
```

## UI组件

安装 [arco design](https://arco.design/) UI 组件

```shell
 npm install --save-dev @arco-design/web-vue
```

在 [main.js](src/main.js) 中引入 UI 组件

```js
import { createApp } from 'vue'
import App from './App.vue'

// 加载组件
import ArcoVue from '@arco-design/web-vue'
// 引入 UI 组件的样式
import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)

// 使用 UI 组件：这里会见所有的 UI 组件注册到 Vue 的全局组件列表中
app.use(ArcoVue)
app.use(createPinia())
app.use(router)

app.mount('#app')
```

### 测试

我们在 `App.vue` 中引入下面的页面进行测试

```html
<script setup></script>

<template>
  <div>
    <a-space>
      <a-button type="primary">Primary</a-button>
      <a-button>Secondary</a-button>
      <a-button type="dashed">Dashed</a-button>
      <a-button type="outline">Outline</a-button>
      <a-button type="text">Text</a-button>
    </a-space>
  </div>
</template>

<style scoped></style>
```

只要能显示 $5$ 个按钮就证明 UI 组件引入成功

### Icon 组件

**Arco 图标是一个独立的库**，需要额外引入并注册使用

```js
import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// 加载组件
import ArcoVue from '@arco-design/web-vue';
// 额外引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
// 引入 UI 组件的样式
import '@arco-design/web-vue/dist/arco.css';



const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ArcoVue)  // 注册 Arco UI 组件
app.use(ArcoVueIcon) // 注册 Arco Icon 组件

app.mount('#app')
```

## 异常页面

这里我们需要补充 $2$ 种 [异常页面](src/views/README.md)

## 路由

为了切换和展示页面，需要引入 [路由](./src/router/README.md)

## 布局

vblog 界面分为 **前台** 和管理 **后台**

博客前台样式如下

![front.png](.images/front.png)

博客管理后台:

![background.png](.images/background.png)

布局需要 [嵌套路由](./src/router/README.md#嵌套路由) 的支持

无论前台还是后台，都需要一个 **导航栏** 。因此，我们在 `App.vue` 中布局导航栏

```html
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
        <!-- 页面展示区 -->
        <div class="main-page">
            <!-- 路由出口：路由匹配到的组件将在这里渲染 -->
            <RouterView></RouterView>
        </div>
        <!-- 页脚 -->
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
```

因为前台与后台布局样式不一样, 因此分别使用独立布局模版。详细内容在 [布局](./src/views/README.md#布局) 中介绍

## 状态管理

[状态管理](./src/stores/README.md)

## 导航守卫

之前的流程 用户通过前台 跳转到后台管理时, 补充了认证, 但是如果用户直接通过 URL 访问后台管理页面喃?

这时候我们就需要做一个全局的导航守卫, 保护所有的 backend 的页面, **凡事访问到 backend 的页面 都需要检查登录状态, 避免用户直接绕开访问**

[导航守卫](./src/router/README.md#导航守卫)

## 对接后端

开始对接后端 api。对接后端 API 需要使用 [axios](https://www.axios-http.cn/docs/intro)

安装 axios
```shell
npm install axios
```

### 对接登录

首先，我们在 `src/api/login.js` 中封装后端登录的接口

```js
import httpClient from "./client.js"

// 对应后端 login 接口
export async function LOGIN(username, password) {
    return await httpClient.post("/users/token", {
        username: username,
        password: password,
    })
}
```

在 [Login.vue](./src/views/Login.vue) 中对接登录接口

```js
const handleSubmit = (data) => {
    if (!data.errors) {
        // // 对接后端 API
        try {
            let result = LOGIN(form.username, form.password)
            console.log(result)
        } catch (error) {
            Message.error(error.message)
        }
    }
}
```

访问错误

```
Access to XMLHttpRequest at 'http://localhost:8010/vblog/api/v1/users/token' from origin 'http://localhost:5173' has been blocked by CORS policy: Response to preflight request doesn't pass access control check: No 'Access-Control-Allow-Origin' header is present on the requested resource.
```

这里我们的前端和后端运行在不同的域下，由于后端不允许跨域，导致浏览器的跨域策略失败

由于我们开发的 **单体项目**，前后端的开发是分离的，域不必分离，都是用同一个域

这里我们使用 vite 做为代理(Niginx 角色)，完成 API 访问的代理，避免浏览器直接向后端发起请求

在 [vite.config.js] 中配置反向代理

```js
// 配置服务端代理
server:{
proxy: {
  '/vblog/api/v1': {
    target: 'http://localhost:8010/'
  }
}
}
```

`axios` 的实例对象配置如下

```js
import axios from "axios";

// 创建 axios 客户端


const instance = axios.create({
    baseURL: "",  // 后端地址，被 vite 代理了。即 http://localhost:5173/vblog/api/v1/.... => http://localhost:8010/vblog/api/v1/...
    timeout: 5000, // 超时时间
    headers: {
        'Accept': 'application/json',
    }
})

export default instance;
```

#### 响应拦截

拦截响应，统一处理后端响应的数据。在 [client.js](./src/api/client.js)

```js
import axios from "axios";

// 创建 axios 客户端

const instance = axios.create({
    baseURL: "",  // 后端地址，被 vite 代理了。即 http://localhost:5173/vblog/api/v1/....
    timeout: 5000, // 超时时间
    headers: {
        'Content-Type': 'application/json',
    }
})

// 为 instance 添加拦截器，通过拦截器做异常的统一处理

// instance.interceptors.request.use() // 请求拦截
instance.interceptors.response.use(
    // 成功就只返回data
    (response) => {
        return response.data;
    },
    // 异常
    (error) => {
        let message = error.message;
        if (error.response && error.response.data) {
            message = error.response.data.data.message;
        }
        return Promise.reject(new Error(message));
    }
) // 响应拦截

export default instance;
```

#### 完成后端对接

在 [Login.vue](./src/views/Login.vue) 中

```js
const handleSubmit = async (data) => {
    if (!data.errors) {
        // // 对接后端 API
        try {
            let result =  await LOGIN(form.username, form.password)

            console.log(result)

            // 保存一个全局状态，最好的方式使用 localStorage 保存登录状态，方便其他标签页或组件
            loginState.value.username = result.data.username
            loginState.value.token = result.data.access_token
            loginState.value.isLogin = true

            // 跳转
            if (!currentRoute.query.hasOwnProperty("redirect")) {
                await router.push({name: "backend"})
            } else {
                await router.push({name: currentRoute.query.redirect})
            }
        } catch (error) {
            Message.error(error.message)
        }
    }
}
```

### 对接后台管理

对接后台文章列表

1. 获取数据的接口封装
```js
import httpClient from "./client.js"

// 对应后端 blogs 接口
export function BLOG_LIST(queryBlogsRequest) {
    return httpClient.get("/vblog/api/v1/blogs", queryBlogsRequest)
}
```
2. 展示数据
```html
<script setup name="BlogList">

import {BLOG_LIST} from "@/api/blog.js"
import {reactive} from "vue";
import {Message} from "@arco-design/web-vue";

const columns = reactive([
  {title: "ID", dataIndex: "id"},
  {title: "标题", dataIndex: "title"},
])

const data = reactive({
  items: [],
  total: 0
})

async function getBlogList(queryBlogsRequest) {
  try {
    let result = await BLOG_LIST(queryBlogsRequest)

    data.items = result.data.items;
    data.total = result.data.total;

  } catch (error) {
    Message.error(error.message)
  }
}


getBlogList({
  page_size:20,
  page_number:1,
})

</script>

<template>
  <div class="blogs-container">
    <a-table :columns="columns" :data="data.items" />
  </div>
</template>

<style scoped>

</style>
```

## 致谢

- [normalize.css](https://necolas.github.io/normalize.css/)
- [vue](https://cn.vuejs.org/)

- [arco design](https://arco.design/)
- [ant design](https://ant-design.antgroup.com/index-cn)
- [element plus](https://element-plus.org/zh-CN/)
- [TDesign](https://tdesign.tencent.com/)

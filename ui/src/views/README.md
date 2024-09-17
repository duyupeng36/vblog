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









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

在做 Home 页面之前，先清理掉脚手架为我们生成的页面

清理 `App.vue`, 只保留 `router` 视图部分, 其他部分删除掉
















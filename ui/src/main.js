import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from '@/App.vue'
import router from '@/router'

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


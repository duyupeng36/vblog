import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueSetupExtend from "vite-plugin-vue-setup-extend" /* 导入插件 */


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueSetupExtend()
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // 配置服务端代理
  server:{
    proxy: {
      '/vblog/api/v1': {
        target: 'http://localhost:8010/',
      }
    }
  }
})

import axios from "axios";

// 创建 axios 客户端


const instance = axios.create({
    baseURL: "",  // 后端地址，被 vite 代理了。即 http://localhost:5173/vblog/api/v1/....
    timeout: 5000, // 超时时间
    headers: {
        'Accept': 'application/json',
    }
})

export default instance;

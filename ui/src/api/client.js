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

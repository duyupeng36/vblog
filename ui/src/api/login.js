import httpClient from "./client.js"

// 对应后端 login 接口
export function LOGIN(loginRequest) {
    return httpClient.post("/vblog/api/v1/users/token", loginRequest)
}

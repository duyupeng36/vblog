import httpClient from "./index.js"

// 对应后端 login 接口
export async function LOGIN(username, password) {
    return await httpClient.post("/vblog/api/v1/users/token", {
        username: username,
        password: password,
    })
}

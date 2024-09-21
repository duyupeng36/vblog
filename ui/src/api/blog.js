import httpClient from "./client.js"

// 对应后端 login 接口
export function BLOG_LIST(queryBlogsRequest) {
    return httpClient.get("/vblog/api/v1/blogs", queryBlogsRequest)
}

# 博客管理

博客管理需要实现下面几个 RESTfull API 和其对应的业务


|   请求方法   | URI                        | 动作     | 业务接口           |
| :------: | :------------------------- | :----- | -------------- |
|  `GET`   | `/vblogs/api/v1/blogs`     | 获取博客列表 | `QueryBlog`    |
|  `POST`  | `/vblogs/api/v1/blogs`     | 创建博客   | `CreateBlog`   |
|  `GET`   | `/vblogs/api/v1/blogs/:id` | 获取一篇文章 | `DescribeBlog` |
|  `PUT`   | `/vblogs/api/v1/blogs/:id` | 修改一篇文章 | `UpdateBlog`   |
| `DELETE` | `/vblogs/api/v1/blogs/:id` | 删除一篇文章 | `DeleteBlog`   |




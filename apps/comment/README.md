# 评论管理


评论 管理需要实现下面几个 RESTfull API 和其对应的业务


|   请求方法   | URI                           | 动作     | 业务接口              |
|:--------:|:------------------------------|:-------|-------------------|
|  `GET`   | `/vblogs/api/v1/comments`     | 获取博客列表 | `QueryComment`    |
|  `POST`  | `/vblogs/api/v1/comments`     | 创建博客   | `CreateComment`   |
|  `GET`   | `/vblogs/api/v1/comments/:id` | 获取一篇文章 | `DescribeComment` |
|  `PUT`   | `/vblogs/api/v1/comments/:id` | 修改一篇文章 | `UpdateComment`   |
| `DELETE` | `/vblogs/api/v1/comments/:id` | 删除一篇文章 | `DeleteComment`   |




package apps

import (
	// 注册 handler
	_ "vblog/apps/blog/api"
	_ "vblog/apps/comment/api"
	_ "vblog/apps/user/api"

	// 注册业务具体实现
	_ "vblog/apps/blog/impl"
	_ "vblog/apps/comment/impl"
	_ "vblog/apps/user/impl"
)

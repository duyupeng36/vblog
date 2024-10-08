// Pakcage impl 完全被封装，只能通过 IoC 访问
package impl

import (
	"vblog/apps/blog"
	"vblog/conf"
	"vblog/ioc"

	"gorm.io/gorm"
)

// 断言 *impl 是否满足 blog.Service 接口的约束
var _ blog.Service = &impl{
	// 不能这样直接初始化
	// 包被加载的时候，全局变量会被有限赋值
	// 此时，conf 包可能还没有初始化
	// c: conf.C(),
}

// impl 对象上实现了 blog.Service。它就充当了 控制器
type impl struct {
	// c *conf.Config
	db *gorm.DB
}

// 业务实例如何读取配置对象？通常，为了 Impl 定义一个 Init 方法
// 这个方法在 config 初始化完成之后，才调用
func (i *impl) Init() error {
	// 每个业务都建立一个 数据库连接池时比较浪费的，
	// 通常，在配置对象上建立连接，然后每个业务都读取这个连接即可
	var err error
	i.db, err = conf.C().MySQL.GORM() // 这样，业务对象就拿到数据库连接了
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) Name() string {
	return blog.AppName
}

// 实现 impl 包的初始化方法，import 时自动注册到 ioc 容器中
func init() {
	ioc.Container().RegistryController(&impl{})
}

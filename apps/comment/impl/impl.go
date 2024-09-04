package impl

import (
	"vblog/apps/comment"
	"vblog/conf"
	"vblog/ioc"

	"gorm.io/gorm"
)

var _ comment.Service = &impl{}

type impl struct {
	db *gorm.DB
}

// 业务实例如何读取配置对象？通常，为了 Impl 定义一个 init 方法
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
	return comment.AppName
}

func init() {
	ioc.Container().RegistryController(&impl{})
}

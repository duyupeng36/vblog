package impl

import (
	"vblog/apps/user"
	"vblog/conf"
	"vblog/ioc"

	"gorm.io/gorm"
)

var _ user.Service = &impl{}

type impl struct {
	db *gorm.DB
}

func (i *impl) Init() (err error) {

	i.db, err = conf.C().MySQL.GORM()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) Name() string {
	return user.AppName
}

func init() {
	ioc.Container().RegistryController(&impl{})
}

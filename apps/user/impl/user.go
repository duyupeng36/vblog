package impl

import (
	"context"
	"vblog/apps/user"
	"vblog/utils/errors"
)

// 注册用户
func (i *impl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {

	if err := in.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	// 再次验证两次密码输入是否一致
	if in.Password != in.PasswordConfirm {
		return nil, errors.NewBadRequestError("Inconsistent passwords entered twice")
	}

	// 创建一个用户
	ins := user.NewUser(in.UserInfo)
	if err := ins.BuildHashedPassword(); err != nil {
		return nil, err
	}

	// 入库
	if err := i.db.WithContext(ctx).Model(&user.User{}).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

// 删除用户
func (i *impl) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.User, error) {

	ins := user.NewUser(user.NewUserInfo())

	if err := i.db.WithContext(ctx).Where("id = ?", in.Id).Take(ins).Error; err != nil {
		return nil, errors.NewNotFoundError(err.Error())
	}

	if err := i.db.WithContext(ctx).Where("id = ?", in.Id).Delete(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

// 登录
func (i *impl) Login(ctx context.Context, in *user.LoginRequest) (*user.Token, error) {

	if err := in.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	ins := user.NewUser(user.NewUserInfo())

	if err := i.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", in.Username).Take(&ins).Error; err != nil {
		return nil, errors.NewNotFoundError(err.Error())
	}

	if err := ins.CheckPassword(in.Password); err != nil {
		return nil, errors.NewUnauthorizedError(err.Error())
	}

	tk := user.NewToken(ins.Username)

	if err := i.db.WithContext(ctx).Model(&user.Token{}).Create(tk).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

// CheckToken 检查 Token
func (i *impl) CheckToken(ctx context.Context, in *user.CheckTokenRequest) (*user.Token, error) {
	tk := user.NewToken("")

	if err := i.db.WithContext(ctx).Model(&user.Token{}).Where("access_token = ?", in.AccessToken).Take(tk).Error; err != nil {
		return nil, errors.NewUnauthorizedError(err.Error())
	}

	return tk, nil
}

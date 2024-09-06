package user

import (
	"context"
	"vblog/utils"
)

const (
	AppName = "users"
)

type Service interface {
	// 注册用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	// 登录
	Login(context.Context, *LoginRequest) (*Token, error)
	// 检查 Token
	CheckToken(context.Context, *CheckTokenRequest) (*Token, error)
}

type CreateUserRequest struct {
	*UserInfo
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

func (r *CreateUserRequest) Validate() error {
	return utils.Validate(r)
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		UserInfo: NewUserInfo(),
	}
}

type DeleteUserRequest struct {
	Id int
}

func NewDeleteUserRequest(id int) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: id,
	}
}

type LoginRequest struct {
	*UserInfo
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{
		UserInfo: NewUserInfo(),
	}
}

type CheckTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func NewCheckTokenRequest(token string) *CheckTokenRequest {
	return &CheckTokenRequest{
		AccessToken: token,
	}
}

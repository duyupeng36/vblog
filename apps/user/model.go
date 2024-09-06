package user

import (
	"encoding/base64"
	"encoding/json"
	"time"
	"vblog/utils"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type Meta struct {
	ID        int   `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}

type User struct {
	*Meta
	*UserInfo
}

func (u *User) String() string {

	b, _ := json.MarshalIndent(u, "", "  ")
	return string(b)
}

func NewUser(info *UserInfo) *User {
	return &User{
		Meta:     NewMeta(),
		UserInfo: info,
	}
}

// BuildHashedPassword 构建 hashed 密码
// password -> hash: 一个密码对应一个 hash，无法防止彩虹表：hash --> password
// 解决方案：添加随机的字节，并将字节附加到 hash 的结果中，方便后续验证
// salt.password -> hash -> salt.hash
// 验证：查询出用户的 salt.hash，提取  salt，加密登录时提供的密码 salt.password -> hash ，对比 salt.hash
func (u *User) BuildHashedPassword() error {
	// hash := sha256.New()
	// hash.Write([]byte(u.Password))
	// u.Password = base64.StdEncoding.EncodeToString(hash.Sum(nil))
	// return nil

	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = base64.StdEncoding.EncodeToString(b)
	return nil
}

func (u *User) CheckPassword(password string) error {

	bytesHashedPassword, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(bytesHashedPassword, []byte(password))
}

type UserInfo struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (ui *UserInfo) Validate() error {
	return utils.Validate(ui)
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

type Token struct {
	*Meta
	Username    string `json:"username"`
	AccessToken string `json:"acess_token"`
}

func NewToken(username string) *Token {
	return &Token{
		Meta:        NewMeta(),
		Username:    username,
		AccessToken: xid.New().String(),
	}
}

func (tk *Token) String() string {
	b, _ := json.MarshalIndent(tk, "", "  ")
	return string(b)
}

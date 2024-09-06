package impl_test

import (
	"testing"
	"vblog/apps/user"
)

func TestCreateUser(t *testing.T) {

	in := user.NewCreateUserRequest()
	in.Username = "admin1"
	in.Password = "123456"
	in.PasswordConfirm = "123456"

	ins, err := controller.CreateUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteUser(t *testing.T) {

	in := user.NewDeleteUserRequest(2)

	ins, err := controller.DeleteUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestLogin(t *testing.T) {
	in := user.NewLoginRequest()
	in.Username = "admin"
	in.Password = "123456"

	tk, err := controller.Login(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestCheckToken(t *testing.T) {
	in := user.NewCheckTokenRequest("crd70cfdunf7ut2el37g")

	tk, err := controller.CheckToken(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)

}
